# pa-irse-codesim: Code similarity detector
> MG1933080 Yi-Jiang Yang

This is an project assignment for NJUCS Introduction of Software Engineering Research (2019). 
Our goal is an implementation of code similarity detector. This a **Go** project. You can use Go tools for
compilation and unit test.

Our major technique is K-grams and Winnowing. We employs external github repo [golang-set](github.com/deckarep/golang-set).
Hence, you have to manually download their source code into `src` directory, or you can use IDE like Jetbrains Goland for
automatic configuration (see `src/feature/winnow.go`, follow Goland's hint and begin downloading).

## Usages & Options

You can use ``codesim -h`` for usage information.

```text
Usage: codesim [options] code1 code2.
Options can be:
  -b uint
    	Base of Karp-Rabin String Matching. (default 3)
  -ft string
    	Feature Type. Your choice can be "winnow" or "multi-winnow". (default "multi-winnow")
  -k int
    	Kgrams Parameter. (default 5)
  -mm string
    	Choose measurement. Your choice can be "max" or "mean". (default "max")
  -pm string
    	Choose text preprocess mode. Your choice can be "func-raw", "func-no-comment" or "func-squeeze". (default "func-squeeze")
  -sm string
    	Choose similarity. Your choice can be "smc", "overlap" or "jaccard". (default "jaccard")
  -v	Show progress.
  -w int
    	Winnow size. Default to 4. (default 4)
```

## Compilation & Run

We use Go 1.12 to build our project. Before compilation, you MUST deal with external github repo [golang-set](github.com/deckarep/golang-set).
We recommend you to use IDE to automatic configuration. To compile, please follow these steps:

1. open `src/feature/winnow.go` in Jetbrains Goland

2. go to File/Settings, set `GOPATH` (or double `shift` and search `GOPATH`)

3. watch the file head, find imports

4. click on line.5 (`	mapset "github.com/deckarep/golang-set"`) and `alt`+`enter`

5. Choose `'Download ...'`, and wait for completion

6. If success, you can see `src/github.com` and `pkg`. Now, you can compile the project now.

To run:

    go run src/main.go code1 code2


## Source Code Tree 

```text
src
├── feature
│   ├── feature.go
│   ├── feature_test.go
│   ├── multiwinnow.go
│   └── winnow.go
├── fingerprint
│   └── fingerprint.go
├── main.go
├── measurement
│   ├── measurement.go
│   ├── measurement_test.go
│   └── utils.go
├── parser
│   ├── parser.go
│   └── parser_test.go
├── preprocess
│   ├── getutils.go
│   ├── preprocess.go
│   └── preprocess_test.go
└── syscmd
    ├── clangdump.go
    ├── llvm-dump.sh
    ├── syscmd_test.go
    └── utils.go

```