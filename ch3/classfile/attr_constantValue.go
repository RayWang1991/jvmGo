package classfile

type AttrConstantValue struct {
	cp    ConstantPool
	index uint16 // index to constant pool, pointing to a constant value
}

func (attrConst *AttrConstantValue) ReadInfo(reader ClassReader) uint64 {
	num := reader.ReadUint64()
	attrConst.index = reader.ReadUint16()
	return num
}
