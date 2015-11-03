package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageNameFromAbsFileName(t *testing.T) {
	testData := []struct {
		filename    string
		packagename string
	}{
		{
			filename:    "/some/go/path.go",
			packagename: "go",
		},
		{
			filename:    "/some/path.go",
			packagename: "some",
		},
		{
			filename:    "/abs.go",
			packagename: "",
		},
	}
	for _, tt := range testData {
		assert.Equal(t, tt.packagename,
			packageNameFromAbsFileName(tt.filename))
	}
}
