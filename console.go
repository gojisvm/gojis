package gojis

import "io"

// Console is an io.Writer that is also an Object.
type Console interface {
	io.Writer
	Object
}

type console struct {
	io.Writer
}

// NewConsole creates a ready to use console that can be used inside a VM.
func NewConsole(w io.Writer) Console {
	c := new(console)
	c.Writer = w
	return c
}

func (c *console) Lookup(string) Object {
	panic("TODO")
}

func (c *console) SetFunction(string, func(Args) Object) {
	panic("TODO")
}

func (c *console) CallWithArgs(...interface{}) (Object, error) {
	panic("TODO")
}

func (c *console) SetObject(string, Object) {
	panic("TODO")
}

func (c *console) IsUndefined() bool {
	panic("TODO")
}

func (c *console) IsNull() bool {
	panic("TODO")
}

func (c *console) IsFunction() bool {
	panic("TODO")
}

func (c *console) Type() Type {
	panic("TODO")
}

func (c *console) Value() interface{} {
	panic("TODO")
}
