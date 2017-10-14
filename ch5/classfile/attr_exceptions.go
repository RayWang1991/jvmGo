package classfile

type AttrExceptions struct {
	cp               ConstantPool
	exceptionIndexes []uint16 // indexes to constant pool, must be class info
}

func (attrExt *AttrExceptions) ReadInfo(reader *ClassReader) uint32 {
	num := reader.ReadUint32()
	indexNum := reader.ReadUint16()
	exceptions := make([]uint16, 0, indexNum)
	for i := uint16(0); i < indexNum; i++ {
		exceptions = append(exceptions, reader.ReadUint16())
	}
	attrExt.exceptionIndexes = exceptions
	return num
}
