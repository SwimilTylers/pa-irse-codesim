package feature

import (
	"fingerprint"
)

type FPFeature interface {
	FlexibleFeature
	MeasurableFeature
}

type FlexibleFeature interface {
	AddFingerprint(fingerprint *fingerprint.Fingerprint)
	AddDimension()
}

type MeasurableFeature interface {
	GetSimilarity(other MeasurableFeature, simMode string) float64
}
