package classfile

type AttrEnclosingMethod struct {
	cp ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (attrMethod *AttrEnclosingMethod) ReadInfo(reader ClassReader) uint64 {
	num := reader.ReadUint64()
	attrMethod.classIndex = reader.ReadUint16()
	attrMethod.methodIndex = reader.ReadUint16()
	return num
}
