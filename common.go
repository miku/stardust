package stardust

import (
	"errors"

	"github.com/juju/utils/set"
)

const Version = "0.0.1"

func JaccardSets(a, b set.Strings) float64 {
	return float64(a.Intersection(b).Size()) / float64(a.Union(b).Size())
}

func Unigrams(s string) set.Strings {
	return Ngrams(s, 1)
}

func Bigrams(s string) set.Strings {
	return Ngrams(s, 2)
}

func Trigrams(s string) set.Strings {
	return Ngrams(s, 2)
}

func Ngrams(s string, n int) set.Strings {
	result := set.NewStrings()
	if n > 0 {
		lastIndex := len(s) - n + 1
		for i := 0; i < lastIndex; i++ {
			result.Add(s[i : i+n])
		}
	}
	return result
}

func NgramSimilaritySize(s, t string, n int) float64 {
	sg := Ngrams(s, n)
	tg := Ngrams(t, n)
	return JaccardSets(sg, tg)
}

func NgramSimilarity(s, t string) float64 {
	return NgramSimilaritySize(s, t, 3)
}

func HammingDistance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings but be of equal length")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
