package classfile

type AttrBootstrapMethods struct {
	cp              ConstantPool
	bootMethodTable []AttrBootstrapMethodEntry
}

func (bsm *AttrBootstrapMethods) ReadInfo(reader *ClassReader) uint32 {
	num := reader.ReadUint32()
	tableN := reader.ReadUint16()
	table := make([]AttrBootstrapMethodEntry, 0, tableN)
	for i := uint16(0); i < tableN; i++ {
		entry := AttrBootstrapMethodEntry{}
		entry.ReadInfo(reader)
		table = append(table, entry)
	}
	bsm.bootMethodTable = table
	return num
}

type AttrBootstrapMethodEntry struct {
	methodRef uint16
	args      []uint16
	// index to constant pool, available for string_info, class_info, integer_info,
	// long, float, double, method handle, method type
}

func (entry *AttrBootstrapMethodEntry) ReadInfo(reader *ClassReader) {
	entry.methodRef = reader.ReadUint16()
	n := reader.ReadUint16()
	args := make([]uint16, 0, n)
	for i := uint16(0); i < n; i++ {
		args = append(args, reader.ReadUint16())
	}
}
