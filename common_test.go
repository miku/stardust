package stardust

import (
	"math"
	"reflect"
	"testing"

	"github.com/juju/utils/set"
)

// AlmostEqualRelative is a float comparison helper
// Via http://randomascii.wordpress.com/2012/02/25/comparing-floating-point-numbers-2012-edition/
func AlmostEqualRelative(a, b, maxRelDiff float64) bool {
	diff := math.Abs(a - b)
	A := math.Abs(a)
	B := math.Abs(b)
	largest := A
	if B > A {
		largest = B
	}
	if diff <= largest*maxRelDiff {
		return true
	}
	return false
}

func TestJaccardSets(t *testing.T) {
	var jaccardSetsTests = []struct {
		a   set.Strings
		b   set.Strings
		out float64
	}{
		{set.NewStrings("a"), set.NewStrings("a"), 1.0},
		{set.NewStrings("a"), set.NewStrings("a", "b"), 0.5},
		{set.NewStrings("a"), set.NewStrings("a", "b", "c"), 1.0 / 3},
		{set.NewStrings("a", "b"), set.NewStrings("a", "b", "c"), 2.0 / 3},
	}

	for _, tt := range jaccardSetsTests {
		out := JaccardSets(tt.a, tt.b)
		if out != tt.out {
			t.Errorf("Jaccard(%v, %v) => %f, want: %f", tt.a, tt.b, out, tt.out)
		}
	}
}

func TestNgram(t *testing.T) {
	var ngramTests = []struct {
		s   string
		n   int
		out set.Strings
	}{
		{"abc", 2, set.NewStrings("ab", "bc")},
		{"abc", 6, set.NewStrings()},
		{"abc", 3, set.NewStrings("abc")},
		{"abc", 1, set.NewStrings("a", "b", "c")},
		{"abc", 0, set.NewStrings()},
		{"abc", -10, set.NewStrings()},
		{"abc def", 3, set.NewStrings("abc", "bc ", "c d", " de", "def")},
	}
	for _, tt := range ngramTests {
		out := Ngrams(tt.s, tt.n)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("Ngrams(%s, %d) => %v, want: %v", tt.s, tt.n, out, tt.out)
		}
	}
}

func TestNgramSimilarity(t *testing.T) {
	var ngramSimilarityTests = []struct {
		a   string
		b   string
		out float64
	}{
		{"Hello World", "Hello Earth", 0.285714},
		{"Hello World", "Hello Wookie", 0.461538},
		{"Hello World", "Hello", 0.333333},
		{"The quick brown fox", "The qiuck brown fox", 0.619048},
		{"The quick brown fox", "The quick brown fox", 1.000000},
	}
	for _, tt := range ngramSimilarityTests {
		out := NgramSimilarity(tt.a, tt.b)
		if !AlmostEqualRelative(out, tt.out, 1e-5) {
			t.Errorf("NgramSimilarity(%s, %s) => %f, want: %f", tt.a, tt.b, out, tt.out)
		}
	}
}

func TestNgramSimilaritySize(t *testing.T) {
	var ngramSimilaritySizeTests = []struct {
		a    string
		b    string
		size int
		out  float64
	}{
		{"Hello World", "Hello Earth", 2, 0.333333},
		{"Hello World", "Hello Wookie", 7, 0.222222},
		{"Hello World", "Hello", 1, 0.500000},
		{"The quick brown fox", "The qiuck brown fox", 3, 0.619048},
		{"The quick brown fox", "The qiuck brown fox", 2, 0.714286},
		{"The quick brown fox", "The qiuck brown fox", 1, 1.000000},
	}
	for _, tt := range ngramSimilaritySizeTests {
		out := NgramSimilaritySize(tt.a, tt.b, tt.size)
		if !AlmostEqualRelative(out, tt.out, 1e-5) {
			t.Errorf("NgramSimilaritySize(%s, %s, %d) => %f, want: %f", tt.a, tt.b, tt.size, out, tt.out)
		}
	}
}

func TestHammingDistance(t *testing.T) {
	var hammingDistanceTests = []struct {
		a   string
		b   string
		out int
	}{
		{"Hello World", "Hello Earth", 4},
	}
	for _, tt := range hammingDistanceTests {
		out, _ := HammingDistance(tt.a, tt.b)
		if out != tt.out {
			t.Errorf("HammingDistance(%s, %s) => %d, want: %d", tt.a, tt.b, out, tt.out)
		}
	}
}
