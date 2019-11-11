package feature

import (
	"fingerprint"
	"math"
	"testing"
)

func TestWinnowFeature(t *testing.T) {
	a := NewWinnowFeature("test0")
	a.AddFingerprint(&fingerprint.Fingerprint{
		Hash:  100,
		Where: 8,
	})
	a.AddFingerprint(&fingerprint.Fingerprint{
		Hash:  200,
		Where: 4,
	})

	b := NewWinnowFeature("test1")
	b.AddFingerprint(&fingerprint.Fingerprint{
		Hash:  100,
		Where: 8,
	})
	b.AddFingerprint(&fingerprint.Fingerprint{
		Hash:  50,
		Where: 4,
	})

	cov, sim := a.Covariance(b), a.Similarity(b)
	if math.Abs(cov-1.0/3) >= 1e-5 {
		t.Error("Expected cov=1/3, get", sim)
	}
	if math.Abs(sim-0.5) >= 1e-5 {
		t.Error("Expect sim=1/2", sim)
	}
}
