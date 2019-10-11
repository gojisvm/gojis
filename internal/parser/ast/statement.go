package ast

type StatementList struct {
	StatementListItems []*StatementListItem
}

type StatementListItem struct {
	Statement *Statement
	// Declaration *Declaration
}

type Statement struct {
	BlockStatement      *BlockStatement
	VariableStatement   *VariableStatement
	EmptyStatement      *EmptyStatement
	ExpressionStatement *ExpressionStatement
	IfStatement         *IfStatement
	BreakableStatement  *BreakableStatement
	ContinueStatement   *ContinueStatement
	BreakStatement      *BreakStatement
	ReturnStatement     *ReturnStatement
	WithStatement       *WithStatement
	LabelledStatement   *LabelledStatement
	ThrowStatement      *ThrowStatement
	TryStatement        *TryStatement
	DebuggerStatement   *DebuggerStatement
}

type BlockStatement struct {
	Block *Block
}

type Block struct {
	StatementList *StatementList
}

type VariableStatement struct {
}

type EmptyStatement struct {
}

type ExpressionStatement struct {
}

type IfStatement struct {
}

type BreakableStatement struct {
}

type ContinueStatement struct {
}

type BreakStatement struct {
}

type ReturnStatement struct {
}

type WithStatement struct {
}

type LabelledStatement struct {
}

type ThrowStatement struct {
}

type TryStatement struct {
}

type DebuggerStatement struct {
}
