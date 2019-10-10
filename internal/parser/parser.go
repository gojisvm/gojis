package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gojisvm/gojis/internal/parser/ast"
)

// Parser parses ECMAScript code and generates an AST
// that can be evaluated by the runtime.
type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFiles(paths ...string) (*ast.Script, error) {
	var asts []*ast.Script
	for _, path := range paths {
		ast, err := p.ParseFile(path)
		if err != nil {
			return nil, err
		}
		asts = append(asts, ast)
	}
	return ast.Merge(asts...), nil
}

func (p *Parser) ParseFile(path string) (*ast.Script, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}
	return p.Parse(path, f)
}

func (p *Parser) Parse(srcName string, r io.Reader) (*ast.Script, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read all: %v", err)
	}
	return p.parse(srcName, data)
}

func (p *Parser) parse(srcName string, input []byte) (*ast.Script, error) {
	i := newIsolate(srcName, input)
	return i.parse()
}
