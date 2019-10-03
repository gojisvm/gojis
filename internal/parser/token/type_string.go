// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Error-1]
	_ = x[LineTerminator-2]
	_ = x[Whitespace-3]
	_ = x[MultiLineComment-4]
	_ = x[SingleLineComment-5]
	_ = x[CommonToken-6]
	_ = x[IdentifierName-7]
	_ = x[Punctuator-8]
	_ = x[NumericLiteral-9]
	_ = x[StringLiteral-10]
	_ = x[Template-11]
	_ = x[ReservedWord-12]
}

const _Type_name = "UnknownErrorLineTerminatorWhitespaceMultiLineCommentSingleLineCommentCommonTokenIdentifierNamePunctuatorNumericLiteralStringLiteralTemplateReservedWord"

var _Type_index = [...]uint8{0, 7, 12, 26, 36, 52, 69, 80, 94, 104, 118, 131, 139, 151}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}