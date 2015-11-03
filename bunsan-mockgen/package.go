package main

import (
	"path"
	"path/filepath"
)

func packageNameFromAbsFileName(filename string) string {
	pkg := path.Base(path.Dir(filename))
	if path.IsAbs(pkg) {
		return ""
	}
	return pkg
}

func PackageNameFromFileName(filename string) (string, error) {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	return packageNameFromAbsFileName(abs), nil
}
