package classfile

type MethodHandleInfo struct {
	refKind  byte   // [1-9]
	refIndex uint16 // index to method ref
}

func (m *MethodHandleInfo) ReadInfo(reader *ClassReader) {
	m.refKind = reader.ReadUint8()
	m.refIndex = reader.ReadUint16()
}