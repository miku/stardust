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

func BenchmarkJaroSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JaroSimilarity("Hello World", "Hey young world")
	}
}

func BenchmarkJaroWinklerSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JaroWinklerSimilarity("Hello World", "Hey young world", 0.1, 3)
	}
}
