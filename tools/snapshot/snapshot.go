package snapshot

import (
	"io"
	"log"
	"reflect"
	"unsafe"
)

func Load(bytes []byte, v interface{}, size uintptr) error {
	ptr := unsafe.Pointer(reflect.ValueOf(v).Pointer()) // #nosec
	data := (*(*[1<<31 - 1]byte)(ptr))[:size]

	copy(data, bytes)

	return nil
}

func Store(w io.Writer, v interface{}) {
	data := toBytes(v)
	w.Write(data)
}

// toBytes returns a byte slice pointing to the memory segment
// where v points to. v has to be a pointer.
func toBytes(v interface{}) []byte {
	// convert the value to a byte slice
	val := reflect.ValueOf(v)
	size := val.Elem().Type().Size()
	log.Printf("size: %v\n", size)
	ptr := unsafe.Pointer(val.Pointer())
	data := (*(*[1<<31 - 1]byte)(ptr))[:size] // #nosec

	// pointers are not yet supported

	return data
}

// getPointers returns a map with offsets within v. These offsets
// are pointers pointing to some memory location.
// The returned map associates the memory offsets of the pointers
// with their destinations.
func getPointers(v interface{}) map[uintptr]unsafe.Pointer {
	panic("TODO")
}
