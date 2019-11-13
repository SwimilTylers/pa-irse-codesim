package measurement

import (
	"feature"
	"log"
	"math"
)

func MeasureTwoFeatures(fs []feature.MeasurableFeature, mMode string, simMode string) float64 {
	if len(fs) != 2 {
		log.Fatalln("Need Exactly two features")
	}

	var measure float64

	switch mMode {
	case "max":
		measure = math.Max(MaxMatch(fs[0], fs[1:], simMode), MaxMatch(fs[1], fs[:1], simMode))
		break
	case "mean":
		measure = (MeanMatch(fs[0], fs[1:], simMode) + MeanMatch(fs[1], fs[:1], simMode)) / 2
		break
	default:
		log.Fatalln("Unsupported Measurement:", mMode)
	}

	return measure
}
