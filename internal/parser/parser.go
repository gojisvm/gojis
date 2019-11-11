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

// New creates a new ready to use parser.
func New() *Parser {
	return &Parser{}
}

// ParseFiles parses all given files. The file must exist and must be a file
// (not a directory). Effectively, all files will be parsed with
// parser.Parser.ParseFile, and the resulting ASTs will be merged.
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

// ParseFile parses the given file and returns the resulting AST.
func (p *Parser) ParseFile(path string) (*ast.Script, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}
	return p.Parse(path, f)
}

// Parse parses frmo the given reader, which will be handled with the given
// srcName as source name. If this method is used to parse a file, srcName
// should be the file name.
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
