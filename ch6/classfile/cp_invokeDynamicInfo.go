package classfile

type InvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16 // index to bootstrap_methods[] ?
	nameTypeIndex            uint16 // index to Name and Type info
}

func (i *InvokeDynamicInfo) ReadInfo(reader *ClassReader) {
	i.bootstrapMethodAttrIndex = reader.ReadUint16()
	i.nameTypeIndex = reader.ReadUint16()
}

func (i *InvokeDynamicInfo) NameType(cp ConstantPool) (string, string) {
	nt := cp[i.nameTypeIndex].(*NameTypeInfo)
	return nt.Name(cp), nt.Type(cp)
}

func (i *InvokeDynamicInfo) BstmI() uint16 {
	return i.bootstrapMethodAttrIndex
}
