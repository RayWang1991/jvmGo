package classfile

type AttrUnsupported struct {
}

func (s *AttrUnsupported) ReadInfo(reader ClassReader) uint64 {
	num := reader.ReadUint64()
	reader.ReadBytes(uint(num))
	return num
}
