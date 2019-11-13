package measurement

import (
	"feature"
)

func MaxMatch(codeTested feature.MeasurableFeature, codeBase []feature.MeasurableFeature, simMode string) float64 {
	maxMatch := codeTested.GetSimilarity(codeBase[0], simMode)

	for i := 1; i < len(codeBase); i++ {
		match := codeTested.GetSimilarity(codeBase[i], simMode)
		if match > maxMatch {
			maxMatch = match
		}
	}

	return maxMatch
}

func MeanMatch(codeTested feature.MeasurableFeature, codeBase []feature.MeasurableFeature, simMode string) float64 {
	var match float64

	for _, value := range codeBase {
		match += codeTested.GetSimilarity(value, simMode)
	}

	return match / float64(len(codeBase))
}
