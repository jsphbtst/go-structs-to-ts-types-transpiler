package mapper

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/jsphbtst/go-to-ts-transpiler/pkg/utils"
)

func tsTypeFromGoType(
	expr ast.Expr,
	typeMapping *map[string]string,
) string {
	switch t := expr.(type) {
	case *ast.Ident:
		// Simple types
		tsType, ok := (*typeMapping)[t.Name]
		if !ok {
			tsType = "any"
		}
		return tsType
	case *ast.ArrayType:
		// Array types; recursive call
		innerType := tsTypeFromGoType(t.Elt, typeMapping)
		return fmt.Sprintf("%s[]", innerType)
	}

	return "any"
}

func GenerateTypeScriptInterface(
	structName string,
	fields []*ast.Field,
	typeMapping *map[string]string,
) string {
	var tsCode strings.Builder

	fmt.Fprintf(&tsCode, "type %s = {\n", structName)

	for _, field := range fields {
		fieldName := field.Names[0].Name
		tsType := tsTypeFromGoType(field.Type, typeMapping)

		// Convert fieldName to camelCase
		camelCaseFieldName := utils.PascalToCamelCase(fieldName)

		fmt.Fprintf(&tsCode, "  %s: %s\n", camelCaseFieldName, tsType)
	}

	tsCode.WriteString("}\n")
	return tsCode.String()
}
