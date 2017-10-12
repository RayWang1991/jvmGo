package marea

import cf "jvmGo/ch6/classfile"

type MethodRef struct {
	ClassRef
	name string
	desc string
}

func (m *MethodRef) Name() string {
	return m.name
}

func (m *MethodRef) Desc() string {
	return m.desc
}

func NewMethodRef(from *Class, cp cf.ConstantPool, mi *cf.MethodRefInfo) *MethodRef {
	cr := NewClassRef(from, mi.GetClassName(cp))
	n, t := mi.GetNameAndType(cp)
	m := &MethodRef{
		ClassRef: *cr,
		name:     n,
		desc:     t,
	}
	return m
}
