package measurement

import (
	"feature"
	"math"
)

func Straightforward(fs []feature.MeasurableFeature) float64 {
	maxMatch := math.Max(fs[0].Similarity(fs[1]), fs[1].Similarity(fs[0]))
	for k := 2; k < len(fs); k++ {
		match := math.Max(fs[0].Similarity(fs[k]), fs[k].Similarity(fs[0]))
		if match > maxMatch {
			maxMatch = match
		}
	}

	return maxMatch
}
