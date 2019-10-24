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

	tokens []token.Token
	pos    int
}

func newIsolate(sourceName string, input []byte) *isolate {
	lx := lexer.New(input)
	return &isolate{
		sourceName: sourceName,
		lx:         lx,
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

func (i *isolate) done() bool {
	return i.pos >= len(i.tokens)
}

func (i *isolate) next() (token.Token, bool) {
	if i.done() {
		return token.Token{}, false
	}

	t := i.tokens[i.pos]
	i.pos++
	return t, true
}

func (i *isolate) unread() {
	if i.pos-1 < 0 {
		// noop, cannot unread
		return
	}

	i.pos--
}

func (i *isolate) accept(t token.Type) (token.Token, bool) {
	next, ok := i.next()
	if !ok || next.Type != t {
		i.unread()
		return token.Token{}, false
	}
	return next, true
}

func (i *isolate) acceptOneOf(ts ...token.Type) (token.Token, bool) {
	next, ok := i.next()
	if !ok {
		return token.Token{}, false
	}

	for _, t := range ts {
		if next.Type == t {
			return next, true
		}
	}

	i.unread()
	return token.Token{}, false
}

func (i *isolate) acceptOneOfTypes(ts ...token.Type) bool {
	for _, t := range ts {
		_, ok := i.accept(t)
		if ok {
			return true
		}
	}
	return false
}

type checkpoint struct {
	pos int
}

func (i *isolate) checkpoint() checkpoint {
	return checkpoint{
		pos: i.pos,
	}
}

func (i *isolate) restore(chck checkpoint) {
	i.pos = chck.pos
}

// actual parsing including all rules starts here

func (i *isolate) parse() (script *ast.Script, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			parseError, ok := recovered.(Error)
			if !ok {
				panic(recovered) // re-panic if not parser error
			}
			err = parseError
		}
	}()

	go func() {
		err := i.lx.StartLexing()
		if err != nil {
			i.fatalf("Lexer failed: %v", err)
		}
	}()

	for t := range i.lx.TokenStream().Tokens() {
		i.tokens = append(i.tokens, t)
	}

	script = i.parseScript()
	return
}

func (i *isolate) parseScript() *ast.Script {
	return parseScript(i, 0)
}
