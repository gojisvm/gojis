package ast

// StatementList represents a StatementList node.
type StatementList struct {
	StatementListItems []*StatementListItem
}

// StatementListItem represents a StatementListItem node.
type StatementListItem struct {
	Statement   *Statement
	Declaration *Declaration
}

// Statement represents a Statement node.
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

// BlockStatement represents a BlockStatement node.
type BlockStatement struct {
	Block *Block
}

// Block represents a Block node.
type Block struct {
	StatementList *StatementList
}

// VariableStatement represents a VariableStatement node.
type VariableStatement struct {
	VariableDeclarationList *VariableDeclarationList
}

// EmptyStatement represents a EmptyStatement node.
type EmptyStatement struct {
	// ";"
}

// ExpressionStatement represents a ExpressionStatement node.
type ExpressionStatement struct {
	Expression *Expression
}

// IfStatement represents a IfStatement node.
type IfStatement struct {
	Expression    *Expression
	Statement     *Statement
	ElseStatement *Statement
}

// BreakableStatement represents a BreakableStatement node.
type BreakableStatement struct {
	IterationStatement *IterationStatement
	SwitchStatement    *SwitchStatement
}

// IterationStatement represents a IterationStatement node.
type IterationStatement struct {
	Statement               *Statement
	Expression1             *Expression
	Expression2             *Expression
	Expression3             *Expression
	VariableDeclarationList *VariableDeclarationList
	LexicalDeclaration      *LexicalDeclaration
	LeftHandSideExpression  *LeftHandSideExpression
	ForBinding              *ForBinding
	ForDeclaration          *ForDeclaration
	AssignmentExpression    *AssignmentExpression
	Do                      bool
	While                   bool
	For                     bool
	In                      bool
	Var                     bool
	Of                      bool
	Await                   bool
}

// SwitchStatement represents a SwitchStatement node.
type SwitchStatement struct {
	Expression *Expression
	CaseBlock  *CaseBlock
}

// CaseBlock represents a CaseBlock node.
type CaseBlock struct {
	CaseClauses   []*CaseClause
	DefaultClause *DefaultClause
}

// CaseClause represents a CaseClause node.
type CaseClause struct {
	Expression    *Expression
	StatementList *StatementList
}

// DefaultClause represents a DefaultClause node.
type DefaultClause struct {
	StatementList *StatementList
}

// ContinueStatement represents a ContinueStatement node.
type ContinueStatement struct {
	LabelIdentifier *LabelIdentifier
}

// BreakStatement represents a BreakStatement node.
type BreakStatement struct {
	LabelIdentifier *LabelIdentifier
}

// ReturnStatement represents a ReturnStatement node.
type ReturnStatement struct {
	Expression *Expression
}

// WithStatement represents a WithStatement node.
type WithStatement struct {
	Expression *Expression
	Statement  *Statement
}

// LabelledStatement represents a LabelledStatement node.
type LabelledStatement struct {
	LabelIdentifier *LabelIdentifier
	LabelledItem    *LabelledItem
}

// LabelledItem represents a LabelledItem node.
type LabelledItem struct {
	Statement           *Statement
	FunctionDeclaration *FunctionDeclaration
}

// ThrowStatement represents a ThrowStatement node.
type ThrowStatement struct {
	Expression *Expression
}

// TryStatement represents a TryStatement node.
type TryStatement struct {
	Block   *Block
	Catch   *Catch
	Finally *Finally
}

// Catch represents a Catch node.
type Catch struct {
	CatchParameter *CatchParameter
	Block          *Block
}

// CatchParameter represents a CatchParameter node.
type CatchParameter struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
}

// Finally represents a Finally node.
type Finally struct {
	Block *Block
}

// DebuggerStatement represents a DebuggerStatement node.
type DebuggerStatement struct {
	// "debugger ;"
}
