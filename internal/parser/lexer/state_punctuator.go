package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

func lexPunctuator(l *Lexer) state {
	switch r := l.peek(); r {
	case '>':
		return lexPunctuatorStartingWithGreaterThan
	}
	return tokenMismatch(PunctuatorStart)
}

func lexPunctuatorStartingWithGreaterThan(l *Lexer) state {
	// >>>=
	// >>>
	// >>=
	// >>
	// >=
	// >
	if l.accept(GreaterThan) {
		if l.accept(GreaterThan) {
			if l.accept(GreaterThan) {
				if l.accept(Equals) {
					// >>>=
					l.emit(token.UnsignedRightShiftAssign)
				} else {
					// >>>
					l.emit(token.UnsignedRightShift)
				}
			} else {
				if l.accept(Equals) {
					// >>=
					l.emit(token.RightShiftAssign)
				} else {
					// >>
					l.emit(token.RightShift)
				}
			}
		} else {
			if l.accept(Equals) {
				// >=
				l.emit(token.GreaterThanOrEqualTo)
			} else {
				// >
				l.emit(token.GreaterThan)
			}
		}
		return lexToken
	}
	return unexpectedToken
}
