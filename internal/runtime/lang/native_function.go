package lang

import "github.com/gojisvm/gojis/internal/runtime/errors"

// NativeFunction is a type alias for Go functions that represent an ECMAScript
// function algorithm. It takes a 'this' value, which will be interpreted as the
// receiver, and zero or more arguments.
//
// A function can return one argument, which may be Null or Undefined, and an
// error. If an error is returned, the function call is interpreted as a
// function call with abnormal completion, and the returned error will be
// interpreted as a thrown error. The runtime must handle this correctly.
type NativeFunction func(this Value, args ...Value) (Value, errors.Error)
