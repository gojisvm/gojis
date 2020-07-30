//lint:file-ignore U1000 "work in progress"

package parser

import (
	"unicode"

	"github.com/gojisvm/gojis/internal/parser/matcher"
)

// Defined matchers
var (
	ampersand                         = matcher.String(`&`)
	assign                            = matcher.String(`=`)
	asterisk                          = matcher.String(`*`)
	backslash                         = matcher.String(`\`)
	backtick                          = matcher.String("`")
	binaryDigit                       = matcher.String(`01`)
	binaryIndicator                   = matcher.String(`bB`)
	braceClose                        = matcher.String(`}`)
	braceOpen                         = matcher.String(`{`)
	bracketClose                      = matcher.String(`]`)
	bracketOpen                       = matcher.String(`[`)
	caret                             = matcher.String(`^`)
	dash                              = matcher.String(`-`)
	decimalDigit                      = matcher.String(`0123456789`)
	dollar                            = matcher.String(`$`)
	dot                               = matcher.String(`.`)
	doubleQuote                       = matcher.String(`"`)
	exclamationMark                   = matcher.String(`!`)
	exponentIndicator                 = matcher.String(`eE`)
	greaterThan                       = matcher.String(`>`)
	hexDigit                          = matcher.String(`0123456789abcdefABCDEF`)
	hexIndicator                      = matcher.String(`xX`)
	lessThan                          = matcher.String(`<`)
	lowerU                            = matcher.String(`u`)
	lowerX                            = matcher.String(`x`)
	nonZeroDigit                      = matcher.String(`123456789`)
	octalDigit                        = matcher.String(`01234567`)
	octalIndicator                    = matcher.String(`oO`)
	percent                           = matcher.String(`%`)
	pipe                              = matcher.String(`|`)
	plus                              = matcher.String(`+`)
	punctuatorStart                   = matcher.String(`!%&()*+,-.:;<=>?[]^{|~`)
	regularExpressionCharPartial      = matcher.Diff(regularExpressionNonTerminator, matcher.String(`\/[`))
	regularExpressionClassCharPartial = matcher.Diff(regularExpressionNonTerminator, matcher.String(`]\`))
	regularExpressionFirstCharPartial = matcher.Diff(regularExpressionNonTerminator, matcher.String(`*\/[`))
	regularExpressionNonTerminator    = matcher.Diff(sourceCharacter, lineTerminator)
	sign                              = matcher.String(`+-`)
	singleQuote                       = matcher.String(`'`)
	slash                             = matcher.String(`/`)
	templateCharacterPartial          = matcher.Diff(sourceCharacter, matcher.Merge(backtick, backslash, dollar, lineTerminator))
	underscore                        = matcher.String(`_`)
	zero                              = matcher.String(`0`)

	sourceCharacter = matcher.New("SourceCharacter", codePoint) // 10.1

	whiteSpace            = matcher.Merge(formFeed, noBreakSpace, space, horizontalTab, unicodeSpace, verticalTab, zeroWidthJoiner, zeroWidthNoBreakSpace, zeroWidthNonJoiner) // 11.2
	formFeed              = matcher.RuneWithDesc("<FF>", '\u000C')                                                                                                             // 11.2
	noBreakSpace          = matcher.RuneWithDesc("<NBSP>", '\u00A0')                                                                                                           // 11.2
	space                 = matcher.RuneWithDesc("<SP>", '\u0020')                                                                                                             // 11.2
	horizontalTab         = matcher.RuneWithDesc("<TAB>", '\u0009')                                                                                                            // 11.2
	unicodeSpace          = matcher.New("<USP>", unicode.Z)                                                                                                                    // 11.2
	verticalTab           = matcher.RuneWithDesc("<VT>", '\u000B')                                                                                                             // 11.2
	zeroWidthJoiner       = matcher.RuneWithDesc("<ZWJ>", '\u200D')                                                                                                            // 11.1
	zeroWidthNoBreakSpace = matcher.RuneWithDesc("<ZWNBSP>", '\uFEFF')                                                                                                         // 11.1
	zeroWidthNonJoiner    = matcher.RuneWithDesc("<ZWNJ>", '\u200C')                                                                                                           // 11.1

	lineTerminator     = matcher.Merge(lineFeed, carriageReturn, lineSeparator, paragraphSeparator) // 11.3
	lineFeed           = matcher.RuneWithDesc("<LF>", '\u000A')                                     // 11.3
	carriageReturn     = matcher.RuneWithDesc("<CR>", '\u000D')                                     // 11.3
	lineSeparator      = matcher.RuneWithDesc("<LS>", '\u2028')                                     // 11.3
	paragraphSeparator = matcher.RuneWithDesc("<PS>", '\u2029')                                     // 11.3

	singleLineCommentChar = matcher.Diff(sourceCharacter, lineTerminator) // 11.4

	multiLineNotAsteriskChar               = matcher.Diff(sourceCharacter, asterisk)
	multiLineNotForwardSlashOrAsteriskChar = matcher.Diff(sourceCharacter, matcher.Merge(asterisk, slash))

	unicodeIDStart         = matcher.New("ID_Start", iDStart)
	unicodeIDContinue      = matcher.New("ID_Continue", iDContinue)
	identifierStartPartial = matcher.Merge(unicodeIDStart, dollar, underscore)
	identifierStart        = matcher.Merge(unicodeIDStart, dollar, underscore, backslash)
	identifierPartPartial  = matcher.Merge(unicodeIDContinue, dollar, zeroWidthNonJoiner, zeroWidthJoiner)
	identifierPart         = matcher.Merge(unicodeIDContinue, dollar, zeroWidthNonJoiner, zeroWidthJoiner, backslash)

	doubleStringCharacterPartial = matcher.Diff(sourceCharacter, matcher.Merge(doubleQuote, backslash, lineTerminator))
	singleStringCharacterPartial = matcher.Diff(sourceCharacter, matcher.Merge(singleQuote, backslash, lineTerminator))

	singleEscapeCharacter = matcher.String(`'"\bfnrtv`)
	escapeCharacter       = matcher.Merge(singleEscapeCharacter, decimalDigit, lowerX, lowerU)
	nonEscapeCharacter    = matcher.Diff(sourceCharacter, matcher.Merge(escapeCharacter, lineTerminator))
)
