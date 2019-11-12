package measurement

import (
	"feature"
	"log"
	"os"
	"parser"
	"path/filepath"
	"preprocess"
	"syscmd"
	"testing"
)

func fetchTexts(folder string) [][]byte {
	currentPwd, _ := os.Getwd()
	currentPwd = filepath.Clean(currentPwd + "/../../testfile/" + folder)

	bs := make([][]byte, 0, 10)

	if err := filepath.Walk(currentPwd, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		l := len(bs)
		l++

		if l == cap(bs) {
			newBs := make([][]byte, l+1, 2*l)
			copy(newBs, bs)
			bs = newBs
		} else {
			bs = bs[:l+1]
		}

		bs[l] = syscmd.ClangDump(path)

		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	return bs
}

func extractFeature(bs [][]byte, fpParser *parser.FeatureSelector, prep func([]byte) [][]byte,
	fvector func() feature.FlexibleFeature, measurement func(feature.FlexibleFeature) feature.MeasurableFeature) []feature.MeasurableFeature {
	ret := make([]feature.MeasurableFeature, len(bs))

	for i, b := range bs {
		vec := fvector()
		for _, value := range prep(b) {
			vec.AddDimension()
			fpParser.SelectFeatures(value, vec)
		}
		ret[i] = measurement(vec)
	}

	return ret
}

func TestStraightforward(t *testing.T) {
	fp := parser.NewParser(5, 4, 3)

	features := extractFeature(fetchTexts("comparison_0"), fp, preprocess.GetPreprocessFunc("func-squeeze"), func() feature.FlexibleFeature {
		return feature.NewMultiWinnowFeature("test")
	}, func(flexibleFeature feature.FlexibleFeature) feature.MeasurableFeature {
		return flexibleFeature.(feature.MeasurableFeature)
	})

	t.Log(features)
}
