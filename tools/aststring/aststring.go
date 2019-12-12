package main // import "github.com/gojisvm/gojis/tools/aststring"

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("aststring: ")

	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("expecting one argument, got %v (%v)", len(args), args)
	}

	matches, err := filepath.Glob(args[0] + "/*")
	if err != nil {
		log.Fatalf("cannot resolve directory '%v'", args[0])
	}

	for _, match := range matches {
		start := time.Now()
		process(match)
		log.Printf("%-20v (%v)", match, time.Since(start).Round(time.Millisecond))
	}
}

type structDecl struct {
	name       string
	structType *ast.StructType
}

func process(path string) {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, path, nil, parser.AllErrors)
	if err != nil {
		log.Printf("parse file: %v", err)
		return // abort, but don't kill the tool completely
	}

	var structs []*structDecl

	// gather all types
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						structs = append(structs, &structDecl{
							name:       typeSpec.Name.Name,
							structType: structType,
						})
					}
				}
			}
		}
	}

	generate(filepath.Base(path), structs)
}
