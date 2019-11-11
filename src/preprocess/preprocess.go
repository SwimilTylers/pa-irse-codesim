package preprocess

import (
	"log"
	"regexp"
	"strings"
)

var funcBlockRegex = regexp.MustCompile(`(?mU)^define .* \{(.|\n)*\}$`)
var commentRegex = regexp.MustCompile(`(?mU);.*$`)

func toSingleLine(a string) string {
	return strings.ReplaceAll(a, "\n\t", " ")
}

func findFuncBlocks(a string) []string {
	return funcBlockRegex.FindAllString(a, -1)
}

func removeComments(a string) string {
	return commentRegex.ReplaceAllString(a, "")
}

func preprocess(bs []byte, mode string) [][]byte {
	if mode == "raw" {
		return [][]byte{bs}
	}

	funcs := findFuncBlocks(string(bs))
	bfuncs := make([][]byte, len(funcs))

	for i, s := range funcs {
		switch mode {
		case "func-raw":
			bfuncs[i] = []byte(s)
			break
		case "func-no-comment":
			bfuncs[i] = []byte(removeComments(s))
			break
		case "func-squeeze":
			bfuncs[i] = []byte(toSingleLine(removeComments(s)))
			break
		default:
			log.Fatalln("Unidentified mode:", mode)
		}
	}

	return bfuncs
}

func GetPreprocessFunc(mode string) func([]byte) [][]byte {
	return func(bytes []byte) [][]byte {
		return preprocess(bytes, mode)
	}
}
