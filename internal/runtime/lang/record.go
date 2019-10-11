package lang

// Record is an object that holds fields.
// The type of the fields as well as the name of such fields
// are not specified by the record, but by the entity that is
// using the record.
type Record struct {
	fields map[interface{}]interface{}
}

// NewRecord creates a new record with no fields.
func NewRecord() *Record {
	r := new(Record)
	r.fields = make(map[interface{}]interface{})
	return r
}

// GetField returns the value of the field associated with
// the given name, and a flag indicating whether such a field
// is held by the record.
func (r *Record) GetField(n string) (interface{}, bool) {
	val, ok := r.fields[n]
	return val, ok
}

// SetField associates a value with a given field name.
// If the field already exists, it will be overwritten.
func (r *Record) SetField(n string, val interface{}) {
	r.fields[n] = val
}

// Type returns lang.TypeInternal.
func (r *Record) Type() Type { return TypeInternal }

// Value returns the Record itself.
func (r *Record) Value() interface{} { return r }
