//go:generate go run generate_template.go

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

type ElementType float64
type SliceType []ElementType
var ElementZeroValue ElementType

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func findType(pkgs map[string]*ast.Package, name string) (packageName, elementType string) {
	for pkgName, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if typeSpec.Name.String() == name {
								if t, ok := typeSpec.Type.(*ast.ArrayType); ok {
									return pkgName, t.Elt.(*ast.Ident).Name
								} else {
									panic(fmt.Sprintf("type %s must be a slice", name))
								}
							}
						}
					}
				}
			}
		}
	}

	panic(fmt.Sprintf("type %s does not exist", name))
}

func isNumeric(name string) bool {
	switch name {
	case "int8", "uint8", "byte", "int16", "uint16", "int32", "rune", "uint32",
		"int64", "uint64", "int", "uint", "uintptr", "float32", "float64",
		"complex64", "complex128":
		return true
	}

	return false
}

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, 0)
	check(err)

	for _, sliceType := range os.Args[1:] {
		packageName, elementType := findType(pkgs, sliceType)
		t := pieTemplate

		numeric := isNumeric(elementType)
		if !numeric {
			t = strings.Split(t, "// ---")[0]
		}

		t = strings.Replace(t, "package main", "package "+packageName, -1)
		t = strings.Replace(t, "SliceType", sliceType, -1)
		t = strings.Replace(t, "ElementType", elementType, -1)
		t = strings.Replace(t, "ElementConditionFunc", sliceType+"ConditionFunc", -1)
		t = strings.Replace(t, "ElementTransformFunc", sliceType+"TransformFunc", -1)

		if numeric {
			t = strings.Replace(t, "ElementZeroValue", "0", -1)
		} else {
			t = strings.Replace(t, "ElementZeroValue", `""`, -1)
		}

		err := ioutil.WriteFile(strings.ToLower(sliceType)+"_pie.go", []byte(t), 0755)
		check(err)
	}
}
