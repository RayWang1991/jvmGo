package classfile

type StringInfo struct {
	index uint16 // index to constant pool, a utf8_info
}

func (s *StringInfo) ReadInfo(reader *ClassReader) {
	s.index = reader.ReadUint16()
}
