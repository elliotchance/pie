//go:generate go run generate.go

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/elliotchance/pie/functions"
)

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

func getKeyAndElementType(pkg *ast.Package, name string, typeSpec *ast.TypeSpec) (string, string, string, *TypeExplorer) {
	pkgName := pkg.Name

	if t, ok := typeSpec.Type.(*ast.ArrayType); ok {
		explorer := NewTypeExplorer(pkg, getIdentName(t.Elt))

		return pkgName, "", getIdentName(t.Elt), explorer
	}

	if t, ok := typeSpec.Type.(*ast.MapType); ok {
		explorer := NewTypeExplorer(pkg, getIdentName(t.Value))

		return pkgName, getIdentName(t.Key), getIdentName(t.Value), explorer
	}

	panic(fmt.Sprintf("type %s must be a slice or map", name))
}

func findType(pkgs map[string]*ast.Package, name string) (packageName, keyType, elementType string, explorer *TypeExplorer) {
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if typeSpec.Name.String() == name {
								return getKeyAndElementType(pkg, name, typeSpec)
							}
						}
					}
				}
			}
		}
	}

	panic(fmt.Sprintf("type %s does not exist", name))
}

func getType(keyType, elementType string) int {
	if keyType != "" {
		return functions.ForMaps
	}

	switch elementType {
	case "int8", "uint8", "byte", "int16", "uint16", "int32", "rune", "uint32",
		"int64", "uint64", "int", "uint", "uintptr", "float32", "float64",
		"complex64", "complex128":
		return functions.ForNumbers

	case "string":
		return functions.ForStrings
	}

	return functions.ForStructs
}

func getImports(packageName, s string) (imports []string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", s, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	for _, s := range f.Imports {
		importName := s.Path.Value

		if importName == `"github.com/elliotchance/pie/pie"` &&
			isSelfPackage(packageName) {
			continue
		}

		imports = append(imports, importName)
	}

	return
}

func getAllImports(packageName string, files []string, explorer *TypeExplorer) (imports []string) {
	mapImports := map[string]struct{}{}

	for _, file := range files {
		if !explorer.HasString() && strings.Contains(file, "mightBeString") {
			mapImports[`"fmt"`] = struct{}{}
		}

		for _, imp := range getImports(packageName, file) {
			mapImports[imp] = struct{}{}
		}
	}

	for imp := range mapImports {
		imports = append(imports, imp)
	}

	sort.Strings(imports)

	return
}

// We have to generate imports slightly differently when we are building code
// that will go into its own packages vs an external package.
func isSelfPackage(packageName string) bool {
	return packageName == "pie"
}

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, 0)
	check(err)

	for _, arg := range os.Args[1:] {
		mapOrSliceType, fns := getFunctionsFromArg(arg)
		packageName, keyType, elementType, explorer := findType(pkgs, mapOrSliceType)
		kind := getType(keyType, elementType)

		var templates []string
		for _, function := range functions.Functions {
			if fns[0] != "*" && !stringSliceContains(fns, function.Name) {
				continue
			}

			if function.For&kind != 0 {
				templates = append(templates, pieTemplates[function.Name])
			}
		}

		// Aggregate imports.
		t := fmt.Sprintf("package %s\n\n", packageName)

		imports := getAllImports(packageName, templates, explorer)
		if len(imports) > 0 {
			t += fmt.Sprintf("import (")
			for _, imp := range imports {
				t += fmt.Sprintf("\n\t%s", imp)
			}
			t += "\n)\n\n"
		}

		for _, tmpl := range templates {
			i := strings.Index(tmpl, "//")
			t += tmpl[i:] + "\n"
		}

		t = strings.Replace(t, "StringSliceType", mapOrSliceType, -1)
		t = strings.Replace(t, "StringElementType", elementType, -1)
		t = strings.Replace(t, "ElementType", elementType, -1)
		t = strings.Replace(t, "MapType", mapOrSliceType, -1)
		t = strings.Replace(t, "KeyType", elementType, -1)
		t = strings.Replace(t, "KeySliceType", "[]"+keyType, -1)
		t = strings.Replace(t, "SliceType", mapOrSliceType, -1)

		if !explorer.HasEquals() {
			re := regexp.MustCompile(`([\w_]+|[\w_]+\[\w+\])\.Equals\(([^)]+)\)`)
			t = ReplaceAllStringSubmatchFunc(re, t, func(groups []string) string {
				return fmt.Sprintf("%s == %s", groups[1], groups[2])
			})
		}

		if !explorer.HasString() {
			t = strings.Replace(t, "mightBeString.String()", `fmt.Sprintf("%v", mightBeString)`, -1)
		}

		switch kind {
		case functions.ForNumbers:
			t = strings.Replace(t, "ElementZeroValue", "0", -1)

		case functions.ForStrings:
			t = strings.Replace(t, "ElementZeroValue", `""`, -1)

		case functions.ForStructs:
			zeroValue := fmt.Sprintf("%s{}", elementType)

			// If its a pointer we need to replace '*' -> '&' when
			// instantiating.
			if elementType[0] == '*' {
				zeroValue = "&" + zeroValue[1:]
			}

			t = strings.Replace(t, "ElementZeroValue", zeroValue, -1)
		}

		if isSelfPackage(packageName) {
			t = strings.Replace(t, "pie.Strings", "Strings", -1)
			t = strings.Replace(t, "pie.Ints", "Ints", -1)
			t = strings.Replace(t, "pie.Float64s", "Float64s", -1)
		}

		// The TrimRight is important to remove an extra new line that conflicts
		// with go fmt.
		t = strings.TrimRight(t, "\n") + "\n"

		err := ioutil.WriteFile(strings.ToLower(mapOrSliceType)+"_pie.go", []byte(t), 0755)
		check(err)
	}
}

func getFunctionsFromArg(arg string) (mapOrSliceType string, fns []string) {
	parts := strings.Split(arg, ".")

	if len(parts) < 2 {
		panic("must specify at least one function or *: " + arg)
	}

	return parts[0], parts[1:]
}

func stringSliceContains(haystack []string, needle string) bool {
	if haystack == nil {
		return false
	}

	for _, w := range haystack {
		if w == needle {
			return true
		}
	}
	return false
}

// http://elliot.land/post/go-replace-string-with-regular-expression-callback
func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		if isNegative(str[v[0]-1]) {
			result += str[lastIndex:v[0]] + addBrackets(repl(groups))
		} else {
			result += str[lastIndex:v[0]] + repl(groups)
		}
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}

func isNegative(b byte) bool {
	return b == '!'
}

func addBrackets(str string) string {
	return `(` + str + `)`
}
