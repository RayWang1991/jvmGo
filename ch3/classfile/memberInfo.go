package classfile

type MemberInfo struct {
	cp          ConstantPool
	accessFlags uint16
	nameIndex   uint16
	descIndex   uint16
	attrs       []AttrInfo
}

func (m *MemberInfo) ReadInfo(reader *ClassReader) {
	m.accessFlags = reader.ReadUint16()
	m.nameIndex = reader.ReadUint16()
	m.descIndex = reader.ReadUint16()
	n := reader.ReadUint16()
	attrs := make([]AttrInfo, 0, n)
	for i := uint16(0); i < n; i++ {
		attr := NewAttributeInfo(reader, m.cp)
		attr.ReadInfo(reader)
		attrs = append(attrs, attr)
	}
	m.attrs = attrs
}

type FieldInfo struct {
	MemberInfo
}

type MethodInfo struct {
	MemberInfo
}
