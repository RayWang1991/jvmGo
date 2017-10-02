package classfile

type AttrLocalVarTable struct {
	localVarTable []AttrLocalVarEntry
}

func (table *AttrLocalVarTable) ReadInfo(reader *ClassReader) uint32 {
	num := reader.ReadUint32()
	eNum := reader.ReadUint16()
	entries := make([]AttrLocalVarEntry, 0, eNum)
	for i := uint16(0); i < eNum; i++ {
		entries = append(entries, AttrLocalVarEntry{
			reader.ReadUint16(),
			reader.ReadUint16(),
			reader.ReadUint16(),
			reader.ReadUint16(),
			reader.ReadUint16(),
		})
	}
	table.localVarTable = entries
	return num
}

type AttrLocalVarEntry struct {
	startPC   uint16
	length    uint16
	nameIndex uint16
	descIndex uint16
	index     uint16
}
