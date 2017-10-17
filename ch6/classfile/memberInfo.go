package classfile

type MemberInfo struct {
	cp          ConstantPool
	accessFlags uint16
	nameIndex   uint16
	descIndex   uint16
	attrs       []AttrInfo
}

func (m *MemberInfo) ReadInfo(reader *ClassReader) {
	m.accessFlags = reader.ReadUint16()
	m.nameIndex = reader.ReadUint16()
	m.descIndex = reader.ReadUint16()
	n := reader.ReadUint16()
	attrs := make([]AttrInfo, 0, n)
	for i := uint16(0); i < n; i++ {
		attr := NewAttributeInfo(reader, m.cp)
		attr.ReadInfo(reader)
		attrs = append(attrs, attr)
	}
	m.attrs = attrs
}

func (m *MemberInfo) Name() string {
	return m.cp.GetUTF8(m.nameIndex)
}

func (m *MemberInfo) Description() string {
	return m.cp.GetUTF8(m.descIndex)
}

func (m *MemberInfo) AccFlags() uint16 {
	return m.accessFlags
}

func (m *MemberInfo) ConstantPool() ConstantPool {
	return m.cp
}

type FieldInfo struct {
	MemberInfo
	slotNum uint8
}

// index to constant value in constant pool, -1 on non
func (f *FieldInfo) GetConstantValueIndex() int {
	for _, a := range f.attrs {
		if a, ok := a.(*AttrConstantValue); ok {
			return int(a.index)
		}
	}
	return -1
}

func (f *FieldInfo) SlotNum() uint8 {
	return f.slotNum
}

type MethodInfo struct {
	MemberInfo
}

func (m *MethodInfo) GetCodeAttr() *AttrCode {
	for _, a := range m.attrs {
		if a, ok := a.(*AttrCode); ok {
			return a
		}
	}
	return nil
}
