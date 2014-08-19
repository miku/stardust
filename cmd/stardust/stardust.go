package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strings"

	"github.com/miku/stardust"
)

func main() {

	distanceFuncMap := map[string]interface{}{
		"hamming":     stardust.HammingDistance,
		"levenshtein": stardust.LevenshteinDistance,
		"ngram":       stardust.NgramSimilarity,
		"jaro":        stardust.JaroSimilarity,
	}

	measure := flag.String("m", "ngram", "distance measure")
	listFuncs := flag.Bool("l", false, "list available measures")
	version := flag.Bool("v", false, "prints current program version")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()

	if *version {
		fmt.Println(stardust.Version)
		os.Exit(0)
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *listFuncs {
		var keys []string
		for k := range distanceFuncMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Println(k)
		}
		return
	}

	var PrintUsage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] STRING STRING\n", os.Args[0])
		flag.PrintDefaults()
	}

	// check for stdin or filename
	if len(flag.Args()) == 1 {
		if flag.Args()[0] == "-" {
			fmt.Fprintf(os.Stderr, "reading from stdin...\n")
			return
		} else {
			// try open a file
			if _, err := os.Stat(flag.Args()[0]); err == nil {
				fmt.Fprintf(os.Stderr, "reading from file...\n")
				return
			}
		}
	}

	if len(flag.Args()) != 2 {
		PrintUsage()
		os.Exit(1)
	}

	// find the right prefix function
	var keys []string
	for k := range distanceFuncMap {
		keys = append(keys, k)
	}
	result := stardust.CompleteString(keys, *measure)
	if len(result) > 1 {
		log.Fatalf("ambiguous name: %s\n", strings.Join(result, ", "))
	} else if len(result) == 0 {
		log.Fatal("no such distance function")
	}
	fn, _ := distanceFuncMap[result[0]]

	a := flag.Args()[0]
	b := flag.Args()[1]

	// we have both int and float functions
	switch fn.(type) {
	default:
		log.Fatal("unknown signature")
	case func(string, string) (float64, error):
		result, err := fn.(func(string, string) (float64, error))(a, b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	case func(string, string) (int, error):
		result, err := fn.(func(string, string) (int, error))(a, b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
