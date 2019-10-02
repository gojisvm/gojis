package lexer

import (
	"unicode/utf8"

	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
	"github.com/gojisvm/gojis/internal/parser/token"
)

// lexer is a tokenizer/scanner for ECMAScript source code. It operates on a
// byte slice and with runes. A lexer has a token stream where it pushes all
// tokens onto. lexer states are defined recursively. It is recommended to
// return to an initial state after every token, since token sequences should be
// a responsibility of the parser.
type lexer struct {
	input []byte
	start int
	pos   int
	width int

	current state
	tokens  *token.Stream
}

// New creates a new ready to use lexer that operates on the given input bytes.
// To start pushing tokens onto its token stream, the lexer must be started with
// a call to StartLexing. It is recommended that this call is executed in a
// separate goroutine.
func New(input []byte) *lexer {
	return &lexer{
		input: input,
		start: 0,
		pos:   0,
		width: 0,

		current: lexScript,
		tokens:  token.NewStream(),
	}
}

// StartLexing starts this lexer, so that it will start pusing tokens onto the
// token stream. This should be executed in a separate goroutine, as it will
// block until the lexer is done.
func (l *lexer) StartLexing() {
	defer l.tokens.Close()

	for !l.IsEOF() {
		l.current = l.current(l)
		if l.current == nil {
			// last state was end state
			break
		}
	}
}

// IsEOF determines whether the position marker of the lexer has reached the end
// of its input.
func (l *lexer) IsEOF() bool {
	return l.pos >= len(l.input)
}

// TokenStream returns the token stream of this lexer. This is the stream that
// it will push tokens onto. The token stream is closed after StartLexing is
// done.
func (l *lexer) TokenStream() *token.Stream {
	return l.tokens
}

func (l *lexer) emit(t token.Type) {
	l.tokens.Push(token.New(
		t,                              // token type
		string(l.input[l.start:l.pos]), // token value
		l.start,                        // token start pos
		l.pos-l.start,                  // token length
	))
	l.start = l.pos
}

func (l *lexer) error(msg string) {
	l.tokens.Push(token.New(
		token.Error, // error token type
		msg,         // error message
		l.pos,       // error position
		0,           // length
	))
}

func (l *lexer) next() rune {
	r, w := utf8.DecodeRune(l.input[l.pos:])
	l.width = w
	l.pos += w
	return r
}

func (l *lexer) unread() {
	l.pos -= l.width
}

func (l *lexer) unreadN(n int) {
	l.pos -= n
}

func (l *lexer) Peek() rune {
	r := l.next()
	l.unread()
	return r
}

func (l *lexer) accept(m matcher.M) bool {
	if m.Matches(l.next()) {
		return true
	}
	l.unread()
	return false
}

func (l *lexer) acceptRune(r rune) bool {
	if r == l.next() {
		return true
	}
	l.unread()
	return false
}

func (l *lexer) acceptWord(word string) bool {
	advanced := 1
	for _, r := range word {
		if !l.acceptRune(r) {
			l.unreadN(advanced)
			return false
		}
		advanced++
	}
	return true
}

func (l *lexer) acceptMultiple(m matcher.M) uint {
	var matched uint
	for m.Matches(l.next()) {
		matched++
	}
	l.unread()
	return matched
}
