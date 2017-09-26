// constant pool contains two kinds of type datas
// 1. Literals
// 2. Symbolic references
package classfile

import (
	"fmt"
)

type ConstantPool []ConstInfo

func NewConstantPool(reader ClassReader) ConstantPool {
	n := reader.ReadUint16()
	cp := make([]ConstInfo, n) // nil initiated
	var i uint16 = 1           // start from 1
	var d uint16
	for i < n {
		tag := reader.ReadUint16()
		cp[i], d = NewConstInfo(tag, reader)
		i += d
	}
	return cp
}

// init const info through tag and reader([]byte)
func NewConstInfo(tag uint16, reader ClassReader) (info ConstInfo, l uint16) {
	l = 1
	switch tag {
	case CONST_Utf8_Info:
		info = &Utf8Info{}
	case CONST_Integer_Info:
		info = &IntegerInfo{}
	case CONST_Float_Info:
		info = &FloatInfo{}
	case CONST_Long_Info:
		l = 2
		info = &LongInfo{}
	case CONST_Double_Info:
		l = 2
		info = &DoubleInfo{}
	case CONST_Class_Info:
		info = &ClassInfo{}
	case CONST_String_Info:
		info = &StringInfo{}
	case CONST_FieldRef_Info:
		info = &FieldRefInfo{}
	case CONST_MethodRef_Info:
		info = &MethodRefInfo{}
	case CONST_InterfaceMethodRef_Info:
		info = &InterfaceMethodRefInfo{}
	case CONST_NameAndType_Info:
		info = &NameTypeInfo{}
	case CONST_MethodHandle_Info:
		info = &MethodHandleInfo{}
	case CONST_MethodType_Info:
		info = &MethodTypeInfo{}
	case CONST_InvokeDynamic_Info:
		info = &InvokeDynamic_Info{}
	default:
		panic(fmt.Errorf("java.lang.ClassFormatError: constant pool tag: %d !", tag))
	}
	info.ReadInfo(reader)
	return
}

type ConstInfo interface {
	ReadInfo(reader ClassReader)
}

const (
	CONST_Utf8_Info               = 1
	CONST_Integer_Info            = 3
	CONST_Float_Info              = 4
	CONST_Long_Info               = 5
	CONST_Double_Info             = 6
	CONST_Class_Info              = 7
	CONST_String_Info             = 8
	CONST_FieldRef_Info           = 9
	CONST_MethodRef_Info          = 10
	CONST_InterfaceMethodRef_Info = 11
	CONST_NameAndType_Info        = 12
	CONST_MethodHandle_Info       = 15
	CONST_MethodType_Info         = 16
	CONST_InvokeDynamic_Info      = 18
)

func (cp ConstantPool) getUtf8(index uint16) string {
	return cp[index].(*Utf8Info).val
}
