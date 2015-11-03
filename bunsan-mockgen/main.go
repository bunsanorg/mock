package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/bunsanorg/buildutils"
)

var mockgen = flag.String("mockgen", "mockgen",
	"github.com/golang/mock/mockgen location")
var source = flag.String("source", "", "Source file")
var destination = flag.String("destination", "", "Destination file")
var gofile = flag.String("gofile", "",
	"Shortcut for -source=$gofile -destination=mock/$gofile")
var packageName = flag.String("package", "", "Package of the generated code; "+
	"defaults to the package of the input with a 'mock_' prefix.")

func main() {
	flag.Parse()

	if *gofile != "" {
		*source = *gofile
		*destination = path.Join("mock", *gofile)
	}
	if *source == "" {
		log.Fatal("-source is required")
	}
	if *destination == "" {
		log.Fatal("-destination is required")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	importPath, err := buildutils.ImportPath(cwd)
	if err != nil {
		log.Fatal(err)
	}
	interfaces, err := ListInterfacesFromFile(*source)
	if err != nil {
		log.Fatal(err)
	}
	arguments := []string{
		"-destination=" + *destination,
		"-package=" + *packageName,
		importPath,
		strings.Join(interfaces, ","),
	}
	cmd := exec.Command(*mockgen, arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
