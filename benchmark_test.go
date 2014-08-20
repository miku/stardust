package stardust

import "testing"

func BenchmarkNgramSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NgramSimilarity("Hello World", "Hey young world")
	}
}

func BenchmarkHammingDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HammingDistance("Hello World", "Hey young world")
	}
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LevenshteinDistance("Hello World", "Hey young world")
	}
}

func BenchmarkJaroDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JaroDistance("Hello World", "Hey young world")
	}
}

func BenchmarkJaroWinklerDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JaroWinklerDistance("Hello World", "Hey young world", 0.1, 3)
	}
}
