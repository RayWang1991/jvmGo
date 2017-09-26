package classFile

type ClassInfo struct {
	nameIndex uint16 // index to constant pool, a utf8_info to represent the Full Qualified class name
}

func (self *ClassInfo) ReadInfo(reader ClassReader) {
	self.nameIndex = reader.ReadUint16()
}

type Utf8Info struct {
	val string // []byte data
}

func (self *Utf8Info) ReadInfo(reader ClassReader) {
	length := reader.ReadUint16() // length for the utf8 info in bytes
	bs := reader.ReadBytes(uint(length))
	self.val = decodeMUTF8(bs)
}

type IntegerInfo struct {
	val int32 // integer
}

func (self *IntegerInfo) ReadInfo(reader ClassReader) {
	self.val = int32(reader.ReadUint32())
}

type FloatInfo struct {
	val float32 // float
}

func (self *FloatInfo) ReadInfo(reader ClassReader) {
	self.val = float32(reader.ReadUint32())
}

type LongInfo struct {
	val int64 // long
}

func (self *LongInfo) ReadInfo(reader ClassReader) {
	self.val = int64(reader.ReadUint64())
}

type DoubleInfo struct {
	val float64 // double
}

func (self *DoubleInfo) ReadInfo(reader ClassReader) {
	self.val = float64(reader.ReadUint64())
}

type StringInfo struct {
	index uint16 // index to constant pool, a utf8_info
}

func (self *StringInfo) ReadInfo(reader ClassReader) {
	self.index = reader.ReadUint16()
}

type RefInfo struct{
	classIndex    uint16 // index to the class info
	nameTypeIndex uint16 // index to the name and type info
}

func (self *RefInfo) ReadInfo(reader ClassReader) {
	self.classIndex = reader.ReadUint16()
	self.nameTypeIndex = reader.ReadUint16()
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

func (self *NameTypeInfo) ReadInfo(reader ClassReader) {
	self.nameIndex = reader.ReadUint16()
	self.typeIndex = reader.ReadUint16()
}

type MethodHandleInfo struct {
	refKind  byte   // [1-9]
	refIndex uint16 // index to method ref
}

func (self *MethodHandleInfo) ReadInfo(reader ClassReader) {
	self.refKind = reader.ReadUint8()
	self.refIndex = reader.ReadUint16()
}

type MethodTypeInfo struct {
	descIndex uint16 // index to utf8_info
}

func (self *MethodTypeInfo) ReadInfo(reader ClassReader) {
	self.descIndex = reader.ReadUint16()
}

type InvokeDynamic_Info struct {
	bootstrapMethodAttrIndex uint16 // index to bootstrap_methods[] ?
	nameTypeIndex            uint16 // index to Name and Type info
}

func (self *InvokeDynamic_Info) ReadInfo(reader ClassReader) {
	self.bootstrapMethodAttrIndex = reader.ReadUint16()
	self.nameTypeIndex = reader.ReadUint16()
}

