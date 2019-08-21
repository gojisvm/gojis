package snapshot

import (
	"fmt"
	"io"
	"reflect"
	"unsafe"

	endian "github.com/virtao/GoEndian"
)

func Load(bytes []byte, v interface{}, size uintptr) error {
	ptr := unsafe.Pointer(reflect.ValueOf(v).Pointer()) // #nosec
	data := (*(*[1<<31 - 1]byte)(ptr))[:size]

	copy(data, bytes)

	return nil
}

func Store(w io.Writer, v interface{}) error {
	val := reflect.ValueOf(v)
	ptr := unsafe.Pointer(val.Pointer()) // #nosec

	data := toBytes(ptr, val.Elem().Type())
	_, err := w.Write(data)
	if err != nil {
		return fmt.Errorf("Store: %v", err)
	}
	return nil
}

// toBytes returns a byte slice pointing to the memory segment
// where v points to. v has to be a pointer.
func toBytes(ptr unsafe.Pointer, t reflect.Type) []byte {
	data := bytesAt(ptr, t.Size())

	ptrs := getPointers(data, ptr, t)
	_ = ptrs

	return data
}

func bytesAt(p unsafe.Pointer, size uintptr) []byte {
	return (*(*[1<<31 - 1]byte)(p))[:size]
}

type nestedPtr struct {
	ptr unsafe.Pointer
	t   reflect.Type
}

// getPointers returns a map with offsets within v. These offsets
// are pointers pointing to some memory location.
// The returned map associates the memory offsets of the pointers
// with their destinations.
func getPointers(data []byte, ptr unsafe.Pointer, t reflect.Type) map[uintptr]nestedPtr {
	ptrs := make(map[uintptr]nestedPtr)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Ptr {
			tField := t.Field(i)
			offset := tField.Offset

			var encPtr uintptr
			const uintptrsize = unsafe.Sizeof(encPtr)
			ptrBytes := data[offset : offset+uintptrsize]
			switch unsafe.Sizeof(encPtr) {
			case 4:
				ptrs[offset] = nestedPtr{
					// this use of unsafe.Pointer should be correct
					ptr: unsafe.Pointer(uintptr(endian.Endian.Uint32(ptrBytes))), // #nosec
					t:   tField.Type,
				}
			case 8:
				ptrs[offset] = nestedPtr{
					// this use of unsafe.Pointer should be correct
					ptr: unsafe.Pointer(uintptr(endian.Endian.Uint64(ptrBytes))), // #nosec
					t:   tField.Type,
				}
			default:
				panic("Unknown uintptr size: " + string(uintptrsize))
			}
		}
	}

	return ptrs
}
