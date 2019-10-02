package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gojisvm/gojis/internal/parser/ast"
)

// Parser parses ECMAScript code and generates an AST
// that can be evaluated by the runtime.
type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFile(path string) (ast.ParseNode, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %v", err)
	}
	return p.Parse(path, f)
}

func (p *Parser) Parse(srcName string, r io.Reader) (ast.ParseNode, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read all: %v", err)
	}
	return p.parse(srcName, data)
}

func (p *Parser) parse(srcName string, input []byte) (ast.ParseNode, error) {
	i := newIsolate(srcName, input)
	return i.parse()
}
