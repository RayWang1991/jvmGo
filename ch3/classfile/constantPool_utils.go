package classfile

import (
	"fmt"
	"unicode/utf16"
)


// get string through utf8 index
func (info *StringInfo) StringWith(cp ConstantPool) string {
	utf8I := cp[info.index].(*Utf8Info)
	return utf8I.val
}

// get class name through utf8 index
func (info *ClassInfo) ClassName(cp ConstantPool) string {
	utf8I := cp[info.nameIndex].(*Utf8Info)
	return utf8I.val
}

// debug string for name and type
func (info *NameTypeInfo) String(cp ConstantPool) string {
	return info.Name(cp) + ":" + info.Type(cp)
}

// get field or method name through utf8 index
func (info *NameTypeInfo) Name(cp ConstantPool) string {
	utf8I := cp[info.nameIndex].(*Utf8Info)
	return utf8I.val
}

// get field or method type descriptor through utf8 index
func (info *NameTypeInfo) Type(cp ConstantPool) string {
	utf8I := cp[info.typeIndex].(*Utf8Info)
	return utf8I.val
}

// get class info for field reference
func (info *FieldRefInfo) ClassInfo(cp ConstantPool) *ClassInfo {
	return cp[info.classIndex].(*ClassInfo)
}

// get name and type info for field reference
func (info *FieldRefInfo) NameTypeInfo(cp ConstantPool) *NameTypeInfo {
	return cp[info.nameTypeIndex].(*NameTypeInfo)
}

// get class info for Method reference
func (info *MethodRefInfo) ClassInfo(cp ConstantPool) *ClassInfo {
	return cp[info.classIndex].(*ClassInfo)
}

// get name and type info for Method reference
func (info *MethodRefInfo) NameTypeInfo(cp ConstantPool) *NameTypeInfo {
	return cp[info.nameTypeIndex].(*NameTypeInfo)
}

// get class info for Interface Method reference
func (info *InterfaceMethodRefInfo) ClassInfo(cp ConstantPool) *ClassInfo {
	return cp[info.classIndex].(*ClassInfo)
}

// get name and type info for Interface Method reference
func (info *InterfaceMethodRefInfo) NameTypeInfo(cp ConstantPool) *NameTypeInfo {
	return cp[info.nameTypeIndex].(*NameTypeInfo)
}
