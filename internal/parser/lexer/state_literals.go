package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

// Defined literals
const (
	Null  = "null"
	True  = "true"
	False = "false"
)

func lexNull(l *Lexer) state {
	if !l.acceptWord(Null) {
		return unexpectedWord(Null)
	}
	l.emit(token.Null)
	return lexToken
}

func lexBoolean(l *Lexer) state {
	switch l.peek() {
	case 't':
		if !l.acceptWord(True) {
			return unexpectedWord(True)
		}
	case 'f':
		if !l.acceptWord(False) {
			return unexpectedWord(False)
		}
	}
	l.emit(token.Boolean)
	return lexToken
}
