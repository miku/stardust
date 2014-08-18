package main

import (
	"fmt"

	"github.com/miku/stardust"
)

func main() {
	fmt.Printf("The unigram-distance between stardust and strdist is %f\n", stardust.NgramSimilaritySize("stardust", "strdist", 1))
}
