package lexer

// Defined keywords
const (
	// Keyword

	KeywordAwait      = "await"
	KeywordBreak      = "break"
	KeywordCase       = "case"
	KeywordCatch      = "catch"
	KeywordClass      = "class"
	KeywordConst      = "const"
	KeywordContinue   = "continue"
	KeywordDebugger   = "debugger"
	KeywordDefault    = "default"
	KeywordDelete     = "delete"
	KeywordDo         = "do"
	KeywordElse       = "else"
	KeywordExport     = "export"
	KeywordExtends    = "extends"
	KeywordFinally    = "finally"
	KeywordFor        = "for"
	KeywordFunction   = "function"
	KeywordIf         = "if"
	KeywordImport     = "import"
	KeywordIn         = "in"
	KeywordInstanceof = "instanceof"
	KeywordNew        = "new"
	KeywordReturn     = "return"
	KeywordSuper      = "super"
	KeywordSwitch     = "switch"
	KeywordThis       = "this"
	KeywordThrow      = "throw"
	KeywordTry        = "try"
	KeywordTypeof     = "typeof"
	KeywordVar        = "var"
	KeywordVoid       = "void"
	KeywordWhile      = "while"
	KeywordWith       = "with"
	KeywordYield      = "yield"

	// Future reserved word

	ReservedWordEnum       = "enum"
	ReservedWordImplements = "implements" // only reserved in strict mode
	ReservedWordPackage    = "package"    // only reserved in strict mode
	ReservedWordProtected  = "protected"  // only reserved in strict mode
	ReservedWordInterface  = "interface"  // only reserved in strict mode
	ReservedWordPrivate    = "private"    // only reserved in strict mode
	ReservedWordPublic     = "public"     // only reserved in strict mode
)

var (
	// allReservedWords is a slice containing all keywords and future reserved
	// words.
	allReservedWords = []string{
		KeywordAwait,
		KeywordBreak,
		KeywordCase,
		KeywordCatch,
		KeywordClass,
		KeywordConst,
		KeywordContinue,
		KeywordDebugger,
		KeywordDefault,
		KeywordDelete,
		KeywordDo,
		KeywordElse,
		KeywordExport,
		KeywordExtends,
		KeywordFinally,
		KeywordFor,
		KeywordFunction,
		KeywordIf,
		KeywordImport,
		KeywordIn,
		KeywordInstanceof,
		KeywordNew,
		KeywordReturn,
		KeywordSuper,
		KeywordSwitch,
		KeywordThis,
		KeywordThrow,
		KeywordTry,
		KeywordTypeof,
		KeywordVar,
		KeywordVoid,
		KeywordWhile,
		KeywordWith,
		KeywordYield,
		ReservedWordEnum,
		ReservedWordImplements,
		ReservedWordPackage,
		ReservedWordProtected,
		ReservedWordInterface,
		ReservedWordPrivate,
		ReservedWordPublic,
	}
)
