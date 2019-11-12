package main

import (
	"feature"
	"flag"
	"fmt"
	"log"
	"measurement"
	"os"
	"parser"
	"path/filepath"
	"preprocess"
	"syscmd"
	"time"
)

var verbose = flag.Bool("v", false, "Show progress.")
var k = flag.Int("k", 5, "Kgrams Parameter. Default to 5.")
var w = flag.Int("w", 4, "Winnow size. Default to 4.")
var hashBase = flag.Uint("b", 3, "Base of Karp-Rabin String Matching. Default to 3.")
var featureType = flag.String("ft", "winnow", "Feature Type. Default to winnow.")
var preprocessMode = flag.String("ppm", "func-squeeze", "Choose text preprocess mode. Default to func-squeeze.")
var measurementMode = flag.String("mm", "str8", "Choose similarity measurement. Default to str8.")

func progress(v ...interface{}) {
	if *verbose {
		log.Println(v...)
	}
}

func main() {
	flag.Parse()

	files, err := pathChange(flag.Args())

	if err != nil {
		log.Fatalln(err)
	}

	if !checkFileExistence(files) {
		log.Fatalln("File not found!")
	}

	progress("Code Comparison between", files)

	startTs := time.Now()

	fpParser := parser.NewParser(*k, *w, *hashBase)
	fpParser.Verbose = *verbose

	prep := preprocess.GetPreprocessFunc(*preprocessMode)

	features := extractFeatures(files, fpParser, prep)

	midTs := time.Now()

	switch *measurementMode {
	case "str8":
		fmt.Println(measurement.Straightforward(features))
		break
	default:
		log.Fatalln("Unknown measurement:", *measurementMode)
	}

	endTs := time.Now()

	progress("Extract takes", midTs.Sub(startTs), "\tComparison takes", endTs.Sub(midTs))
}

func pathChange(paths []string) ([]string, error) {
	ret := make([]string, len(paths))
	var err error

	for i, p := range paths {
		if !filepath.IsAbs(p) {
			ret[i], err = filepath.Abs(p)
		} else {
			ret[i] = p
		}
	}

	return ret, err
}

func fileExistence(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func checkFileExistence(files []string) bool {
	isExist := true
	for _, f := range files {
		ex, _ := fileExistence(f)
		isExist = isExist && ex
	}
	return isExist
}

func extractFeatures(files []string, fpParser *parser.FeatureSelector, prep func([]byte) [][]byte) []feature.MeasurableFeature {
	features := make([]feature.MeasurableFeature, len(files))

	for i, fname := range files {
		startTs := time.Now()

		progress("Dump file:", fname)
		text := syscmd.Dump(fname)

		if len(text) == 0 {
			log.Fatalln("Dump Error")
		}

		progress("Dump takes", time.Now().Sub(startTs))

		var fp feature.FlexibleFeature

		progress("Select", *featureType, "Feature for", fname, ", preprocessed by", *preprocessMode)

		switch *featureType {
		case "winnow":
			fp = feature.NewWinnowFeature(fname)
			break
		case "multi-winnow":
			fp = feature.NewMultiWinnowFeature(fname)
			break
		default:
			log.Fatalln("Unknown feature type", *featureType)
		}

		startTs = time.Now()

		for _, st := range prep(text) {
			fp.AddDimension()
			fpParser.SelectFeatures(st, fp)
		}

		progress("Winnow Takes", time.Now().Sub(startTs))

		features[i] = fp.(feature.MeasurableFeature)
	}

	return features
}
