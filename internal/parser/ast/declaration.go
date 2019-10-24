package ast

type Declaration struct {
	HoistableDeclaration *HoistableDeclaration
	ClassDeclaration     *ClassDeclaration
	LexicalDeclaration   *LexicalDeclaration
}

type HoistableDeclaration struct {
	FunctionDeclaration       *FunctionDeclaration
	GeneratorDeclaration      *GeneratorDeclaration
	AsyncFunctionDeclaration  *AsyncFunctionDeclaration
	AsyncGeneratorDeclaration *AsyncGeneratorDeclaration
}

type FunctionDeclaration struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	FunctionBody      *FunctionBody
}

type GeneratorDeclaration struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	GeneratorBody     *GeneratorBody
}

type VariableDeclarationList struct {
	VariableDeclarations []*VariableDeclaration
}

type VariableDeclaration struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
	Initializer       *Initializer
}

type ForDeclaration struct {
	LetOrConst *LetOrConst
	ForBinding *ForBinding
}

type LexicalDeclaration struct {
	LetOrConst  *LetOrConst
	BindingList *BindingList
}

type ClassDeclaration struct {
	BindingIdentifier *BindingIdentifier
	ClassTail         *ClassTail
}

type AsyncFunctionDeclaration struct {
	BindingIdentifier *BindingIdentifier
	AsyncFunctionBody *AsyncFunctionBody
	FormalParameters  *FormalParameters
}

type AsyncGeneratorDeclaration struct {
	BindingIdentifier  *BindingIdentifier
	AsyncGeneratorBody *AsyncGeneratorBody
	FormalParameters   *FormalParameters
}
