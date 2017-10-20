package marea

import (
	cf "jvmGo/ch6/classfile"
	"jvmGo/ch6/utils"
)

func NewMethodRef(cp cf.ConstantPool, i *cf.MethodRefInfo, cls *Class) *MethodRef {
	return &MethodRef{*NewRef(cp, &(i.RefInfo), cls), nil}
}

func NewInterfaceMethodRef(cp cf.ConstantPool, i *cf.InterfaceMethodRefInfo, cls *Class) *InterfaceMethodRef {
	return &InterfaceMethodRef{*NewRef(cp, &(i.RefInfo), cls)}
}

type MethodRef struct {
	MemberRef
	method *Method
}

func (m *MethodRef) GetMethod() *Method {
	if m.method != nil {
		return m.method
	}
	c := m.ClassRef.Ref()
	m.method = c.LookUpMethod(m.name, m.desc)
	if m.method == nil {
		panic(utils.NoSuchMethodError)
	} else {
		// TODO check
	}
	return
}

type InterfaceMethodRef struct {
	MemberRef
}
