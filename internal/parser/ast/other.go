package ast

// Initializer represents a Initializer node.
type Initializer struct {
	AssignmentExpression *AssignmentExpression
}

// Elision represents a Elision node.
type Elision struct {
	Commas int
}

// SuperProperty represents a SuperProperty node.
type SuperProperty struct {
	Expression     *Expression
	IdentifierName string
}

// MetaProperty represents a MetaProperty node.
type MetaProperty struct {
	NewTarget *NewTarget
}

// NewTarget represents a NewTarget node.
type NewTarget struct {
	// "new.target"
}

// Arguments represents a Arguments node.
type Arguments struct {
	ArgumentList *ArgumentList
	Comma        bool
}

// ArgumentList represents a ArgumentList node.
type ArgumentList struct {
	AssignmentExpression *AssignmentExpression
	Ellipsis             bool
	ArgumentList         *ArgumentList
	Comma                bool
}

// LetOrConst represents a LetOrConst node.
type LetOrConst struct {
	Let   bool
	Const bool
}
