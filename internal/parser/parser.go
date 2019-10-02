package parser

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Parser parses ECMAScript code and generates an AST
// that can be evaluated by the runtime.
type Parser struct {
	ast interface{}
}

// New creates a new ready-to-use parser.
func New() *Parser {
	p := new(Parser)
	// p.ast = NewEmptyAst()
	return p
}

// ParseFiles parses all given files and merges the produced ASTs.
// ParseFiles can produce one error per file, but these errors
// may be collections of multiple parse errors in a file.
func (p *Parser) ParseFiles(paths ...string) (errs []error) {
	for _, path := range paths {
		err := p.ParseFile(path)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return
}

// ParseFile parses the file at the given path.
// If parse errors occurred, they are collected and returned
// as a single error.
func (p *Parser) ParseFile(path string) error {
	input, err := antlr.NewFileStream(path)
	if err != nil {
		return fmt.Errorf("Error while loading file: %v", err)
	}

	_ = input

	// lexer := NewECMAScriptLexer(input)
	// lexer.RemoveErrorListeners()

	// stream := antlr.NewCommonTokenStream(lexer, 0)
	// par := NewECMAScriptParser(stream)
	// par.RemoveErrorListeners()

	// errorCollector := NewCollectingErrorListener()
	// par.AddErrorListener(errorCollector)
	// par.BuildParseTrees = true

	// tree := par.Program()
	// if errs, hasErrors := errorCollector.Errors(); hasErrors {
	// 	return NewError(path, errs...)
	// }

	// // only append root if no errors occurred while parsing
	// p.ast.AddRoot(path, tree)

	panic("TODO")
}

// AST returns the AST of all files that the parser has parsed.
func (p *Parser) AST() interface{} {
	return p.ast
}

// Error summarizes all parse errors of a single file
// and associates these errors with the parsed file.
type Error struct {
	file string
	errs []error
}

// NewError creates a new parser error with the given file and errors.
func NewError(file string, errs ...error) Error {
	return Error{
		file: file,
		errs: errs,
	}
}

func (e Error) Error() string {
	if len(e.errs) == 0 {
		return ""
	}

	var buf bytes.Buffer
	_, err := buf.WriteString(fmt.Sprintf("Errors while parsing '%v':", e.file))
	if err != nil {
		// this cannot happen as we use a bytes.Buffer, which never returns an error
		panic(err)
	}

	for i, err := range e.errs {
		_, err := buf.WriteString("\n\t" + strconv.Itoa(i+1) + ") " + err.Error())
		if err != nil {
			// this cannot happen as we use a bytes.Buffer, which never returns an error
			panic(err)
		}
	}

	return buf.String()
}
