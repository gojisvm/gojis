package lexer

func lexReservedWord(l *Lexer) state {
	for reservedWord, tokenType := range allReservedWords {
		if l.acceptWord(reservedWord) {
			// reserved word was matched
			l.emit(tokenType)
			return lexToken
		}
	}
	return errorf("Unexpected token '%s', expected reserved word", string(l.peek()))
}
