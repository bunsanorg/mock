package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/bunsanorg/buildutils"
)

var mockgen = flag.String("mockgen", "mockgen", "github.com/golang/mock/mockgen location")
var source = flag.String("source", "", "Source file")
var destination = flag.String("destination", "", "Destination file")
var importSelf = flag.Bool("import-self", false, "Import source package")

func main() {
	flag.Parse()

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
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
