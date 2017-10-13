package cmn

import "jvmGo/ch6/marea"

// An Object represents a general reference type, for class type, interface type, and array type
type Object interface{
	Class()
}

type Object1 struct {
	// TODO
	marea.Class
}
