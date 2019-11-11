package feature

import (
	"fingerprint"
	mapset "github.com/deckarep/golang-set"
)

type WinnowFeature struct {
	belongsTo string
	set       mapset.Set
}

func NewFeature(belongsTo string) *WinnowFeature {
	s := &WinnowFeature{
		belongsTo: belongsTo,
		set:       mapset.NewSet(),
	}
	return s
}

func (fp *WinnowFeature) AddFingerprint(fingerprint *fingerprint.Fingerprint) {
	fp.set.Add(fingerprint.ToString())
}

func (fp *WinnowFeature) Covariance(other MeasurableFeature) float64 {
	var measurement float64

	switch other.(type) {
	case *WinnowFeature:
		otherf := other.(*WinnowFeature)
		intersect := fp.set.Intersect(otherf.set)
		union := fp.set.Union(otherf.set)
		measurement = float64(intersect.Cardinality()) / float64(union.Cardinality())
		break
	}

	return measurement
}

func (fp *WinnowFeature) Similarity(basedOn MeasurableFeature) float64 {
	var measurement float64

	switch basedOn.(type) {
	case *WinnowFeature:
		otherf := basedOn.(*WinnowFeature)
		intersect := fp.set.Intersect(otherf.set)
		measurement = float64(intersect.Cardinality()) / float64(fp.set.Cardinality())
		break
	}

	return measurement
}
