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

func NewMethodRef(cp cf.ConstantPool, i *cf.MethodRefInfo, cls *Class) *MethodRef {
	return &MethodRef{*NewRef(cp, &(i.RefInfo), cls)}
}

func NewInterfaceMethodRef(cp cf.ConstantPool, i *cf.InterfaceMethodRefInfo, cls *Class) *InterfaceMethodRef {
	return &InterfaceMethodRef{*NewRef(cp, &(i.RefInfo), cls)}
}

type MethodRef struct {
	MemberRef
}

type InterfaceMethodRef struct {
	MemberRef
}
