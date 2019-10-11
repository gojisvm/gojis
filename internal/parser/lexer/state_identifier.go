package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

func lexIdentifierName(l *Lexer) state {
	return lexIdentifierStart
}

func lexIdentifierStart(l *Lexer) state {
	if !l.accept(identifierStartPartial) {
		if l.accept(backslash) {
			// next sequence must be a unicode escape sequence
			errState := acceptUnicodeEscapeSequence(l)
			if errState != nil {
				return errState
			}
		}
	}
	return lexIdentifierPart
}

func lexIdentifierPart(l *Lexer) state {
	for l.acceptMultiple(identifierPartPartial) >= 0 {
		if l.accept(backslash) {
			errState := acceptUnicodeEscapeSequence(l)
			if errState != nil {
				return errState
			}
		} else {
			break
		}
	}

	l.emit(token.IdentifierName)
	return lexToken
}
