package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

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
	KeywordStatic     = "static"
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
	allReservedWords = map[string]token.Type{
		KeywordAwait:           token.Await,
		KeywordBreak:           token.Break,
		KeywordCase:            token.Case,
		KeywordCatch:           token.Catch,
		KeywordClass:           token.Class,
		KeywordConst:           token.Const,
		KeywordContinue:        token.Continue,
		KeywordDebugger:        token.Debugger,
		KeywordDefault:         token.Default,
		KeywordDelete:          token.Delete,
		KeywordDo:              token.Do,
		KeywordElse:            token.Else,
		KeywordExport:          token.Export,
		KeywordExtends:         token.Extends,
		KeywordFinally:         token.Finally,
		KeywordFor:             token.For,
		KeywordFunction:        token.Function,
		KeywordIf:              token.If,
		KeywordImport:          token.Import,
		KeywordIn:              token.In,
		KeywordInstanceof:      token.Instanceof,
		KeywordNew:             token.NewT,
		KeywordReturn:          token.Return,
		KeywordStatic:          token.Static,
		KeywordSuper:           token.Super,
		KeywordSwitch:          token.Switch,
		KeywordThis:            token.This,
		KeywordThrow:           token.Throw,
		KeywordTry:             token.Try,
		KeywordTypeof:          token.Typeof,
		KeywordVar:             token.Var,
		KeywordVoid:            token.Void,
		KeywordWhile:           token.While,
		KeywordWith:            token.With,
		KeywordYield:           token.Yield,
		ReservedWordEnum:       token.Enum,
		ReservedWordImplements: token.Implements,
		ReservedWordPackage:    token.Package,
		ReservedWordProtected:  token.Protected,
		ReservedWordInterface:  token.Interface,
		ReservedWordPrivate:    token.Private,
		ReservedWordPublic:     token.Public,
	}
)
