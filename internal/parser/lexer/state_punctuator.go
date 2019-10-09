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
	case ';':
		_ = l.next() // consume the rune
		l.emit(token.SemiColon)
	case ':':
		_ = l.next() // consume the rune
		l.emit(token.Colon)
	case '?':
		_ = l.next() // consume the rune
		l.emit(token.QuestionMark)
	case '(':
		_ = l.next() // consume the rune
		l.emit(token.ParOpen)
	case ')':
		_ = l.next() // consume the rune
		l.emit(token.ParClose)
	case '[':
		_ = l.next() // consume the rune
		l.emit(token.BracketOpen)
	case ']':
		_ = l.next() // consume the rune
		l.emit(token.BracketClose)
	case '{':
		_ = l.next() // consume the rune
		l.emit(token.BraceOpen)
	case '}':
		_ = l.next() // consume the rune
		l.emit(token.BraceClose)
	case '~':
		_ = l.next() // consume the rune
		l.emit(token.Tilde)
	}
	return tokenMismatch(PunctuatorStart)
}

func lexPunctuatorStartingWithSlash(l *Lexer) state {
	if !l.accept(Slash) {
		return tokenMismatch(Slash)
	}

	if l.accept(Assign) {
		l.emit(token.DivAssign)
	} else {
		l.emit(token.Div)
	}
	return lexToken
}

func lexPunctuatorStartingWithPipe(l *Lexer) state {
	if !l.accept(Pipe) {
		return tokenMismatch(Pipe)
	}

	if l.accept(Pipe) {
		l.emit(token.LogicalOr)
	} else if l.accept(Assign) {
		l.emit(token.OrAssign)
	} else {
		l.emit(token.BitwiseOr)
	}

	return lexToken
}

func lexPunctuatorStartingWithPlus(l *Lexer) state {
	if !l.accept(Plus) {
		return tokenMismatch(Plus)
	}

	if l.accept(Plus) {
		l.emit(token.UpdatePlus)
	} else if l.accept(Assign) {
		l.emit(token.PlusAssign)
	} else {
		l.emit(token.Plus)
	}

	return lexToken
}

func lexPunctuatorStartingWithCaret(l *Lexer) state {
	if !l.accept(Caret) {
		return tokenMismatch(Caret)
	}

	if l.accept(Assign) {
		l.emit(token.XorAssign)
	} else {
		l.emit(token.BitwiseXor)
	}

	return lexToken
}

func lexPunctuatorStartingWithPercent(l *Lexer) state {
	if !l.accept(Percent) {
		return tokenMismatch(Percent)
	}

	if l.accept(Assign) {
		l.emit(token.ModuloAssign)
	} else {
		l.emit(token.Modulo)
	}

	return lexToken
}

func lexPunctuatorStartingWithAmpersand(l *Lexer) state {
	if !l.accept(Ampersand) {
		return tokenMismatch(Ampersand)
	}

	if l.accept(Ampersand) {
		l.emit(token.LogicalAnd)
	} else if l.accept(Assign) {
		l.emit(token.AndAssign)
	} else {
		l.emit(token.BitwiseAnd)
	}

	return lexToken
}

func lexPunctuatorStartingWithAsterisk(l *Lexer) state {
	if !l.accept(Asterisk) {
		return tokenMismatch(Asterisk)
	}

	if l.accept(Asterisk) {
		if l.accept(Assign) {
			l.emit(token.PowerAssign)
		} else {
			l.emit(token.Power)
		}
	} else {
		if l.accept(Assign) {
			l.emit(token.MultiplyAssign)
		} else {
			l.emit(token.Multiply)
		}
	}
	return lexToken
}

func lexPunctuatorStartingWithExclamationMark(l *Lexer) state {
	if !l.accept(ExclamationMark) {
		return tokenMismatch(ExclamationMark)
	}

	if l.accept(Assign) {
		if l.accept(Assign) {
			l.emit(token.StrictNotEquals)
		} else {
			l.emit(token.NotEquals)
		}
	} else {
		l.emit(token.LogicalNot)
	}
	return lexToken
}

func lexPunctuatorStartingWithDash(l *Lexer) state {
	if !l.accept(Dash) {
		return tokenMismatch(Dash)
	}
	if l.accept(Dash) {
		l.emit(token.UpdateMinus)
	} else if l.accept(Assign) {
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
	if !l.accept(Assign) {
		return unexpectedToken
	}
	if l.accept(Assign) {
		if l.accept(Assign) {
			l.emit(token.StrictEquals)
		} else {
			l.emit(token.Equals)
		}
	} else if l.accept(GreaterThan) {
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
	if !l.accept(Dot) {
		return unexpectedToken
	} else {
		l.emit(token.Dot)
	}
	return lexToken
}

func lexPunctuatorStartingWithLessThan(l *Lexer) state {
	// <<=
	// <<
	// <=
	// <
	if !l.accept(LessThan) {
		return unexpectedToken
	}
	if l.accept(LessThan) {
		if l.accept(Assign) {
			l.emit(token.LeftShiftAssign)
		} else {
			l.emit(token.LeftShift)
		}
	} else {
		if l.accept(Assign) {
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
	if !l.accept(GreaterThan) {
		return unexpectedToken
	}
	if l.accept(GreaterThan) {
		if l.accept(GreaterThan) {
			if l.accept(Assign) {
				// >>>=
				l.emit(token.UnsignedRightShiftAssign)
			} else {
				// >>>
				l.emit(token.UnsignedRightShift)
			}
		} else {
			if l.accept(Assign) {
				// >>=
				l.emit(token.RightShiftAssign)
			} else {
				// >>
				l.emit(token.RightShift)
			}
		}
	} else {
		if l.accept(Assign) {
			// >=
			l.emit(token.GreaterThanOrEqualTo)
		} else {
			// >
			l.emit(token.GreaterThan)
		}
	}
	return lexToken
}
