package classfile

// Deprecation
type AttrDeprecated struct {
}

func (dep *AttrDeprecated) ReadInfo(reader ClassReader) uint64 {
	return reader.ReadUint64()
}

// Synthetic
type AttrSynthetic struct {
}

func (_ *AttrSynthetic) ReadInfo(reader ClassReader) uint64 {
	return reader.ReadUint64()
}

// Signature
type AttrSignature struct {
	cp  ConstantPool
	val string
}

func (sig *AttrSignature) ReadInfo(reader ClassReader) uint64 {
	num := reader.ReadUint64()
	index := reader.ReadUint16()
	sig.val = sig.cp[index].(*Utf8Info).val
	return num
}

// SourceFile
type AttrSourceFile struct {
	cp  ConstantPool
	val string
}

func (s *AttrSourceFile) ReadInfo(reader ClassReader) uint64 {
	num := reader.ReadUint64()
	index := reader.ReadUint16()
	s.val = s.cp[index].(*Utf8Info).val
	return num
}
