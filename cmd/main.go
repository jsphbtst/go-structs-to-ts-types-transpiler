package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/jsphbtst/go-to-ts-transpiler/pkg/mapper"
)

var typeMapping = map[string]string{
	"int":       "number",
	"int8":      "number",
	"int16":     "number",
	"int32":     "number",
	"int64":     "number",
	"uint":      "number",
	"uint8":     "number",
	"uint16":    "number",
	"uint32":    "number",
	"uint64":    "number",
	"float32":   "number",
	"float64":   "number",
	"bool":      "boolean",
	"string":    "string",
	"time.Time": "Date",
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	data, err := os.ReadFile(pwd + "/cmd/structs.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	src := string(data)
	tokenFileset := token.NewFileSet()

	// In case the file doesn't have a top-level package declaration, append it
	file, err := parser.ParseFile(tokenFileset, "", src, 0)
	if err != nil {
		srcWithPackage := "package main\n" + src
		file, err = parser.ParseFile(tokenFileset, "", srcWithPackage, 0)
		if err != nil {
			fmt.Println("Error parsing source code: ", err)
			return
		}
	}

	packageName := file.Name.Name
	fmt.Printf("// %s.go\n", packageName)
	fmt.Println(src)

	var typeScriptCode strings.Builder
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.TypeSpec:
			structName := node.Name.Name
			typeMapping[structName] = structName
			if structType, ok := node.Type.(*ast.StructType); ok {
				tsInterface := mapper.GenerateTypeScriptInterface(
					structName,
					structType.Fields.List,
					&typeMapping,
				)
				typeScriptCode.WriteString(tsInterface)
				typeScriptCode.WriteString("\n")
			}
		}
		return true
	})

	fmt.Printf("// %s.ts\n", packageName)
	fmt.Println(typeScriptCode.String())
}
