package classfile

type MethodTypeInfo struct {
	descIndex uint16 // index to utf8_info
}

func (m *MethodTypeInfo) ReadInfo(reader *ClassReader) {
	m.descIndex = reader.ReadUint16()
}