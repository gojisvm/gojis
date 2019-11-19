package lexer

// state is a definition of a state. The state returned will be the one that is
// executed next.
type state func(*Lexer) state

// lexScript is the default entry point for the lexer. This may be overridden by
// tests. It delegates directly to lexToken, which is responsible for all
// tokens.
func lexScript(l *Lexer) state {
	return lexToken
}

func lexToken(l *Lexer) state {
	if l.eof() {
		return nil
	}

	switch r := l.peek(); r {
	case '/':
		// lookahead one more rune to determine if this is a comment or a
		// regular expression literal
		next, ok := l.lookahead(1)
		if !ok {
			return unexpectedToken
		}
		switch next {
		case '/', '*':
			return lexComment
		default:
			return lexRegularExpressionLiteral
		}
	case '\\': // backslash or identifierStart (handled in default)
		return lexIdentifierName
	case '}':
		return lexPunctuator
	case '"', '\'', '`':
		return lexStringLiteral
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		next, ok := l.lookahead(1)
		if !ok {
			return unexpectedToken
		}

		if r == '.' && !decimalDigit.Matches(next) {
			return lexPunctuator
		}
		return lexNumericLiteral
	case '+', '-':
		next, ok := l.lookahead(1)
		if !ok {
			return unexpectedToken
		}

		if whiteSpace.Matches(next) {
			return lexPunctuator
		} else if decimalDigit.Matches(next) {
			return lexSignedInteger
		}
		return lexPunctuator
	default:
		// handle all cases that cannot be expressed in a switch block. for
		// optimization, matcher matches from here can be converted to cases if
		// necessary.

		if identifierStart.Matches(r) {
			return lexIdentifierName
		}
		if punctuatorStart.Matches(r) {
			return lexPunctuator
		}
		if whiteSpace.Matches(r) {
			return lexWhitespace
		}
		if lineTerminator.Matches(r) {
			return lexLineTerminator
		}
	}
	return unexpectedToken
}
