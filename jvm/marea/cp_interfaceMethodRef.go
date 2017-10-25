package marea

import (
	cf "jvmGo/jvm/classfile"
	"jvmGo/jvm/utils"
)

func NewInterfaceMethodRef(cp cf.ConstantPool, i *cf.InterfaceMethodRefInfo, cls *Class) *InterfaceMethodRef {
	return &InterfaceMethodRef{*NewRef(cp, &(i.RefInfo), cls), nil}
}

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func (m *InterfaceMethodRef) GetRef() *MemberRef {
	return &m.MemberRef
}

func (m *InterfaceMethodRef) GetMethod() *Method {
	if m.method != nil {
		return m.method
	}
	c := m.ClassRef.Ref()
	if !c.IsInterface() {
		panic(utils.IncompatibleClassChangeError)
	}
	m.method = LookUpMethodVirtual(c, m.from, m.name, m.desc)
	return m.method
}
