package classfile

import (
	"fmt"
	"bytes"
	"jvmGo/ch6/instructions"
)

/*
Code {
  u2 attr_name_index
  u4 attr_length
  u2 max_stack
  u2 max_locals
  u4 code_length
  u1 code[code_length]
  u2 exception_table_length
  {
    u2 start_pc
    u2 end_pc
    u2 handler_pc
    u2 catch_type
   } exception_table[exception_table_length]
  u2 attributes_count;
  attribute_info attributes[attributes_count]
}
 */

type AttrCode struct {
	cp        ConstantPool
	maxStack  uint16
	maxLocals uint16
	code      []byte
	excTable  []AttrExceptionTableEntry
	attrs     []AttrInfo //LineNumberTable, LocalVariableTable, LocalVariableTypeTable, StackMapTable
}

// getter
func (code *AttrCode) MaxStack() uint16 {
	return code.maxStack
}

func (code *AttrCode) MaxLocals() uint16 {
	return code.maxLocals
}

func (code *AttrCode) Code() []byte {
	return code.code
}

func (code *AttrCode) ReadInfo(reader *ClassReader) uint32 {
	num := reader.ReadUint32()
	code.maxStack = reader.ReadUint16()
	code.maxLocals = reader.ReadUint16()
	codeLen := reader.ReadUint32()
	code.code = reader.ReadBytes(uint(codeLen))

	entryN := reader.ReadUint16()
	code.excTable = make([]AttrExceptionTableEntry, entryN)
	for _, entry := range code.excTable {
		entry.ReadInfo(reader)
	}

	attrN := reader.ReadUint16()
	code.attrs = make([]AttrInfo, attrN)
	for i := range code.attrs {
		start := reader.length()
		attr := NewAttributeInfo(reader, code.cp) // override point1
		num := attr.ReadInfo(reader)              // override point2
		code.attrs[i] = attr
		if reader.length()+num+ATTRBASELEN != start { // verification
			panic(fmt.Errorf("wrong number for parsing %#v", attr))
		}
	}
	return num
}

func (code *AttrCode) AttrString() string {
	buf := &bytes.Buffer{}
	buf.WriteString("code:\n")
	buf.WriteString(fmt.Sprintf("maxStack: %d", code.maxStack))
	buf.WriteByte(',')
	buf.WriteString(fmt.Sprintf("maxLocals: %d\n", code.maxLocals))
	// code
	cr := instructions.NewCodeReader(code.code)
	cr.ReadCode()
	for i, c := range cr.Code() {
		buf.WriteString(fmt.Sprintf("#%d %s\n", i, instructions.InstStr(c)))
	}
	return buf.String()
}

type AttrExceptionTableEntry struct {
	startPC   uint16 // startPC and endPC specified the range of the exception handler code
	endPC     uint16 // index into code array [startPC, endPC) historical bug for JVM designer
	handlerPC uint16 // index into code array, indicating where to start the handler
	catchType uint16 // index into constant pool to class info
}

func (entry *AttrExceptionTableEntry) ReadInfo(reader *ClassReader) {
	entry.startPC = reader.ReadUint16()
	entry.endPC = reader.ReadUint16()
	entry.handlerPC = reader.ReadUint16()
	entry.catchType = reader.ReadUint16()
}
