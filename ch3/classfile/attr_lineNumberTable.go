package classfile

type AttrLineNumberTable struct {
	lineNumberTable []AttrLineNumberEntry
}

type AttrLineNumberEntry struct {
	startPC    uint16
	lineNumber uint16
}

func (table *AttrLineNumberTable) ReadInfo(reader ClassReader) uint64 {
	num := reader.ReadUint64()
	lineNum := reader.ReadUint16()
	entries := make([]AttrLineNumberEntry, 0, lineNum)
	for i := uint16(0); i < lineNum; i++ {
		entries = append(entries, AttrLineNumberEntry{
			reader.ReadUint16(),
			reader.ReadUint16(),
		})
	}
	table.lineNumberTable = entries
	return num
}
