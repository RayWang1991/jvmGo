package classfile

type AttrConstantValue struct {
	cp    ConstantPool
	index uint16 // index to constant pool, pointing to a constant value
}

func (attrConst *AttrConstantValue) ReadInfo(reader *ClassReader) uint32 {
	num := reader.ReadUint32()
	attrConst.index = reader.ReadUint16()
	return num
}
