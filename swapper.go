package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var verbose *bool
var cpAll *bool
var workspace *string

var wspace = "C:/Users/lirodrig/Documents/test-env/dir/*"
var backup = "C:/Users/lirodrig/Documents/test-env/back/"
var flist []string

func main() {

	parseFlags()

	validateInputs()

	files, err := filepath.Glob(wspace)

	if err != nil {
		log.Fatal(err)
	}

	for _, fpath := range files {
		copyFile(fpath)
	}
}

func parseFlags() {
	verbose = flag.Bool("v", false, "Verbose mode.")
	cpAll = flag.Bool("a", false, "Replace all files in the workspace.")
	workspace = flag.String("w", "", "Workspace directory. Required for the replace all option.")

	flag.Parse()
}

func validateInputs() {
	fmt.Printf(*workspace)

	if *cpAll && *workspace == "" {
		log.Fatal("Error. Workspace directory required when replacing all files.")
	}
}

func copyFile(fpath string) {
	fname := filepath.Base(fpath)
	ndir := backup + fname

	printInfo("=======================================")
	printInfo("Copying file:", fname, "\nfrom:", fpath, "\nto:", ndir)

	from, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}

	defer from.Close()

	to, err := os.OpenFile(ndir, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}

func printInfo(msg ...string) {
	if *verbose {
		fmt.Println(msg)
	}
}
