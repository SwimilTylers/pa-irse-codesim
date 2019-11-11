package parser

import (
	feature2 "feature"
	"preprocess"
)

const fileBase = "../../testfile/"

func ExampleFeatureSelector_SelectFeatures() {
	p := NewParser(5, 4, 3)

	oname := "text11.ll"

	observs := preprocess.GetFromFile(fileBase+oname, preprocess.GetPreprocessFunc("func-squeeze"), false)
	feature := feature2.NewFeature(oname)

	for _, value := range observs {
		p.SelectFeatures(value, feature)
	}
}
