package classfile

type AttrMethodParameters struct {
	cp     ConstantPool
	params []AttrMethodParameterEntry
}

func (ps *AttrMethodParameters) ReadInfo(reader ClassReader) uint64 {
	num := reader.ReadUint64()
	n := reader.ReadUint8()
	params := make([]AttrMethodParameterEntry, 0, n)
	for i := uint8(0); i < n; i++ {
		entry := AttrMethodParameterEntry{
			reader.ReadUint16(),
			reader.ReadUint16(),
		}
		params = append(params, entry)
	}
	ps.params = params
	return num
}

type AttrMethodParameterEntry struct {
	nameIndex uint16
	accFlags  uint16
}
