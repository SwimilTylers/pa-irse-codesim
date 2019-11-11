package parser

import (
	"feature"
	"fingerprint"
	"log"
	"math"
)

type FeatureSelector struct {
	k int // k-grams
	w int // w-winnow window

	base            uint
	compensateConst uint

	Verbose bool
}

func NewParser(k int, w int, base uint) *FeatureSelector {
	p := &FeatureSelector{
		k:               k,
		w:               w,
		base:            base,
		compensateConst: uint(math.Pow(float64(base), float64(k))),
		Verbose:         false,
	}

	return p
}

func (fs *FeatureSelector) Status(s ...interface{}) {
	if fs.Verbose {
		log.Println(s...)
	}
}

func (fs *FeatureSelector) SelectFeatures(s []byte, out feature.FlexibleFeature) {
	// first step: k-grams & hashing
	fs.Status("I: k-grams & hashing ...")

	hash := make([]uint, len(s)-fs.k+1)

	firstShingle := s[0:fs.k]
	//log.Println("First Shingle:", string(firstShingle))

	var compensate uint

	for _, value := range firstShingle {
		hash[0] *= fs.base
		hash[0] += uint(value)
	}

	compensate = fs.compensateConst * uint(firstShingle[0])

	for i := 1; i+fs.k <= len(s); i++ {
		shingle := s[i : i+fs.k]
		//log.Println("Next Shingle:", string(shingle))
		hash[i] = hash[i-1]*fs.base - compensate + uint(shingle[fs.k-1])
		compensate = fs.compensateConst * uint(shingle[0])
	}

	//log.Println("Hash:", hash)

	// second step: winnowing
	fs.Status("II: winnowing ...")
	var lastLeastIndex int = 0

	for key, value := range hash[1:fs.w] {
		if hash[lastLeastIndex] >= value {
			lastLeastIndex = key
		}
	}

	out.AddFingerprint(&fingerprint.Fingerprint{
		Hash:  hash[lastLeastIndex],
		Where: lastLeastIndex,
	})

	for i := 1; i+fs.w <= len(hash); i++ {
		if lastLeastIndex < i {
			lastLeastIndex = i
			for key, value := range hash[i+1 : i+fs.w] {
				if hash[lastLeastIndex] >= value {
					lastLeastIndex = i + 1 + key
				}
			}
			out.AddFingerprint(&fingerprint.Fingerprint{
				Hash:  hash[lastLeastIndex],
				Where: lastLeastIndex,
			})
		} else if hash[lastLeastIndex] >= hash[i+fs.w-1] {
			lastLeastIndex = i + fs.w - 1
			out.AddFingerprint(&fingerprint.Fingerprint{
				Hash:  hash[lastLeastIndex],
				Where: lastLeastIndex,
			})
		}
	}
}
