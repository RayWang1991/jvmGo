package classfile

// for methodHandle
type MethodTypeInfo struct {
	descIndex uint16 // index to utf8_info
}

func (m *MethodTypeInfo) ReadInfo(reader *ClassReader) {
	m.descIndex = reader.ReadUint16()
}

func (m *MethodTypeInfo) Desc(cp ConstantPool) string {
	return cp[m.descIndex].(*Utf8Info).val
}
