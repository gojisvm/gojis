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
	for l.acceptMultiple(identifierPartPartial) > 0 {
		if l.accept(backslash) {
			// next sequence must be a unicode escape sequence
			errState := acceptUnicodeEscapeSequence(l)
			if errState != nil {
				return errState
			}
		}
	}

	// check if the identifier is a reserved word, and if so, emit it as such
	accepted := string(l.input[l.start:l.pos])
	if t, ok := allReservedWords[accepted]; ok {
		l.emit(t)
	} else {
		// identifier is not a reserved word, thus emit it as identifier name
		l.emit(token.IdentifierName)
	}
	return lexToken
}
