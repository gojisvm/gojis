package snapshot

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"unsafe"

	endian "github.com/virtao/GoEndian"
	capnp "zombiezen.com/go/capnproto2"
)

const uintptrsize = unsafe.Sizeof(uintptr(0)) // #nosec

// Load unsafely loads the given bytes into the memory location
// of v. size must be equal to the result of
//
//	unsafe.Sizeof(v)
func Load(b []byte, v interface{}, size uintptr) error {
	ptr := unsafe.Pointer(reflect.ValueOf(v).Pointer()) // #nosec
	data := bytesAt(ptr, size)

	msg, err := capnp.NewDecoder(bytes.NewReader(b)).Decode()
	if err != nil {
		return fmt.Errorf("Load: %v", err)
	}

	sn, err := ReadRootSnapshot(msg)

	root, err := sn.Nested()
	if err != nil {
		return fmt.Errorf("Load(root): %v", err)
	}

	err = loadTo(&data, root)
	return err
}

func loadTo(dst *[]byte, nst Nested) error {
	data, err := nst.Data()
	if err != nil {
		return fmt.Errorf("loadTo(data): %v", err)
	}
	copy(*dst, data)

	ptrList, err := nst.Pointers()
	if err != nil {
		return fmt.Errorf("loadTo(ptrList): %v", err)
	}

	for i := 0; i < ptrList.Len(); i++ {
		ptr := ptrList.At(i)
		offset := uintptr(ptr.Offset())
		target, err := ptr.Target()
		if err != nil {
			return fmt.Errorf("loadTo(target): %v", err)
		}

		nextData := make([]byte, target.Len())
		copy((*dst)[offset:offset+uintptrsize], (*[uintptrsize]byte)(unsafe.Pointer(&((*reflect.StringHeader)(unsafe.Pointer(&nextData)).Data)))[:]) // #nosec
		err = loadTo(&nextData, target)
		if err != nil {
			return fmt.Errorf("loadTo(loadTo): %v", err)
		}
	}

	return nil
}

// Store stores the given interface to a writer in a proprietary, GOARCH and GOOS
// dependent format, that will not be documented further.
func Store(w io.Writer, v interface{}) error {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return fmt.Errorf("Store: %v", err)
	}

	sn, err := storeToCapnp(v, seg)
	_ = msg.SetRoot(sn)

	err = encodeCapnp(w, msg)
	if err != nil {
		return fmt.Errorf("Store: %v", err)
	}

	return nil
}

type internalSnapshot struct {
	nested *internalNested
}

type internalNested struct {
	data    []byte
	targets []*internalPointer
}

type internalPointer struct {
	offset int64
	target *internalNested
}

func storeToCapnp(v interface{}, s *capnp.Segment) (Snapshot, error) {
	snapshot, err := NewRootSnapshot(s)
	if err != nil {
		return snapshot, err
	}

	sn := storeInternal(v)

	nested, err := storeNestedInSegment(sn.nested, s)
	if err != nil {
		return snapshot, fmt.Errorf("storeToCapnp: %v", err)
	}

	_ = snapshot.SetNested(nested)

	return snapshot, err
}

func storeNestedInSegment(internalNested *internalNested, s *capnp.Segment) (Nested, error) {
	nested, err := NewNested(s)
	if err != nil {
		return nested, fmt.Errorf("storeNestedInSegment: %v", err)
	}

	pointers, err := NewPointer_List(s, int32(len(internalNested.targets)))
	if err != nil {
		return nested, fmt.Errorf("storeNestedInSegment: %v", err)
	}

	for i, targets := range internalNested.targets {
		pointer, err := NewPointer(s)
		if err != nil {
			return nested, fmt.Errorf("storeNestedInSegment: %v", err)
		}

		next, err := storeNestedInSegment(targets.target, s)

		pointer.SetOffset(targets.offset)
		_ = pointer.SetTarget(next)

		_ = pointers.Set(i, pointer)
	}

	_ = nested.SetData(internalNested.data)
	nested.SetLen(int64(len(internalNested.data)))
	_ = nested.SetPointers(pointers)

	return nested, nil
}

func encodeCapnp(w io.Writer, msg *capnp.Message) error {
	err := capnp.NewEncoder(w).Encode(msg)
	if err != nil {
		return fmt.Errorf("encodeCapnp: %v", err)
	}
	return nil
}

func storeInternal(v interface{}) *internalSnapshot {
	sn := &internalSnapshot{}

	val := reflect.ValueOf(v)
	ptr := unsafe.Pointer(val.Pointer()) // #nosec

	data := toInternalNested(ptr, val.Elem().Type())
	sn.nested = data

	return sn
}

func toInternalNested(ptr unsafe.Pointer, t reflect.Type) *internalNested {
	result := &internalNested{}
	result.data = bytesAt(ptr, t.Size())

	ptrs := getPointers(result.data, ptr, t)
	for ptrOffset, nested := range ptrs {
		next := toInternalNested(nested.ptr, nested.t)

		result.targets = append(result.targets, &internalPointer{
			offset: int64(uintptr(ptrOffset)),
			target: next,
		})
	}

	return result
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
// relative to ptr with their destinations.
/* #nosec */
func getPointers(data []byte, ptr unsafe.Pointer, t reflect.Type) map[uintptr]nestedPtr {
	ptrs := make(map[uintptr]nestedPtr)

	if t.Kind() != reflect.Struct {
		// not a pointer we can follow, abort
		return ptrs
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Ptr {
			tField := t.Field(i)
			offset := tField.Offset

			switch uintptrsize {
			case 4:
				ptrs[offset] = nestedPtr{
					ptr: unsafe.Pointer(uintptr(endian.Endian.Uint32(data[offset : offset+uintptrsize]))),
					t:   tField.Type,
				}
			case 8:
				ptrs[offset] = nestedPtr{
					ptr: unsafe.Pointer(uintptr(endian.Endian.Uint64(data[offset : offset+uintptrsize]))),
					t:   tField.Type,
				}
			default:
				panic("Unknown uintptr size: " + string(uintptrsize))
			}
		}
	}

	return ptrs
}
