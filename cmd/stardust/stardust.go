package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/miku/stardust"
)

func main() {

	distanceFuncMap := map[string]interface{}{
		"ngram":       stardust.NgramSimilarity,
		"hamming":     stardust.HammingDistance,
		"levenshtein": stardust.LevenshteinDistance,
	}

	measure := flag.String("f", "ngram", "distance measure")
	listFuncs := flag.Bool("l", false, "list available measures")

	flag.Parse()

	if *listFuncs {
		for k, _ := range distanceFuncMap {
			fmt.Println(k)
		}
		return
	}

	var PrintUsage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] STRING STRING\n", os.Args[0])
		flag.PrintDefaults()
	}

	if len(flag.Args()) != 2 {
		PrintUsage()
		os.Exit(1)
	}

	a := flag.Args()[0]
	b := flag.Args()[1]

	// find the right prefix function
	var keys []string
	for k := range distanceFuncMap {
		keys = append(keys, k)
	}
	result := stardust.CompleteString(keys, *measure)
	if len(result) > 1 {
		log.Fatalf("ambiguous measure name: %s\n", strings.Join(result, ", "))
	}
	if len(result) == 0 {
		log.Fatal("no such distance function")
	}
	fn, _ := distanceFuncMap[result[0]]

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
