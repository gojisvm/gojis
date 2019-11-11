package parser

// Error represents an error that occurred while parsing. Inside the parser
// package, an instance of this struct can be panicked. The panic will be
// recovered from, and will be propagated to outside the package as a common
// error. If anything else is panicked, the panic will not be recovered from.
type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}
