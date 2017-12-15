package classfile

type AttrRuntimeVisibleParameterAnnotationsAttributeData struct {
	data []byte
}

// TODO
func (attr *AttrRuntimeVisibleParameterAnnotationsAttributeData) ReadInfo(reader *ClassReader) uint32 {
	length := reader.ReadUint32()
	attr.data = reader.ReadBytes(uint(length))
	return length
}
