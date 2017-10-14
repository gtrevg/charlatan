package main // import "github.com/percolate/charlatan"

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type stringSliceValue []string

func (s *stringSliceValue) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func (s *stringSliceValue) String() string {
	return strings.Join(*s, ", ")
}

func (s *stringSliceValue) Get() interface{} {
	return []string(*s)
}

const (
	usageFormat = `charlatan
https://github.com/percolate/charlatan

Usage:
  charlatan [options] <interface> ...
  charlatan -h | --help

Options:
`
)

var (
	outputPath    = flag.String("output", "", "output file path [default: ./charlatan.go]")
	outputPackage = flag.String("package", "", "output package name [default: \"<current package>\"]")
	dirName       = flag.String("dir", "", "input package directory [default: current package directory]")
	fileNames     stringSliceValue
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("charlatan: ")
	flag.Usage = usage
	flag.Var(&fileNames, "file", "name of input file, may be repeated, ignored if -dir is present")
}

func usage() {
	fmt.Fprintf(os.Stderr, usageFormat)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Print("interface parameters are required")
		flag.Usage()
		os.Exit(2)
	}

	if *outputPath != "" && !strings.HasSuffix(*outputPath, ".go") {
		log.Print("output path must be a Go source file name")
		flag.Usage()
		os.Exit(2)
	}

	var (
		dir string
		g   Generator
	)

	g.targetPackage = *outputPackage
	g.interfaces = flag.Args()

	if *dirName != "" {
		dir = *dirName
		g.parsePackageDir(dir)
	} else if len(fileNames) != 0 {
		dir = filepath.Dir(fileNames[0])
		for _, name := range fileNames[1:] {
			if dir != filepath.Dir(name) {
				log.Fatal("all input source files must be in the same package directory")
			}
		}
		g.parsePackageFiles(fileNames)
	} else {
		// process the package in current directory.
		dir = *dirName
		g.parsePackageDir(".")
	}

	src := g.generate()

	if *outputPath == "" {
		*outputPath = "charlatan.go"
	}

	if err := os.MkdirAll(filepath.Dir(*outputPath), 0755); err != nil {
		log.Fatalf("writing output: %s", err)
	}

	if err := ioutil.WriteFile(*outputPath, src, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
