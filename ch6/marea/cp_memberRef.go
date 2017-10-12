package marea

import cf "jvmGo/ch6/classfile"

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

func NewRef(cp cf.ConstantPool, mi *cf.RefInfo) *MemberRef {
	cr := NewClassRef(mi.GetClassName(cp))
	n, t := mi.GetNameAndType(cp)
	m := &MemberRef{
		ClassRef: *cr,
		name:     n,
		desc:     t,
	}
	return m
}

func NewFieldRef(cp cf.ConstantPool, i *cf.FieldRefInfo) *FieldRef {
	return &FieldRef{*NewRef(cp, &(i.RefInfo))}
}

func NewMethodRef(cp cf.ConstantPool, i *cf.MethodRefInfo) *MethodRef {
	return &MethodRef{*NewRef(cp, &(i.RefInfo))}
}

func NewInterfaceMethodRef(cp cf.ConstantPool, i *cf.InterfaceMethodRefInfo) *InterfaceMethodRef {
	return &InterfaceMethodRef{*NewRef(cp, &(i.RefInfo))}
}

type FieldRef struct {
	MemberRef
}

type MethodRef struct {
	MemberRef
}

type InterfaceMethodRef struct {
	MemberRef
}
