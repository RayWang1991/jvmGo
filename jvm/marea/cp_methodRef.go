package marea

import (
	cf "jvmGo/jvm/classfile"
	"jvmGo/jvm/utils"
)

func NewMethodRef(cp cf.ConstantPool, i *cf.MethodRefInfo, cls *Class) *MethodRef {
	return &MethodRef{*NewRef(cp, &(i.RefInfo), cls), nil}
}

type MethodRef struct {
	MemberRef
	method *Method
}

func (m *MethodRef) GetRef() *MemberRef {
	return &m.MemberRef
}

func (m *MethodRef) GetMethod() *Method {
	if m.method != nil {
		return m.method
	}
	c := m.ClassRef.Ref()
	if c.IsInterface() {
		panic(utils.IncompatibleClassChangeError)
	}
	//fmt.Printf("name : %s, desc : %s\n", m.name, m.desc)
	m.method = c.LookUpMethod(m.name, m.desc)
	m.method = LookUpMethodVirtual(c, m.from, m.name, m.desc)
	return m.method
}
