package marea

import (
	cf "jvmGo/jvm/classfile"
)

type MemberRef struct {
	ClassRef
	name string
	desc string
}

func (m *MemberRef) Name() string {
	return m.name
}

func (m *MemberRef) Desc() string {
	return m.desc
}

func NewRef(cp cf.ConstantPool, mi *cf.RefInfo, cls *Class) *MemberRef {
	cr := NewClassRef(mi.GetClassName(cp), cls)
	n, t := mi.GetNameAndType(cp)
	m := &MemberRef{
		ClassRef: *cr,
		name:     n,
		desc:     t,
	}
	return m
}
