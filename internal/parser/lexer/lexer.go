package lexer

import (
	"unicode/utf8"

	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
	"github.com/gojisvm/gojis/internal/parser/token"
)

// Lexer is a tokenizer/scanner for ECMAScript source code. It operates on a
// byte slice and with runes. A Lexer has a token stream where it pushes all
// tokens onto. Lexer states are defined recursively. It is recommended to
// return to an initial state after every token, since token sequences should be
// a responsibility of the parser.
type Lexer struct {
	input []byte
	start int
	pos   int
	width int

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
		input: input,
		start: 0,
		pos:   0,
		width: 0,

		current: initial,
		tokens:  token.NewStream(),
	}
}

// StartLexing starts this lexer, so that it will start pusing tokens onto the
// token stream. This should be executed in a separate goroutine, as it will
// block until the lexer is done.
func (l *Lexer) StartLexing() {
	defer l.tokens.Close()

	for {
		l.current = l.current(l)
		if l.current == nil {
			// last state was end state
			break
		}
	}
}

// IsEOF determines whether the position marker of the lexer has reached the end
// of its input.
func (l *Lexer) IsEOF() bool {
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

func (l *Lexer) emit(t ...token.Type) {
	l.tokens.Push(token.New(
		t,                              // token types
		string(l.input[l.start:l.pos]), // token value
		l.start,                        // token start pos
		l.pos-l.start,                  // token length
	))
	l.start = l.pos
}

func (l *Lexer) error(msg string) {
	l.tokens.Push(token.New(
		[]token.Type{token.Error}, // error token type
		msg,                       // error message
		l.pos,                     // error position
		0,                         // length
	))
}

// next consumes the next rune from the input. When using next, always check if
// the lexer is EOF.
func (l *Lexer) next() rune {
	r, w := utf8.DecodeRune(l.input[l.pos:])
	l.width = w
	l.pos += w
	return r
}

// unread unreads the last read rune. This must only be called once per call to
// next.
func (l *Lexer) unread() {
	l.unreads++
	l.pos -= l.width
}

// unreadN unreads n bytes (NOT runes). n must be the sum of the width of all
// runes to unread.
func (l *Lexer) unreadN(n int) {
	l.pos -= n
}

// peek returns the next rune without consuming it.
func (l *Lexer) peek() rune {
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
	if m.Matches(l.next()) {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) acceptRune(r rune) bool {
	if r == l.next() {
		return true
	}
	l.unread()
	return false
}

func (l *Lexer) acceptWord(word string) bool {
	advanced := 0
	for _, r := range word {
		if !l.acceptRune(r) {
			l.unreadN(advanced)
			return false
		}
		advanced += len(string(r))
	}
	return true
}

func (l *Lexer) acceptMultiple(m matcher.M) uint {
	var matched uint
	for m.Matches(l.next()) {
		matched++

		if l.IsEOF() {
			return matched
		}
	}
	l.unread()
	return matched
}
