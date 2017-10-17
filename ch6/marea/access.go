package marea

import "jvmGo/ch6/utils"

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

// class hierarchy
// whether S is assignable to T
func IsAssignable(S, T *Class) bool {
	if !S.IsArray() {
		if T.IsInterface() {
			if !S.IsInterface() {
				return S.ClassName() == utils.CLASSNAME_Object
			} else {
				return IsDescandent(S, T)
			}
		} else {
			return DoesImplement(S, T)
		}
	} else {
		// TODO array
		panic("todo array")
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
