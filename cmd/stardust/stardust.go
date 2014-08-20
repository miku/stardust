package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/miku/stardust"
)

func main() {
	app := cli.NewApp()
	app.Name = "stardust"
	app.Usage = "String similarity measures for tab separated values."
	app.Author = "Martin Czygan"
	app.Email = "martin.czygan@gmail.com"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "f",
			Value: "1,2",
			Usage: "c1,c2 the two columns to use for the comparison",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "ngram",
			Usage:       "Ngram similarity",
			Description: "Compute Ngram similarity, which lies between 0 and 1.",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, err := stardust.NgramSimilaritySize(r.Left(), r.Right(), c.Int("size"))
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "size, s",
					Value: 3,
					Usage: "value of n",
				},
			},
		},
		{
			Name:  "hamming",
			Usage: "Hamming distance",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, _ := stardust.HammingDistance(r.Left(), r.Right())
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:  "levenshtein",
			Usage: "Levenshtein distance",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, _ := stardust.LevenshteinDistance(r.Left(), r.Right())
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:        "jaro",
			Usage:       "Jaro similarity",
			Description: "Similar to Ngram, but faster.",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, _ := stardust.JaroSimilarity(r.Left(), r.Right())
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:  "plain",
			Usage: "Plain passthrough (for IO benchmarks)",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					fmt.Printf("%s\n", strings.Join(r.Fields, "\t"))
				}
			},
		},
	}
	app.Run(os.Args)
}
