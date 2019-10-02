package lexer

import "github.com/gojisvm/gojis/internal/parser/lexer/matcher"

// state is a definition of a state. The state returned will be the one that is
// executed next.
type state func(*Lexer) state

// Defined matchers
var (
	BraceClose   = matcher.String("}")
	BraceOpen    = matcher.String("{")
	BracketClose = matcher.String("]")
	BracketOpen  = matcher.String("[")
)

// Below this point, all lexer states are defined.

func lexScript(l *Lexer) state {
	return lexToken
}

func lexToken(l *Lexer) state {
	switch r := l.Peek(); r {
	case '}':
		return lexBraceClose
	case '{':
		return lexBraceOpen
	case ']':
		return lexBracketClose
	case '[':
		return lexBracketOpen
	default:
		// handle all cases that cannot be expressed in a switch block
	}
	return unexpectedToken
}

func lexBraceClose(l *Lexer) state {
	if !l.accept(BraceClose) {
		return unexpectedToken
	}
	return lexToken
}

func lexBraceOpen(l *Lexer) state {
	if !l.accept(BraceOpen) {
		return unexpectedToken
	}
	return lexToken
}

func lexBracketClose(l *Lexer) state {
	if !l.accept(BracketClose) {
		return unexpectedToken
	}
	return lexToken
}

func lexBracketOpen(l *Lexer) state {
	if !l.accept(BracketOpen) {
		return unexpectedToken
	}
	return lexToken
}
