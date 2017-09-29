package classfile

type AttrUnsupported struct {
}

func (s *AttrUnsupported) ReadInfo(reader *ClassReader) uint32 {
	num := reader.ReadUint32()
	reader.ReadBytes(uint(num))
	return num
}
