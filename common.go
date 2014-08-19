package stardust

import (
	"errors"
	"math"
	"strings"

	"github.com/juju/utils/set"
)

const Version = "0.0.1"

func CompleteString(pool []string, prefix string) []string {
	var candidates []string
	for _, value := range pool {
		if strings.HasPrefix(value, prefix) {
			candidates = append(candidates, value)
		}
	}
	return candidates
}

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

func NgramSimilaritySize(s, t string, n int) (float64, error) {
	sset := Ngrams(s, n)
	tset := Ngrams(t, n)
	if tset.Size() == 0 && sset.Size() == 0 {
		return 0, nil
	}
	return JaccardSets(sset, tset), nil
}

func NgramSimilarity(s, t string) (float64, error) {
	return NgramSimilaritySize(s, t, 3)
}

func HammingDistance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings must be of equal length")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

func maxInt(numbers ...int) int {
	result := math.MinInt64
	for _, k := range numbers {
		if k > result {
			result = k
		}
	}
	return result
}

func minInt(numbers ...int) int {
	result := math.MaxInt64
	for _, k := range numbers {
		if k < result {
			result = k
		}
	}
	return result
}

func LevenshteinDistance(s, t string) (int, error) {
	if len(s) < len(t) {
		return LevenshteinDistance(t, s)
	}
	if len(t) == 0 {
		return len(s), nil
	}

	previous := make([]int, len(t)+1)
	for i, c := range s {
		current := []int{i + 1}
		for j, d := range t {
			insertions := previous[j+1] + 1
			deletions := current[j] + 1
			cost := 0
			if c != d {
				cost = 1
			}
			subtitutions := previous[j] + cost
			current = append(current, minInt(insertions, deletions, subtitutions))
		}
		previous = current
	}
	return previous[len(previous)-1], nil
}

func JaroSimilarity(a, b string) (float64, error) {
	la := float64(len(a))
	lb := float64(len(b))

	matchRange := int(math.Floor(math.Max(la, lb)/2.0)) - 1
	matchRange = int(math.Max(0, float64(matchRange-1)))
	var matches, halfs float64
	transposed := make([]bool, len(b))

	for i := 0; i < len(a); i++ {
		start := int(math.Max(0, float64(i-matchRange)))
		end := int(math.Min(lb-1, float64(i+matchRange)))

		for j := start; j <= end; j++ {
			if transposed[j] {
				continue
			}

			if a[i] == b[j] {
				if i != j {
					halfs++
				}
				matches++
				transposed[j] = true
				break
			}
		}
	}

	if matches == 0 {
		return 0, nil
	}

	transposes := math.Floor(float64(halfs / 2))

	return ((matches / la) + (matches / lb) + (matches-transposes)/matches) / 3.0, nil
}
