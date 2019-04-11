//go:generate go run generate_template.go

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"sort"
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

func getIdentName(e ast.Expr) string {
	switch v := e.(type) {
	case *ast.Ident:
		return v.Name

	case *ast.StarExpr:
		return "*" + getIdentName(v.X)

	default:
		panic(fmt.Sprintf("cannot decode %T", e))
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
									return pkgName, getIdentName(t.Elt)
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

func getType(name string) string {
	switch name {
	case "int8", "uint8", "byte", "int16", "uint16", "int32", "rune", "uint32",
		"int64", "uint64", "int", "uint", "uintptr", "float32", "float64",
		"complex64", "complex128":
		return "number"

	case "string":
		return "string"
	}

	return "struct"
}

func getImports(s string) (imports []string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", s, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	for _, s := range f.Imports {
		imports = append(imports, s.Path.Value)
	}

	return
}

func getAllImports(files []string) (imports []string) {
	mapImports := map[string]struct{}{}

	for _, file := range files {
		for _, imp := range getImports(file) {
			mapImports[imp] = struct{}{}
		}
	}

	for imp := range mapImports {
		imports = append(imports, imp)
	}

	sort.Strings(imports)

	return
}

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, 0)
	check(err)

	for _, sliceType := range os.Args[1:] {
		packageName, elementType := findType(pkgs, sliceType)
		templates := []string{pieAllTemplate}

		kind := getType(elementType)

		if kind == "number" || kind == "string" {
			templates = append(templates, pieStringsNumbersTemplate)
		}

		if kind == "number" {
			templates = append(templates, pieNumbersTemplate)
		}

		// Aggregate imports.
		t := fmt.Sprintf("package %s\n\nimport (", packageName)
		for _, imp := range getAllImports(templates) {
			t += fmt.Sprintf("\n\t%s", imp)
		}
		t += "\n)\n\n"

		for _, tmpl := range templates {
			i := strings.Index(tmpl, "//")
			t += tmpl[i:] + "\n"
		}

		t = strings.Replace(t, "SliceType", sliceType, -1)
		t = strings.Replace(t, "ElementType", elementType, -1)

		switch kind {
		case "number":
			t = strings.Replace(t, "ElementZeroValue", "0", -1)

		case "string":
			t = strings.Replace(t, "ElementZeroValue", `""`, -1)

		case "struct":
			zeroValue := fmt.Sprintf("%s{}", elementType)

			// If its a pointer we need to replace '*' -> '&' when
			// instantiating.
			if elementType[0] == '*' {
				zeroValue = "&" + zeroValue[1:]
			}

			t = strings.Replace(t, "ElementZeroValue", zeroValue, -1)
		}

		// The TrimRight is important to remove an extra new line that conflicts
		// with go fmt.
		t = strings.TrimRight(t, "\n") + "\n"

		err := ioutil.WriteFile(strings.ToLower(sliceType)+"_pie.go", []byte(t), 0755)
		check(err)
	}
}
