package classfile

import "math"

// class info
type ClassInfo struct {
	nameIndex uint16 // index to constant pool, a utf8_info to represent the Full Qualified class name
}

func (c *ClassInfo) ReadInfo(reader *ClassReader) {
	c.nameIndex = reader.ReadUint16()
}

type Utf8Info struct {
	val string // []byte data
}

func (u *Utf8Info) ReadInfo(reader *ClassReader) {
	length := reader.ReadUint16() // length for the utf8 info in bytes
	bs := reader.ReadBytes(uint(length))
	u.val = decodeMUTF8(bs)
}

type IntegerInfo struct {
	val int32 // integer
}

func (i *IntegerInfo) ReadInfo(reader *ClassReader) {
	i.val = int32(reader.ReadUint32())
}

type FloatInfo struct {
	val float32 // float
}

func (f *FloatInfo) ReadInfo(reader *ClassReader) {
	f.val = math.Float32frombits(reader.ReadUint32())
}

type LongInfo struct {
	val int64 // long
}

func (l *LongInfo) ReadInfo(reader *ClassReader) {
	l.val = int64(reader.ReadUint64())
}

type DoubleInfo struct {
	val float64 // double
}

func (d *DoubleInfo) ReadInfo(reader *ClassReader) {
	d.val = math.Float64frombits(reader.ReadUint64())
}

type StringInfo struct {
	index uint16 // index to constant pool, a utf8_info
}

func (s *StringInfo) ReadInfo(reader *ClassReader) {
	s.index = reader.ReadUint16()
}

type RefInfo struct {
	classIndex    uint16 // index to the class info
	nameTypeIndex uint16 // index to the name and type info
}

func (r *RefInfo) ReadInfo(reader *ClassReader) {
	r.classIndex = reader.ReadUint16()
	r.nameTypeIndex = reader.ReadUint16()
}

type FieldRefInfo struct {
	RefInfo
}

type MethodRefInfo struct {
	RefInfo
}

type InterfaceMethodRefInfo struct {
	RefInfo
}

type NameTypeInfo struct {
	nameIndex uint16 // index to the name
	typeIndex uint16 // index to type
}

func (n *NameTypeInfo) ReadInfo(reader *ClassReader) {
	n.nameIndex = reader.ReadUint16()
	n.typeIndex = reader.ReadUint16()
}

type MethodHandleInfo struct {
	refKind  byte   // [1-9]
	refIndex uint16 // index to method ref
}

func (m *MethodHandleInfo) ReadInfo(reader *ClassReader) {
	m.refKind = reader.ReadUint8()
	m.refIndex = reader.ReadUint16()
}

type MethodTypeInfo struct {
	descIndex uint16 // index to utf8_info
}

func (m *MethodTypeInfo) ReadInfo(reader *ClassReader) {
	m.descIndex = reader.ReadUint16()
}

type InvokeDynamic_Info struct {
	bootstrapMethodAttrIndex uint16 // index to bootstrap_methods[] ?
	nameTypeIndex            uint16 // index to Name and Type info
}

func (i *InvokeDynamic_Info) ReadInfo(reader *ClassReader) {
	i.bootstrapMethodAttrIndex = reader.ReadUint16()
	i.nameTypeIndex = reader.ReadUint16()
}
