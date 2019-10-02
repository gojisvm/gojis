package parser

import (
	"fmt"

	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/lexer"
	"github.com/gojisvm/gojis/internal/parser/token"
)

type isolate struct {
	sourceName string
	lx         *lexer.Lexer

	tokens *token.Stream
	cached bool
	cache  token.Token
	done   bool
}

func newIsolate(sourceName string, input []byte) *isolate {
	lx := lexer.New(input)
	return &isolate{
		sourceName: sourceName,
		lx:         lx,
		tokens:     lx.TokenStream(),
	}
}

func (i *isolate) fatalf(msg string, args ...interface{}) {
	i.fatal(fmt.Sprintf(msg, args...))
}

func (i *isolate) fatal(msg string) {
	panic(Error{
		msg: msg,
	})
}

func (i *isolate) peek() token.Token {
	i.ensureCache()
	return i.cache
}

func (i *isolate) next() token.Token {
	i.ensureCache()
	i.cached = false
	return i.cache
}

func (i *isolate) unread() {
	i.cached = true
}

func (i *isolate) ensureCache() {
	if i.cached {
		// noop if there's a token in cache
		return
	}

	token, ok := i.tokens.Next()
	if !ok {
		i.done = true
		return
	}

	i.cached = true
	i.cache = token
}

func (i *isolate) accept(t token.Type) bool {
	next := i.next()
	for _, typ := range next.Types {
		if typ == t {
			return true
		}
	}
	i.unread()
	return false
}

// actual parsing including all rules starts here

func (i *isolate) parse() (script *ast.Script, err error) {
	go i.lx.StartLexing()

	defer func() {
		if recovered := recover(); recovered != nil {
			parseError, ok := recovered.(Error)
			if !ok {
				panic(recovered) // re-panic if not parser error
			}
			err = parseError
		}
	}()

	script = i.parseScript()
	return
}

func (i *isolate) parseScript() *ast.Script {
	i.fatalf("Unimplemented")
	return nil
}
