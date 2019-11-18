package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

func lexPunctuator(l *Lexer) state {
	switch r := l.peek(); r {
	case '>':
		return lexPunctuatorStartingWithGreaterThan
	case '<':
		return lexPunctuatorStartingWithLessThan
	case '.':
		return lexPunctuatorStartingWithDot
	case '=':
		return lexPunctuatorStartingWithAssign
	case '-':
		return lexPunctuatorStartingWithDash
	case '!':
		return lexPunctuatorStartingWithExclamationMark
	case '*':
		return lexPunctuatorStartingWithAsterisk
	case '&':
		return lexPunctuatorStartingWithAmpersand
	case '%':
		return lexPunctuatorStartingWithPercent
	case '^':
		return lexPunctuatorStartingWithCaret
	case '+':
		return lexPunctuatorStartingWithPlus
	case '|':
		return lexPunctuatorStartingWithPipe
	case '/':
		return lexPunctuatorStartingWithSlash
	case ',':
		_ = l.next() // consume the rune
		l.emit(token.Comma)
		return lexToken
	case ';':
		_ = l.next() // consume the rune
		l.emit(token.SemiColon)
		return lexToken
	case ':':
		_ = l.next() // consume the rune
		l.emit(token.Colon)
		return lexToken
	case '?':
		_ = l.next() // consume the rune
		l.emit(token.QuestionMark)
		return lexToken
	case '(':
		_ = l.next() // consume the rune
		l.emit(token.ParOpen)
		return lexToken
	case ')':
		_ = l.next() // consume the rune
		l.emit(token.ParClose)
		return lexToken
	case '[':
		_ = l.next() // consume the rune
		l.emit(token.BracketOpen)
		return lexToken
	case ']':
		_ = l.next() // consume the rune
		l.emit(token.BracketClose)
		return lexToken
	case '{':
		_ = l.next() // consume the rune
		l.emit(token.BraceOpen)
		return lexToken
	case '}':
		_ = l.next() // consume the rune
		l.emit(token.BraceClose)
		return lexToken
	case '~':
		_ = l.next() // consume the rune
		l.emit(token.Tilde)
		return lexToken
	}
	return tokenMismatch(punctuatorStart)
}

func lexPunctuatorStartingWithSlash(l *Lexer) state {
	if !l.accept(slash) {
		return tokenMismatch(slash)
	}

	if l.accept(assign) {
		l.emit(token.DivAssign)
	} else {
		l.emit(token.Slash)
	}
	return lexToken
}

func lexPunctuatorStartingWithPipe(l *Lexer) state {
	if !l.accept(pipe) {
		return tokenMismatch(pipe)
	}

	if l.accept(pipe) {
		l.emit(token.LogicalOr)
	} else if l.accept(assign) {
		l.emit(token.OrAssign)
	} else {
		l.emit(token.BitwiseOr)
	}

	return lexToken
}

func lexPunctuatorStartingWithPlus(l *Lexer) state {
	if !l.accept(plus) {
		return tokenMismatch(plus)
	}

	if l.accept(plus) {
		l.emit(token.UpdatePlus)
	} else if l.accept(assign) {
		l.emit(token.PlusAssign)
	} else {
		l.emit(token.Plus)
	}

	return lexToken
}

func lexPunctuatorStartingWithCaret(l *Lexer) state {
	if !l.accept(caret) {
		return tokenMismatch(caret)
	}

	if l.accept(assign) {
		l.emit(token.XorAssign)
	} else {
		l.emit(token.BitwiseXor)
	}

	return lexToken
}

func lexPunctuatorStartingWithPercent(l *Lexer) state {
	if !l.accept(percent) {
		return tokenMismatch(percent)
	}

	if l.accept(assign) {
		l.emit(token.ModuloAssign)
	} else {
		l.emit(token.Modulo)
	}

	return lexToken
}

func lexPunctuatorStartingWithAmpersand(l *Lexer) state {
	if !l.accept(ampersand) {
		return tokenMismatch(ampersand)
	}

	if l.accept(ampersand) {
		l.emit(token.LogicalAnd)
	} else if l.accept(assign) {
		l.emit(token.AndAssign)
	} else {
		l.emit(token.BitwiseAnd)
	}

	return lexToken
}

func lexPunctuatorStartingWithAsterisk(l *Lexer) state {
	if !l.accept(asterisk) {
		return tokenMismatch(asterisk)
	}

	if l.accept(asterisk) {
		if l.accept(assign) {
			l.emit(token.PowerAssign)
		} else {
			l.emit(token.Power)
		}
	} else {
		if l.accept(assign) {
			l.emit(token.MultiplyAssign)
		} else {
			l.emit(token.Asterisk)
		}
	}
	return lexToken
}

func lexPunctuatorStartingWithExclamationMark(l *Lexer) state {
	if !l.accept(exclamationMark) {
		return tokenMismatch(exclamationMark)
	}

	if l.accept(assign) {
		if l.accept(assign) {
			l.emit(token.StrictNotEquals)
		} else {
			l.emit(token.NotEquals)
		}
	} else {
		l.emit(token.ExclamationMark)
	}
	return lexToken
}

func lexPunctuatorStartingWithDash(l *Lexer) state {
	if !l.accept(dash) {
		return tokenMismatch(dash)
	}
	if l.accept(dash) {
		l.emit(token.UpdateMinus)
	} else if l.accept(assign) {
		l.emit(token.MinusAssign)
	} else {
		l.emit(token.Minus)
	}
	return lexToken
}

func lexPunctuatorStartingWithAssign(l *Lexer) state {
	// ===
	// ==
	// =>
	// =
	if !l.accept(assign) {
		return unexpectedToken
	}
	if l.accept(assign) {
		if l.accept(assign) {
			l.emit(token.StrictEquals)
		} else {
			l.emit(token.Equals)
		}
	} else if l.accept(greaterThan) {
		l.emit(token.Arrow)
	} else {
		l.emit(token.Assign)
	}
	return lexToken
}

func lexPunctuatorStartingWithDot(l *Lexer) state {
	// ...
	// .
	if l.acceptWord("...") {
		l.emit(token.Ellipsis)
	}
	if !l.accept(dot) {
		return unexpectedToken
	}
	l.emit(token.Dot)
	return lexToken
}

func lexPunctuatorStartingWithLessThan(l *Lexer) state {
	// <<=
	// <<
	// <=
	// <
	if !l.accept(lessThan) {
		return unexpectedToken
	}
	if l.accept(lessThan) {
		if l.accept(assign) {
			l.emit(token.LeftShiftAssign)
		} else {
			l.emit(token.LeftShift)
		}
	} else {
		if l.accept(assign) {
			l.emit(token.LessThanOrEqualTo)
		} else {
			l.emit(token.LessThan)
		}
	}
	return lexToken
}

func lexPunctuatorStartingWithGreaterThan(l *Lexer) state {
	// >>>=
	// >>>
	// >>=
	// >>
	// >=
	// >
	if !l.accept(greaterThan) {
		return unexpectedToken
	}
	if l.accept(greaterThan) {
		if l.accept(greaterThan) {
			if l.accept(assign) {
				// >>>=
				l.emit(token.UnsignedRightShiftAssign)
			} else {
				// >>>
				l.emit(token.UnsignedRightShift)
			}
		} else {
			if l.accept(assign) {
				// >>=
				l.emit(token.RightShiftAssign)
			} else {
				// >>
				l.emit(token.RightShift)
			}
		}
	} else {
		if l.accept(assign) {
			// >=
			l.emit(token.GreaterThanOrEqualTo)
		} else {
			// >
			l.emit(token.GreaterThan)
		}
	}
	return lexToken
}
