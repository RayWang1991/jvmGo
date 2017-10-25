package marea

import "jvmGo/jvm/classfile"

type NameTypeInfo struct {
	nameIndex uint16 // index to the name
	typeIndex uint16 // index to type
	name, tp  string
}

func NewNameTypeInfo(cp classfile.ConstantPool, info *classfile.NameTypeInfo) *NameTypeInfo {
	nti := &NameTypeInfo{}
	nti.name = info.Name(cp)
	nti.tp = info.Type(cp)
	return nti
}
