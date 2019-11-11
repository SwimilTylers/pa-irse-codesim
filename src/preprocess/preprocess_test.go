package preprocess

import (
	"bytes"
	"strings"
	"testing"
)

const testBase = "../../testfile/"

func TestGetFromFile(t *testing.T) {
	s0 := GetFromFile(testBase+"text0.txt", nil, false)[0]
	g0 := "Hello, world"

	if strings.Compare(string(s0), g0) != 0 {
		t.Error("Content conflicts:", string(s0), "v.s", g0)
	}
}

func TestPreprocess(t *testing.T) {
	s0 := GetFromFile(testBase+"text1.ll", GetPreprocessFunc("func-raw"), false)
	g0 := [][]byte{
		GetFromFile(testBase+"text1_0.txt", nil, false)[0],
		GetFromFile(testBase+"text1_1.txt", nil, false)[0],
	}

	for i := 0; i < len(s0); i++ {
		if bytes.Compare(s0[0], g0[0]) != 0 {
			t.Error("Content conflicts I -", i)
		}
	}

	s1 := GetFromFile(testBase+"text1.ll", GetPreprocessFunc("func-no-comment"), false)[0]
	g1 := GetFromFile(testBase+"text1_3.txt", nil, false)[0]

	if strings.Compare(string(s1), string(g1)) != 0 {
		t.Error("Content conflicts II")
	}
}
