package classfile

type RefInfo struct {
	classIndex    uint16 // index to the class info
	nameTypeIndex uint16 // index to the name and type info
}

func (r *RefInfo) ReadInfo(reader *ClassReader) {
	r.classIndex = reader.ReadUint16()
	r.nameTypeIndex = reader.ReadUint16()
}

type FieldRefInfo struct {
	RefInfo
}

type MethodRefInfo struct {
	RefInfo
}

type InterfaceMethodRefInfo struct {
	RefInfo
}
