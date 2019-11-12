package feature

import (
	"fingerprint"
)

type MultiWinnowFeature struct {
	belongsTo string
	top       int
	winnows   []*WinnowFeature
}

func NewMultiWinnowFeature(belongsTo string) *MultiWinnowFeature {
	return &MultiWinnowFeature{
		belongsTo: belongsTo,
		top:       -1,
		winnows:   make([]*WinnowFeature, 0, 10),
	}
}

func (mwf *MultiWinnowFeature) AddFingerprint(fingerprint *fingerprint.Fingerprint) {
	mwf.winnows[mwf.top].AddFingerprint(fingerprint)
}

func (mwf *MultiWinnowFeature) AddDimension() {
	mwf.top++
	if mwf.top == cap(mwf.winnows) {
		newWs := make([]*WinnowFeature, mwf.top*2)
		copy(newWs, mwf.winnows)
		mwf.winnows = newWs
	}
	mwf.winnows = mwf.winnows[:mwf.top+1]
	mwf.winnows[mwf.top] = NewWinnowFeature(mwf.belongsTo + "-" + string(mwf.top))
}

func (mwf *MultiWinnowFeature) Covariance(other MeasurableFeature) float64 {
	cov := 0.0

	switch other.(type) {
	case *MultiWinnowFeature:
		otherf := other.(*MultiWinnowFeature)
		mostMatch := make([]float64, len(mwf.winnows))
		mostMatchIdx := make([]int, len(mostMatch))
		for dim, tensor := range mwf.winnows {
			mostMatch[dim] = 0
			mostMatchIdx[dim] = 0

			for odim, otensor := range otherf.winnows {
				match := tensor.Covariance(otensor)
				if match > mostMatch[dim] {
					mostMatchIdx[dim] = odim
					mostMatch[dim] = match
				}
			}
		}

		total_card := 0
		for _, winnow := range mwf.winnows {
			total_card += winnow.set.Cardinality()
		}

		for i, value := range mostMatch {
			cov += value * float64(mwf.winnows[i].set.Cardinality()) / float64(total_card)
		}

		break
	}

	return cov
}

func (mwf *MultiWinnowFeature) Similarity(basedOn MeasurableFeature) float64 {
	sim := 0.0

	switch basedOn.(type) {
	case *MultiWinnowFeature:
		otherf := basedOn.(*MultiWinnowFeature)
		mostMatch := make([]float64, len(mwf.winnows))
		mostMatchIdx := make([]int, len(mostMatch))
		for dim, tensor := range mwf.winnows {
			mostMatch[dim] = 0
			mostMatchIdx[dim] = 0

			for odim, otensor := range otherf.winnows {
				match := tensor.Similarity(otensor)
				if match > mostMatch[dim] {
					mostMatchIdx[dim] = odim
					mostMatch[dim] = match
				}
			}
		}

		total_card := 0
		for _, winnow := range mwf.winnows {
			total_card += winnow.set.Cardinality()
		}

		for i, value := range mostMatch {
			sim += value * float64(mwf.winnows[i].set.Cardinality()) / float64(total_card)
		}

		break
	}

	return sim
}
