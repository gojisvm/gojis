package ast

type StatementList struct {
	StatementListItems []*StatementListItem
}

type StatementListItem struct {
	Statement   *Statement
	Declaration *Declaration
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
	VariableDeclarationList *VariableDeclarationList
}

type EmptyStatement struct {
	// ";"
}

type ExpressionStatement struct {
	Expression *Expression
}

type IfStatement struct {
	Expression    *Expression
	Statement     *Statement
	ElseStatement *Statement
}

type BreakableStatement struct {
	IterationStatement *IterationStatement
	SwitchStatement    *SwitchStatement
}

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

type SwitchStatement struct {
	Expression *Expression
	CaseBlock  *CaseBlock
}

type CaseBlock struct {
	CaseClauses   []*CaseClause
	DefaultClause *DefaultClause
}

type CaseClause struct {
	Expression    *Expression
	StatementList *StatementList
}

type DefaultClause struct {
	StatementList *StatementList
}

type ContinueStatement struct {
	LabelIdentifier *LabelIdentifier
}

type BreakStatement struct {
	LabelIdentifier *LabelIdentifier
}

type ReturnStatement struct {
	Expression *Expression
}

type WithStatement struct {
	Expression *Expression
	Statement  *Statement
}

type LabelledStatement struct {
	LabelIdentifier *LabelIdentifier
	LabelledItem    *LabelledItem
}

type LabelledItem struct {
	Statement           *Statement
	FunctionDeclaration *FunctionDeclaration
}

type ThrowStatement struct {
	Expression *Expression
}

type TryStatement struct {
	Block   *Block
	Catch   *Catch
	Finally *Finally
}

type Catch struct {
	CatchParameter *CatchParameter
	Block          *Block
}

type CatchParameter struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
}

type Finally struct {
	Block *Block
}

type DebuggerStatement struct {
	// "debugger ;"
}
