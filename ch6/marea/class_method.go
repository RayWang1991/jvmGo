package marea

import (
	"jvmGo/ch6/cmn"
	"jvmGo/ch6/classfile"
)

type Method struct {
	ClassMember
	max
	code       []byte
	exceptions []string //strings of exception class name
}

func NewMethod(info *classfile.MethodInfo) *Method {
	m := &Method{}
	m.name = info.Name()
	m.flags = info.AccFlags()
	m.desc = info.Description()
	m.code = info.GetCodeAttr()
	return m.
}

// access methods
func (m *Method) IsPublic() bool {
	return cmn.IsPublic(m.flags)
}
func (m *Method) IsPrivate() bool {
	return cmn.IsPrivate(m.flags)
}
func (m *Method) IsProtected() bool {
	return cmn.IsProtected(m.flags)
}
func (m *Method) IsFinal() bool {
	return cmn.IsFinal(m.flags)
}
func (m *Method) IsSynchronized() bool {
	return cmn.IsSynchronized(m.flags)
}
func (m *Method) IsVarargs() bool {
	return cmn.IsVarargs(m.flags)
}
func (m *Method) IsNative() bool {
	return cmn.IsNative(m.flags)
}
func (m *Method) IsAbstract() bool {
	return cmn.IsAbstract(m.flags)
}
