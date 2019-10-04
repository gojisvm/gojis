package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

func lexReservedWord(l *Lexer) state {
	for _, reservedWord := range AllReservedWords {
		if l.acceptWord(reservedWord) {
			// reserved word was matched
			l.emit(token.ReservedWord)
			return lexToken
		}
	}
	return errorf("Unexpected token '%s', expected reserved word", string(l.peek()))
}
