package main

import (
	"path"
	"testing"

	"github.com/bunsanorg/buildutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	SelfPackage = "github.com/bunsanorg/mock/bunsan-mockgen"
	SelfFile    = "interfaces_test.go"
)

type FirstInterface interface{}
type SecondInterface interface{}

func TestListInterfaces(t *testing.T) {
	src, err := buildutils.SrcDir(SelfPackage)
	require.NoError(t, err)
	list, err := ListInterfacesFromFile(path.Join(src, SelfFile))
	require.NoError(t, err)
	assert.EqualValues(t, []string{"FirstInterface", "SecondInterface"}, list)
}
