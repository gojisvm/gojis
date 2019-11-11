// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[LineTerminator-1]
	_ = x[Whitespace-2]
	_ = x[MultiLineComment-3]
	_ = x[SingleLineComment-4]
	_ = x[CommonToken-5]
	_ = x[IdentifierName-6]
	_ = x[Punctuator-7]
	_ = x[StringLiteral-8]
	_ = x[TemplateHead-9]
	_ = x[TemplateMiddle-10]
	_ = x[TemplateTail-11]
	_ = x[NoSubstitutionTemplate-12]
	_ = x[Async-13]
	_ = x[Let-14]
	_ = x[Asterisk-15]
	_ = x[Await-16]
	_ = x[Break-17]
	_ = x[Case-18]
	_ = x[Catch-19]
	_ = x[Class-20]
	_ = x[Const-21]
	_ = x[Continue-22]
	_ = x[Debugger-23]
	_ = x[Default-24]
	_ = x[Delete-25]
	_ = x[Do-26]
	_ = x[Else-27]
	_ = x[Export-28]
	_ = x[Extends-29]
	_ = x[Finally-30]
	_ = x[For-31]
	_ = x[Function-32]
	_ = x[Get-33]
	_ = x[If-34]
	_ = x[Import-35]
	_ = x[In-36]
	_ = x[Instanceof-37]
	_ = x[NewT-38]
	_ = x[Return-39]
	_ = x[Set-40]
	_ = x[Static-41]
	_ = x[Super-42]
	_ = x[Switch-43]
	_ = x[This-44]
	_ = x[Throw-45]
	_ = x[Try-46]
	_ = x[Typeof-47]
	_ = x[Var-48]
	_ = x[Void-49]
	_ = x[While-50]
	_ = x[With-51]
	_ = x[Yield-52]
	_ = x[Enum-53]
	_ = x[Implements-54]
	_ = x[Package-55]
	_ = x[Protected-56]
	_ = x[Interface-57]
	_ = x[Private-58]
	_ = x[Public-59]
	_ = x[Null-60]
	_ = x[Boolean-61]
	_ = x[DecimalLiteral-62]
	_ = x[BinaryIntegerLiteral-63]
	_ = x[OctalIntegerLiteral-64]
	_ = x[HexIntegerLiteral-65]
	_ = x[RegularExpressionBody-66]
	_ = x[RegularExpressionFlags-67]
	_ = x[Target-68]
	_ = x[AndAssign-69]
	_ = x[Arrow-70]
	_ = x[Assign-71]
	_ = x[BitwiseAnd-72]
	_ = x[BitwiseNot-73]
	_ = x[BitwiseOr-74]
	_ = x[BitwiseXor-75]
	_ = x[BraceClose-76]
	_ = x[BraceOpen-77]
	_ = x[BracketClose-78]
	_ = x[BracketOpen-79]
	_ = x[Colon-80]
	_ = x[Comma-81]
	_ = x[DivAssign-82]
	_ = x[Dot-83]
	_ = x[Ellipsis-84]
	_ = x[Equals-85]
	_ = x[ExclamationMark-86]
	_ = x[GreaterThan-87]
	_ = x[GreaterThanOrEqualTo-88]
	_ = x[LeftShift-89]
	_ = x[LeftShiftAssign-90]
	_ = x[LessThan-91]
	_ = x[LessThanOrEqualTo-92]
	_ = x[LogicalAnd-93]
	_ = x[LogicalOr-94]
	_ = x[Minus-95]
	_ = x[MinusAssign-96]
	_ = x[Modulo-97]
	_ = x[ModuloAssign-98]
	_ = x[MultiplyAssign-99]
	_ = x[NotEquals-100]
	_ = x[OrAssign-101]
	_ = x[ParClose-102]
	_ = x[ParOpen-103]
	_ = x[Plus-104]
	_ = x[PlusAssign-105]
	_ = x[Power-106]
	_ = x[PowerAssign-107]
	_ = x[QuestionMark-108]
	_ = x[RightShift-109]
	_ = x[RightShiftAssign-110]
	_ = x[SemiColon-111]
	_ = x[Slash-112]
	_ = x[StrictEquals-113]
	_ = x[StrictNotEquals-114]
	_ = x[Tilde-115]
	_ = x[UnsignedRightShift-116]
	_ = x[UnsignedRightShiftAssign-117]
	_ = x[UpdateMinus-118]
	_ = x[UpdatePlus-119]
	_ = x[XorAssign-120]
}

const _Type_name = "UnknownLineTerminatorWhitespaceMultiLineCommentSingleLineCommentCommonTokenIdentifierNamePunctuatorStringLiteralTemplateHeadTemplateMiddleTemplateTailNoSubstitutionTemplateAsyncLetAsteriskAwaitBreakCaseCatchClassConstContinueDebuggerDefaultDeleteDoElseExportExtendsFinallyForFunctionGetIfImportInInstanceofNewTReturnSetStaticSuperSwitchThisThrowTryTypeofVarVoidWhileWithYieldEnumImplementsPackageProtectedInterfacePrivatePublicNullBooleanDecimalLiteralBinaryIntegerLiteralOctalIntegerLiteralHexIntegerLiteralRegularExpressionBodyRegularExpressionFlagsTargetAndAssignArrowAssignBitwiseAndBitwiseNotBitwiseOrBitwiseXorBraceCloseBraceOpenBracketCloseBracketOpenColonCommaDivAssignDotEllipsisEqualsExclamationMarkGreaterThanGreaterThanOrEqualToLeftShiftLeftShiftAssignLessThanLessThanOrEqualToLogicalAndLogicalOrMinusMinusAssignModuloModuloAssignMultiplyAssignNotEqualsOrAssignParCloseParOpenPlusPlusAssignPowerPowerAssignQuestionMarkRightShiftRightShiftAssignSemiColonSlashStrictEqualsStrictNotEqualsTildeUnsignedRightShiftUnsignedRightShiftAssignUpdateMinusUpdatePlusXorAssign"

var _Type_index = [...]uint16{0, 7, 21, 31, 47, 64, 75, 89, 99, 112, 124, 138, 150, 172, 177, 180, 188, 193, 198, 202, 207, 212, 217, 225, 233, 240, 246, 248, 252, 258, 265, 272, 275, 283, 286, 288, 294, 296, 306, 310, 316, 319, 325, 330, 336, 340, 345, 348, 354, 357, 361, 366, 370, 375, 379, 389, 396, 405, 414, 421, 427, 431, 438, 452, 472, 491, 508, 529, 551, 557, 566, 571, 577, 587, 597, 606, 616, 626, 635, 647, 658, 663, 668, 677, 680, 688, 694, 709, 720, 740, 749, 764, 772, 789, 799, 808, 813, 824, 830, 842, 856, 865, 873, 881, 888, 892, 902, 907, 918, 930, 940, 956, 965, 970, 982, 997, 1002, 1020, 1044, 1055, 1065, 1074}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
