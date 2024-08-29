/*
Copyright © 2024 - miyamo2 <miyamo2@outlook.com>
*/
package internal

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"os"
	"reflect"
	"slices"
	"strings"
	"text/template"
)

var (
	//go:embed templates/filtgen.gotpl
	tplStr string
	tpl    *template.Template
)

func init() {
	tpl = template.Must(template.New("filt").
		Funcs(
			template.FuncMap{
				"toLowerCamel":               strcase.ToLowerCamel,
				"joinPackageNameAndTypeName": joinPackageNameAndTypeName,
				"importToPath":               importToPath,
				"version":                    func() string { return fmt.Sprintf("@%s", Version) },
			}).
		Parse(tplStr))
}

// Output represents the output of the generation.
type Output struct {
	Package string
	Imports []Import
	Structs []Struct
}

// Import represents an import statement.
type Import struct {
	Path  string
	Alias string
}

// Struct represents the meta of the struct that is the source of the iterator to be generated.
type Struct struct {
	Name   string
	Fields []Field
}

// Field represents the meta of the field that is the source of the iterator to be generated.
type Field struct {
	Name    string
	Package string
	Type    string
	Eq      bool
	Ne      bool
	Gt      bool
	Lt      bool
	Ge      bool
	Le      bool
	Matches bool
	Is      bool
	Isnt    bool
	As      bool
	Asnt    bool
}

// Generate generates the iterator based on the source file.
func Generate(ctx context.Context, sourceFileName string) error {
	fileName := fileNameToBeGen(ctx, sourceFileName)
	fmt.Printf("generating %s from %s...\n", fileName, sourceFileName)
	nodes, err := parser.ParseFile(token.NewFileSet(), sourceFileName, nil, 0)
	if err != nil {
		fmt.Printf("failed to generate. file: %s\n", fileName)
		return err
	}
	imports := nodes.Imports

	output := Output{
		Package: nodes.Name.Name,
		Imports: append(make([]Import, 0, len(imports)), Import{Path: "iter"}),
		Structs: make([]Struct, 0),
	}

	astutil.Apply(nodes, nil, func(c *astutil.Cursor) bool {
		typeSpec, ok := c.Node().(*ast.TypeSpec)
		if !ok {
			return true
		}
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		strct := Struct{
			Name:   typeSpec.Name.Name,
			Fields: make([]Field, 0, len(structType.Fields.List)),
		}
		for _, field := range structType.Fields.List {
			switch ident := field.Type.(type) {
			case *ast.Ident:
				fieldType := ident.Name
				switch fieldType {
				case "string":
					output.Imports = addImport(output.Imports, "strings", "")
				case "error":
					output.Imports = addImport(output.Imports, "errors", "")
				}
				for _, name := range field.Names {
					tag := field.Tag
					var tagStr string
					if tag != nil {
						tagStr = tag.Value
					}
					strct.Fields = append(strct.Fields, newField(ctx, name.Name, "", fieldType, tagStr))
				}
			case *ast.SelectorExpr:
				pkg := ident.X.(*ast.Ident).Name
				fieldType := ident.Sel.Name

				if impPath := getImportPath(ctx, imports, pkg); impPath != "" {
					var alias string
					if !strings.HasSuffix(impPath, pkg) {
						alias = pkg
					}
					output.Imports = addImport(output.Imports, impPath, alias)
				}

				for _, name := range field.Names {
					tag := field.Tag
					var tagStr string
					if tag != nil {
						tagStr = tag.Value
					}
					strct.Fields = append(strct.Fields, newField(ctx, name.Name, pkg, fieldType, tagStr))
				}
			}
		}
		output.Structs = append(output.Structs, strct)
		return true
	})

	f, err := os.Create(fileName)
	err = tpl.Execute(f, output)
	if err != nil {
		fmt.Printf("failed to generate. file: %s\n", fileName)
		return err
	}
	err = f.Close()
	if err != nil {
		fmt.Printf("failed to generate. file: %s\n", fileName)
		return err
	}
	fmt.Printf("generation completed. file: %s\n", fileName)
	return nil
}

// getImportPath returns the import path of the package from the source's import statement.
func getImportPath(_ context.Context, imports []*ast.ImportSpec, pkg string) string {
	impIdx := slices.IndexFunc(imports, func(imp *ast.ImportSpec) bool {
		if imp.Name == nil {
			return false
		}
		return imp.Name.Name == pkg
	})

	if impIdx == -1 {
		impIdx = slices.IndexFunc(imports, func(imp *ast.ImportSpec) bool {
			return strings.HasSuffix(imp.Path.Value, fmt.Sprintf(`%s"`, pkg))
		})
	}
	if impIdx == -1 {
		return ""
	}
	if imports[impIdx].Path == nil {
		return ""
	}

	impPath := imports[impIdx].Path.Value
	impPath, _ = strings.CutSuffix(impPath, `"`)
	impPath, _ = strings.CutPrefix(impPath, `"`)
	return impPath
}

// addImport adds an import statement.
func addImport(imports []Import, impPath, alias string) []Import {
	if slices.ContainsFunc(imports, func(i Import) bool {
		return i.Path == impPath
	}) {
		return imports
	}
	return append(imports, Import{
		Path:  impPath,
		Alias: alias,
	})
}

// newField returns a new Field.
func newField(_ context.Context, name, pkgName, fieldType, tag string) Field {
	field := Field{
		Name:    name,
		Type:    fieldType,
		Package: pkgName,
	}

	tag, _ = strings.CutSuffix(tag, "`")
	tag, _ = strings.CutPrefix(tag, "`")

	structTag := reflect.StructTag(tag)
	filtgenTag := structTag.Get("filtgen")

	filters := strings.Split(filtgenTag, ",")
FiltersRange:
	for _, filter := range filters {
		filter = strings.TrimSpace(filter)
		switch filter {
		case "eq":
			field.Eq = true
		case "ne":
			field.Ne = true
		case "gt":
			field.Gt = true
		case "lt":
			field.Lt = true
		case "ge":
			field.Ge = true
		case "le":
			field.Le = true
		case "matches":
			field.Matches = true
		case "is":
			field.Is = true
		case "isnt":
			field.Isnt = true
		case "as":
			field.As = true
		case "asnt":
			field.Asnt = true
		case "*":
			field.Eq = true
			field.Ne = true
			field.Gt = true
			field.Lt = true
			field.Ge = true
			field.Le = true
			field.Matches = true
			field.Is = true
			field.Isnt = true
			field.As = true
			field.Asnt = true
			break FiltersRange
		}
	}
	return field
}

// fileNameToBeGen returns the file name to be generated.
func fileNameToBeGen(_ context.Context, sourceFileName string) string {
	fileNamePrefix, _ := strings.CutSuffix(sourceFileName, ".go")
	return fmt.Sprintf("%s_filtgen.go", fileNamePrefix)
}

// joinPackageNameAndTypeName joins the package name and the type name.
func joinPackageNameAndTypeName(pkg, tp string) string {
	if pkg == "" {
		return tp
	}
	return fmt.Sprintf("%s.%s", pkg, tp)
}

// importToPath returns the import statement for template.
func importToPath(imp Import) string {
	if imp.Alias != "" {
		return fmt.Sprintf(`%s "%s"`, imp.Alias, imp.Path)
	}
	return fmt.Sprintf(`"%s"`, imp.Path)
}
