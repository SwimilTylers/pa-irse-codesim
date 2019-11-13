package feature

import (
	"fingerprint"
	mapset "github.com/deckarep/golang-set"
	"log"
	"math"
)

type WinnowFeature struct {
	belongsTo string
	set       mapset.Set
}

func NewWinnowFeature(belongsTo string) *WinnowFeature {
	s := &WinnowFeature{
		belongsTo: belongsTo,
		set:       mapset.NewSet(),
	}
	return s
}

func (wf *WinnowFeature) AddFingerprint(fingerprint *fingerprint.Fingerprint) {
	wf.set.Add(fingerprint.ToString())
}

func (wf *WinnowFeature) AddDimension() {
	// nullify it, meaningless to this feature
}

func (wf *WinnowFeature) GetSimilarity(other MeasurableFeature, simMode string) float64 {
	var simFunc func(MeasurableFeature) float64

	switch simMode {
	case "overlap": // overlap coefficient
		simFunc = wf.overlap
		break
	case "smc": // simple matching coefficient
		simFunc = wf.smc
		break
	case "jaccard": // jaccard similarity coefficient, or IoU
		simFunc = wf.jaccard
		break
	default:
		log.Fatalln("Winnow-based Feature do not support this similarity:", simMode)
	}

	return simFunc(other)
}

func (wf *WinnowFeature) overlap(other MeasurableFeature) float64 {
	var measurement float64

	switch other.(type) {
	case *WinnowFeature:
		otherf := other.(*WinnowFeature)
		intersect := wf.set.Intersect(otherf.set)
		smallerCard := math.Min(float64(wf.set.Cardinality()), float64(otherf.set.Cardinality()))
		measurement = float64(intersect.Cardinality()) / smallerCard
		break
	}

	return measurement
}

func (wf *WinnowFeature) smc(other MeasurableFeature) float64 {
	var measurement float64

	switch other.(type) {
	case *WinnowFeature:
		otherf := other.(*WinnowFeature)
		intersect := wf.set.Intersect(otherf.set)
		sumCard := float64(wf.set.Cardinality()) + float64(otherf.set.Cardinality())
		measurement = float64(intersect.Cardinality()) / sumCard
		break
	}

	return measurement
}

func (wf *WinnowFeature) jaccard(other MeasurableFeature) float64 {
	var measurement float64

	switch other.(type) {
	case *WinnowFeature:
		otherf := other.(*WinnowFeature)
		intersect := wf.set.Intersect(otherf.set)
		union := wf.set.Union(otherf.set)
		measurement = float64(intersect.Cardinality()) / float64(union.Cardinality())
		break
	}

	return measurement
}
