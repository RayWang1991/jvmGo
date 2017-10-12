package classfile

type NameTypeInfo struct {
	nameIndex uint16 // index to the name
	typeIndex uint16 // index to type
}

func (n *NameTypeInfo) ReadInfo(reader *ClassReader) {
	n.nameIndex = reader.ReadUint16()
	n.typeIndex = reader.ReadUint16()
}
