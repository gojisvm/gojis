package lexer

import (
	"fmt"

	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
	"github.com/gojisvm/gojis/internal/parser/token"
)

const eofRune = '\ufffd'

// Lexer is a tokenizer/scanner for ECMAScript source code. It operates on a
// byte slice and with runes. A Lexer has a token stream where it pushes all
// tokens onto. Lexer states are defined recursively. It is recommended to
// return to an initial state after every token, since token sequences should be
// a responsibility of the parser.
type Lexer struct {
	input []rune
	start int
	pos   int

	current state
	tokens  *token.Stream

	unreads uint64
}

// New creates a new ready to use lexer that operates on the given input bytes.
// To start pushing tokens onto its token stream, the lexer must be started with
// a call to StartLexing. It is recommended that this call is executed in a
// separate goroutine.
func New(input []byte) *Lexer {
	return newWithInitialState(input, lexScript)
}

func newWithInitialState(input []byte, initial state) *Lexer {
	return &Lexer{
		input: []rune(string(input)),
		start: 0,
		pos:   0,

		current: initial,
		tokens:  token.NewStream(),
	}
}

func (l *Lexer) StartLexing() (err error) {
	defer l.tokens.Close()

	defer func() {
		if recovered := recover(); recovered != nil {
			parseError, ok := recovered.(Error)
			if !ok {
				panic(recovered) // re-panic if not lexer error
			}
			err = parseError
		}
	}()

	for {
		l.current = l.current(l)
		if l.current == nil {
			// last state was end state
			break
		}
	}

	return
}

// eof determines whether the position marker of the lexer has reached the end
// of its input.
func (l *Lexer) eof() bool {
	return l.pos >= len(l.input)
}

// TokenStream returns the token stream of this lexer. This is the stream that
// it will push tokens onto. The token stream is closed after StartLexing is
// done.
func (l *Lexer) TokenStream() *token.Stream {
	return l.tokens
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) emit(t token.Type) {
	l.tokens.Push(token.New(
		t,                              // token types
		string(l.input[l.start:l.pos]), // token value
		l.start,                        // token start pos
		l.pos-l.start,                  // token length
	))
	l.start = l.pos
}

func (l *Lexer) fatalf(format string, args ...interface{}) {
	l.fatal(fmt.Sprintf(format, args...))
}

func (l *Lexer) fatal(msg string) {
	panic(Error{
		msg: msg,
	})
}

// next consumes the next rune from the input. When using next, always check if
// the lexer is EOF.
func (l *Lexer) next() rune {
	r := l.input[l.pos]
	l.pos++
	return r
}

// unread unreads one rune.
func (l *Lexer) unread() {
	l.unreads++
	l.pos--
}

// unreadN unreads n runes.
func (l *Lexer) unreadN(n int) {
	l.unreads++
	l.pos -= n
}

// peek returns the next rune without consuming it. When l.IsEOF(),
// lexer.eofRune will be returned.
func (l *Lexer) peek() rune {
	if l.eof() {
		return eofRune
	}

	r := l.next()
	l.unread()
	return r
}

// acceptEnclosed enters the given state on the stack, and returns after it was
// processed. This function is useful when having states that consist of other
// states, e.g. a string enclosed in quotes.
//
//	if !l.accept(Quote){
//		return tokenMismatch(Quote)
//	}
//	if ok, errState := l.acceptEnclosed(StringCharacters); !ok {
//		return errState
//	}
//	if !l.accept(Quote){
//		return tokenMismatch(Quote)
//	}
//
// In some cases, this removes the need to store state inside the lexer itself.
func (l *Lexer) acceptEnclosed(enclosed state) (ok bool, err state) {
	err = enclosed(l)
	ok = err == nil
	return
}

func (l *Lexer) accept(m matcher.M) bool {
	if l.eof() {
		return false
	}

	if m.Matches(l.next()) {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) acceptOneOf(ms ...matcher.M) bool {
	if l.eof() {
		return false
	}

	for _, m := range ms {
		if l.accept(m) {
			return true
		}
	}
	return false
}

func (l *Lexer) acceptRune(r rune) bool {
	if l.eof() {
		return false
	}

	if r == l.next() {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) acceptWord(word string) bool {
	if l.eof() {
		return false
	}

	for i, r := range word {
		if !l.acceptRune(r) {
			l.unreadN(i)
			return false
		}
	}
	return true
}

func (l *Lexer) acceptMultiple(m matcher.M) uint {
	if l.eof() {
		return 0
	}

	var matched uint
	for !l.eof() && m.Matches(l.next()) {
		matched++

		if l.eof() {
			return matched
		}
	}
	l.unread()
	return matched
}
