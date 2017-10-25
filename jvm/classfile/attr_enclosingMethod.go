package classfile

type AttrEnclosingMethod struct {
	cp          ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (attrMethod *AttrEnclosingMethod) ReadInfo(reader *ClassReader) uint32 {
	num := reader.ReadUint32()
	attrMethod.classIndex = reader.ReadUint16()
	attrMethod.methodIndex = reader.ReadUint16()
	return num
}
