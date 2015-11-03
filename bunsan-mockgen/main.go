package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/bunsanorg/buildutils"
)

var mockgen = flag.String("mockgen", "mockgen",
	"github.com/golang/mock/mockgen location")
var source = flag.String("source", "", "Source file")
var destination = flag.String("destination", "", "Destination file")
var importSelf = flag.Bool("import-self", false, "Import source package")
var gofile = flag.String("gofile", "",
	"Shortcut for -source=$gofile -destination=mock/$gofile")

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
	arguments := []string{
		"-source", *source,
		"-destination", *destination,
	}
	if *importSelf {
		arguments = append(arguments, "-imports", ".="+importPath)
	}
	cmd := exec.Command(*mockgen, arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
