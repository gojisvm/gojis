package lexer

import "github.com/gojisvm/gojis/internal/parser/lexer/matcher"

// state is a definition of a state. The state returned will be the one that is
// executed next.
type state func(*lexer) state

// Defined matchers
var (
	BraceClose   = matcher.String("}")
	BraceOpen    = matcher.String("{")
	BracketClose = matcher.String("]")
	BracketOpen  = matcher.String("[")
)

// Below this point, all lexer states are defined.

func lexScript(l *lexer) state {
	return lexToken
}

func lexToken(l *lexer) state {
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

func lexBraceClose(l *lexer) state {
	if !l.accept(BraceClose) {
		return unexpectedToken
	}
	return lexToken
}

func lexBraceOpen(l *lexer) state {
	if !l.accept(BraceOpen) {
		return unexpectedToken
	}
	return lexToken
}

func lexBracketClose(l *lexer) state {
	if !l.accept(BracketClose) {
		return unexpectedToken
	}
	return lexToken
}

func lexBracketOpen(l *lexer) state {
	if !l.accept(BracketOpen) {
		return unexpectedToken
	}
	return lexToken
}
