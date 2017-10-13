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

func (r *RefInfo) GetClassName(cp ConstantPool) string {
	return cp[r.classIndex].(*ClassInfo).ClassName(cp)
}

func (r *RefInfo) GetNameAndType(cp ConstantPool) (n, t string) {
	info := cp[r.nameTypeIndex].(*NameTypeInfo)
	n = info.Name(cp)
	t = info.Name(cp)
	return
}

func (r *RefInfo) NameTypeInfo(cp ConstantPool) *NameTypeInfo {
	return cp[r.nameTypeIndex].(*NameTypeInfo)
}

// get class info for field reference
func (r *RefInfo) ClassInfo(cp ConstantPool) *ClassInfo {
	return cp[r.classIndex].(*ClassInfo)
}
