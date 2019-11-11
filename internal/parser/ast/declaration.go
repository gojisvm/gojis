package ast

// Declaration represents a Declaration node.
type Declaration struct {
	HoistableDeclaration *HoistableDeclaration
	ClassDeclaration     *ClassDeclaration
	LexicalDeclaration   *LexicalDeclaration
}

// HoistableDeclaration represents a HoistableDeclaration node.
type HoistableDeclaration struct {
	FunctionDeclaration       *FunctionDeclaration
	GeneratorDeclaration      *GeneratorDeclaration
	AsyncFunctionDeclaration  *AsyncFunctionDeclaration
	AsyncGeneratorDeclaration *AsyncGeneratorDeclaration
}

// FunctionDeclaration represents a FunctionDeclaration node.
type FunctionDeclaration struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	FunctionBody      *FunctionBody
}

// GeneratorDeclaration represents a GeneratorDeclaration node.
type GeneratorDeclaration struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	GeneratorBody     *GeneratorBody
}

// VariableDeclarationList represents a VariableDeclarationList node.
type VariableDeclarationList struct {
	VariableDeclarations []*VariableDeclaration
}

// VariableDeclaration represents a VariableDeclaration node.
type VariableDeclaration struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
	Initializer       *Initializer
}

// ForDeclaration represents a ForDeclaration node.
type ForDeclaration struct {
	LetOrConst *LetOrConst
	ForBinding *ForBinding
}

// LexicalDeclaration represents a LexicalDeclaration node.
type LexicalDeclaration struct {
	LetOrConst  *LetOrConst
	BindingList *BindingList
}

// ClassDeclaration represents a ClassDeclaration node.
type ClassDeclaration struct {
	BindingIdentifier *BindingIdentifier
	ClassTail         *ClassTail
}

// AsyncFunctionDeclaration represents a AsyncFunctionDeclaration node.
type AsyncFunctionDeclaration struct {
	BindingIdentifier *BindingIdentifier
	AsyncFunctionBody *AsyncFunctionBody
	FormalParameters  *FormalParameters
}

// AsyncGeneratorDeclaration represents a AsyncGeneratorDeclaration node.
type AsyncGeneratorDeclaration struct {
	BindingIdentifier  *BindingIdentifier
	AsyncGeneratorBody *AsyncGeneratorBody
	FormalParameters   *FormalParameters
}
