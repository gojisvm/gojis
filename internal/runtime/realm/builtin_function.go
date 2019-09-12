package realm

import (
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

// CreateBuiltinFunction creates a callable object, whose Call internal method will be the passed function fn.
// CreateBuiltinFunction is specified in 9.3.3.
func CreateBuiltinFunction(fn lang.NativeFunction, realm *Realm, proto lang.Value, internalSlotsList ...lang.StringOrSymbol) *lang.Object {
	if realm == nil {
		realm = CurrentRealm()
	}
	if proto == nil {
		proto = realm.GetIntrinsicObject(IntrinsicNameFunctionPrototype)
	}
	fobj := lang.ObjectCreate(proto, internalSlotsList...)
	fobj.Call = fn
	fobj.Realm = realm
	fobj.Extensible = true
	fobj.ScriptOrModule = lang.Null
	return fobj
}
