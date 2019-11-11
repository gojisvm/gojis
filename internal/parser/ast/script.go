package ast

// Script represents a Script node.
type Script struct {
	ScriptBody *ScriptBody
}

// ScriptBody represents a ScriptBody node.
type ScriptBody struct {
	StatementList *StatementList
}
