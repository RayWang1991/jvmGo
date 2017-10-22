package classfile

type ClassInfo struct {
	nameIndex uint16 // index to constant pool, a utf8_info to represent the Full Qualified class name
}

func (c *ClassInfo) ReadInfo(reader *ClassReader) {
	c.nameIndex = reader.ReadUint16()
}
