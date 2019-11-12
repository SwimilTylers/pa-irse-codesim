package main

import (
	"feature"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"measurement"
	"os/exec"
	"parser"
	"preprocess"
	"time"
)

var verbose = flag.Bool("v", false, "Show progress.")
var k = flag.Int("k", 5, "Kgrams Parameter. Default to 5.")
var w = flag.Int("w", 4, "Winnow size. Default to 4.")
var hashBase = flag.Uint("b", 3, "Base of Karp-Rabin String Matching. Default to 3.")
var featureType = flag.String("-ft", "winnow", "Feature Type. Default to winnow.")
var preprocessMode = flag.String("-ppm", "func-squeeze", "Choose text preprocess mode. Default to func-squeeze.")
var measurementMode = flag.String("-mm", "str8", "Choose similarity measurement. Default to str8.")
var files = []string{
	"/home/swimiltylers/demo1.cpp",
	"/home/swimiltylers/demo2.cpp",
}

func progress(v ...interface{}) {
	if *verbose {
		log.Println(v...)
	}
}

func main() {
	flag.Parse()

	startTs := time.Now()

	fpParser := parser.NewParser(*k, *w, *hashBase)
	fpParser.Verbose = *verbose

	prep := preprocess.GetPreprocessFunc(*preprocessMode)

	features := extractFeatures(fpParser, prep)

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

func extractFeatures(fpParser *parser.FeatureSelector, prep func([]byte) [][]byte) []feature.MeasurableFeature {
	features := make([]feature.MeasurableFeature, len(files))

	for i, fname := range files {
		startTs := time.Now()
		text := dump(fname)
		progress("Dump takes", time.Now().Sub(startTs))

		var fp feature.FlexibleFeature

		progress("Select", *featureType, "Feature for", fname, ", preprocessed by", *preprocessMode)

		switch *featureType {
		case "winnow":
			fp = feature.NewWinnowFeature(fname)
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

func dump(from string) []byte {
	progress("dump", from)

	r := exec.Command("/bin/bash", "-c", fmt.Sprintf("/bin/bash llvm-dump.sh %s", from))
	f, err := r.StdoutPipe()
	if err != nil {
		log.Fatalln("Cannot get stdout:", err)
	}

	var bs []byte

	if err = r.Start(); err == nil {
		bs, err = ioutil.ReadAll(f)
		if err != nil {
			log.Fatalln("Dump Read Error:", err)
		}
	} else {
		log.Fatalln("Dump Start Error:", err)
	}

	if err = r.Wait(); err != nil {
		log.Fatalln("Dump Wait Error:", err)
	}

	return bs
}
