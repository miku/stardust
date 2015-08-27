package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/gyuho/goling/similar"
	"github.com/karlek/nyfiken/distance"
	"github.com/miku/stardust"
)

func main() {
	app := cli.NewApp()
	app.Name = "stardust"
	app.Usage = "String similarity measures for tab separated values."
	app.Author = "Martin Czygan"
	app.Email = "martin.czygan@gmail.com"
	app.Version = "0.1.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "fields, f",
			Value: "1,2",
			Usage: "c1,c2 the two columns to use for the comparison",
		},
		cli.StringFlag{
			Name:  "delimiter, d",
			Value: "\t",
			Usage: "column delimiter (defaults to tab)",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "adhoc",
			Usage:       "Adhoc distance",
			Description: "Ad-hoc percentage difference found on the web (https://godoc.org/github.com/karlek/nyfiken/distance).",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure := distance.Approx(r.Left(), r.Right())
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:        "cosine",
			Usage:       "Cosine word-wise",
			Description: "A a measure of similarity between two vectors. The bigger the return value is, the more similar the two texts are.",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure := similar.Cosine(r.Left(), r.Right())
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:        "coslev",
			Usage:       "Cosine word-wise and levenshtein combined",
			Description: "Experimenal.",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure := similar.Get(r.Left(), r.Right())
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:        "dice",
			Usage:       "Sørensen–Dice coefficient",
			Description: "Semimetric version of the Jaccard index.",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, _ := stardust.SorensenDiceDistance(r.Left(), r.Right())
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},

		{
			Name:  "hamming",
			Usage: "Hamming distance",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, err := stardust.HammingDistance(r.Left(), r.Right())
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:        "jaro",
			Usage:       "Jaro distance",
			Description: "Similar to Ngram, but faster.",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, err := stardust.JaroDistance(r.Left(), r.Right())
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:        "jaro-winkler",
			Usage:       "Jaro-Winkler distance",
			Description: "It is a variant of the Jaro distance metric.",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, err := stardust.JaroWinklerDistance(r.Left(), r.Right(), c.Float64("boost"), c.Int("size"))
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
			Flags: []cli.Flag{
				cli.Float64Flag{
					Name:  "boost, b",
					Value: 0.5,
					Usage: "boost factor",
				},
				cli.IntFlag{
					Name:  "size, p",
					Value: 3,
					Usage: "prefix size",
				},
			},
		},
		{
			Name:  "levenshtein",
			Usage: "Levenshtein distance",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, err := stardust.LevenshteinDistance(r.Left(), r.Right())
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%s\t%v\n", strings.Join(r.Fields, "\t"), measure)
				}
			},
		},
		{
			Name:        "ngram",
			Usage:       "Ngram distance",
			Description: "Compute Ngram distance, which lies between 0 and 1 (equal).",
			Action: func(c *cli.Context) {
				records := stardust.RecordGenerator(c)
				for r := range records {
					measure, err := stardust.NgramDistanceSize(r.Left(), r.Right(), c.Int("size"))
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
