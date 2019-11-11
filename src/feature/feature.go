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
	Covariance(other MeasurableFeature) float64
	Similarity(basedOn MeasurableFeature) float64
}
