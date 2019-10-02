package lexer

import (
	"unicode"

	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
)

// Defined matchers
var (
	Asterisk     = matcher.Rune('*')
	BraceClose   = matcher.Rune('}')
	BraceOpen    = matcher.Rune('{')
	BracketClose = matcher.Rune(']')
	BracketOpen  = matcher.Rune('[')
	Slash        = matcher.Rune('/')

	SourceCharacter = matcher.New("SourceCharacter", isUnicodeCodePoint) // 10.1

	WhiteSpace = matcher.Merge(_FF, _NBSP, _SP, _TAB, _USP, _VT, _ZWJ, _ZWNBSP, _ZWNJ) // 11.2
	_FF        = matcher.Rune('\u000C')                                                // 11.2
	_NBSP      = matcher.Rune('\u00A0')                                                // 11.2
	_SP        = matcher.Rune('\u0020')                                                // 11.2
	_TAB       = matcher.Rune('\u0009')                                                // 11.2
	_USP       = matcher.New("<USP>", isUnicodeCategoryZ)                              // 11.2
	_VT        = matcher.Rune('\u000B')                                                // 11.2
	_ZWJ       = matcher.Rune('\u200D')                                                // 11.1
	_ZWNBSP    = matcher.Rune('\uFEFF')                                                // 11.1
	_ZWNJ      = matcher.Rune('\u200C')                                                // 11.1

	LineTerminator = matcher.Merge(_LF, _CR, _LS, _PS) // 11.3
	_LF            = matcher.Rune('\u000A')            // 11.3
	_CR            = matcher.Rune('\u000D')            // 11.3
	_LS            = matcher.Rune('\u2028')            // 11.3
	_PS            = matcher.Rune('\u2029')            // 11.3

	SingleLineCommentChar = matcher.Diff(SourceCharacter, LineTerminator) // 11.4

	MultiLineNotAsteriskChar               = matcher.Diff(SourceCharacter, Asterisk)
	MultiLineNotForwardSlashOrAsteriskChar = matcher.Diff(SourceCharacter, matcher.Merge(Asterisk, Slash))
)

func isUnicodeCategoryZ(r rune) bool {
	return unicode.Is(unicode.Z, r)
}

func isUnicodeCodePoint(r rune) bool {
	return r <= unicode.MaxRune
}
