package snapshot

import "io"

func Load(r io.Reader, v interface{}) error {
	panic("TODO")
}

func Store(w io.Writer, v interface{}) error {
	panic("TODO")
}

// toBytes returns a byte slice pointing to the memory segment
// where v points to. v has to be a pointer.
func toBytes(v interface{}) []byte {
	panic("TODO")
}

// getPointers returns a map with offsets within v. These offsets
// are pointers pointing to some memory location.
// The returned map associates the memory offsets of the pointers
// with their destinations.
func getPointers(v interface{}) map[uintptr]uintptr {
	panic("TODO")
}
