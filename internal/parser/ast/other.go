package ast

type Initializer struct {
	AssignmentExpression *AssignmentExpression
}

type Elision struct {
	Commas int
}

type SuperProperty struct {
	Expression     *Expression
	IdentifierName string
}

type MetaProperty struct {
	NewTarget *NewTarget
}

type NewTarget struct {
	// "new.target"
}

type Arguments struct {
	ArgumentList *ArgumentList
	Comma        bool
}

type ArgumentList struct {
	AssignmentExpression *AssignmentExpression
	Ellipsis             bool
	ArgumentList         *ArgumentList
	Comma                bool
}

type LetOrConst struct {
	Let   bool
	Const bool
}
