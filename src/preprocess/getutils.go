package preprocess

import (
	"io/ioutil"
	"log"
	"os"
)

func progress(verbose bool, s ...interface{}) {
	if verbose {
		log.Println(s...)
	}
}

func GetFromFile(file string, preprocess func([]byte) [][]byte, verbose bool) [][]byte {
	f, err := os.Open(file)

	if err != nil {
		log.Fatalf("Unable to open '%s': %s\n", file, err)
	}

	progress(verbose, "Read From:", file)

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Unable to close '%s': %s\n", file, err)
		}
	}()

	if contents, err := ioutil.ReadAll(f); err == nil {
		progress(verbose, "Get Contents: ", file)
		if preprocess == nil {
			return [][]byte{contents}
		} else {
			return preprocess(contents)
		}
	} else {
		log.Fatalln("Read Error:", err)
		return nil
	}
}
