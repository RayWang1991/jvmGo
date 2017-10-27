package marea

import (
	"jvmGo/jvm/cmn"
	"jvmGo/jvm/utils"
)

func isAccessableCls(want, wanted *Class) bool {
	if wanted.IsPublic() {
		return true
	}
	return want.PackageName() == wanted.PackageName()
}

func isAccessableField(want *Class, f *Field) bool {
	if f.IsPublic() {
		return true
	}
	if f.IsPrivate() {
		return want == f.Class()
	}
	if f.IsProtected() && IsDescandent(want, f.Class()) {
		return true
	}
	return want.PackageName() == f.Class().PackageName()
}

func isAccessableMethod(want *Class, m *Method) bool {
	if m.IsPublic() {
		return true
	}
	if m.IsPrivate() {
		return want == m.Class()
	}
	if m.IsProtected() && IsDescandent(want, m.Class()) {
		return true
	}
	return want.PackageName() == m.Class().PackageName()
}

// class hierarchy
// whether S is assignable to T
func IsAssignable(S, T *Class) bool {
	if !S.IsArray() {
		if S.IsInterface() {
			// S is interface type
			if !T.IsInterface() {
				// T is Class type
				return T.ClassName() == utils.CLASSNAME_Object
			} else {
				// T is Interface Type
				return IsDescandent(S, T)
			}
		} else {
			// S is Class type
			if !T.IsInterface() {
				// T is Class type
				return IsDescandent(S, T)
			} else {
				// T is Interface type
				return DoesImplement(S, T)
			}
		}
	} else {
		// s is an array type
		if !T.IsInterface() {
			if T.IsArray() {
				// T  array type
				sn := cmn.ElementName(S.ClassName())
				tn := cmn.ElementName(T.ClassName())
				if cmn.IsPrimitiveType(sn) {
					return tn == sn
				} else {
					if cmn.IsPrimitiveType(tn) {
						return false
					}
					Ss := S.DefineLoader().Load(sn) // element type for S
					Ts := S.DefineLoader().Load(tn) // element type for T
					return IsAssignable(Ss, Ts)
				}
			} else {
				// T is Class type
				return T.name == utils.CLASSNAME_Object
			}
		} else {
			// T is interface type
			return T.name == utils.CLASSNAME_Cloneable || T.name == utils.CLASSNAME_Serializable
		}
	}
}

// return the result that whether c is subclass of d or c is d
func IsDescandent(c, d *Class) bool {
	if d == nil {
		return false
	}
	for ; c != nil; c = c.superClass {
		if c == d {
			return true
		}
	}
	return false
}

// return the result that whether c implements interface d
func DoesImplement(c, d *Class) bool {
	for ; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if IsDescandent(i, d) {
				return true
			}
		}
	}
	return false
}
