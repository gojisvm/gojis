package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

type isolate struct {
	sourceName string
	input      []rune
	start      int
	pos        int
}

func newIsolate(sourceName string, input []byte) *isolate {
	return &isolate{
		sourceName: sourceName,
		input:      []rune(string(input)),
	}
}

func (i *isolate) fatal(msg string) {
	line, col := i.lineAndColPos()
	panic(Error{
		msg:    msg,
		source: i.sourceName,
		line:   line,
		col:    col,
	})
}

func (i *isolate) lineAndColPos() (row, col int) {
	row = 1 // lines start at 1
	col = 1 // cols start at 1

	for idx, r := range i.input {
		if idx >= i.pos {
			break
		}
		if r == '\t' {
			col += 4
		} else {
			col++
		}
		if lineTerminator.Matches(r) {
			col = 1
			row++
		}
	}
	return
}

func (i *isolate) done() bool {
	return i.pos >= len(i.input)
}

// collect collects all accepted runes and returns them as a token. After that,
// it sets the start position to the current position.
func (i *isolate) collect(t token.Type) token.Token {
	tk := token.New(
		t,                              // token type
		string(i.input[i.start:i.pos]), // token value
		i.start,                        // token start pos
		i.pos-i.start,                  // token length
	)
	i.start = i.pos
	return tk
}

func (i *isolate) accept(t token.Type) (token.Token, bool) {
	/*
		White space code points may occur between any two tokens and at the
		start or end of input.
	*/
	for i.acceptMatcher(whiteSpace) {
	}

	if i.lexForToken(t) {
		return i.collect(t), true
	}
	return token.Token{}, false
}

func (i *isolate) acceptOneOf(ts ...token.Type) (token.Token, bool) {
	chck := i.checkpoint()

	for _, t := range ts {
		if t, ok := i.accept(t); ok {
			return t, ok
		}
		i.restore(chck)
	}
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

// negativeLookahead determines, whether the next token's type is NOT one of the
// given types. Given the parameters token.LineTerminator and token.Whitespace,
// this method will return false if the next token is either a line terminator
// or a whitespace. It will return true if the negative lookahead is successful.
// This method will not consume any tokens.
func (i *isolate) negativeLookahead(ts ...token.Type) bool {
	chck := i.checkpoint()
	defer i.restore(chck)

	return !i.acceptOneOfTypes(ts...)
}

type checkpoint struct {
	start int
	pos   int
}

func (i *isolate) checkpoint() checkpoint {
	return checkpoint{
		start: i.start,
		pos:   i.pos,
	}
}

func (i *isolate) restore(chck checkpoint) {
	i.start = chck.start
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

	script = i.parseScript()
	return
}

func (i *isolate) parseScript() *ast.Script {
	return parseScript(i, 0)
}
