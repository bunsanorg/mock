package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func ListInterfacesFromAST(fileAST *ast.File) []string {
	result := make([]string, 0)
	for _, decl := range fileAST.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			_, ok = ts.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}
			result = append(result, ts.Name.String())
		}
	}
	return result
}

func ListInterfacesFromFile(filename string) ([]string, error) {
	fset := token.NewFileSet()
	fileAST, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, err
	}
	return ListInterfacesFromAST(fileAST), nil
}
