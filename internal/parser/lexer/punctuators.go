package lexer

// Defined punctuators
var (
	// Punctuators as specified in 11.7
	Punctuators = []string{
		">>>=",
		">>>",
		">>=",
		">>",
		">=",
		">",
		"...",
		"<<=",
		"===",
		"!==",
		"**=",
		"<<",
		"<=",
		"==",
		"!=",
		"++",
		"--",
		"**",
		"&&",
		"||",
		"+=",
		"-=",
		"*=",
		"%=",
		"&=",
		"|=",
		"^=",
		"=>",
		".",
		";",
		",",
		"<",
		"{",
		"(",
		")",
		"[",
		"]",
		"!",
		"+",
		"-",
		"*",
		"%",
		"&",
		"|",
		"^",
		"~",
		"?",
		":",
		"=",
	}

	// DivPunctuators as specified in 11.7
	DivPunctuators = []string{
		"/=",
		"/",
	}

	RightBracePunctuator = BraceClose
)
