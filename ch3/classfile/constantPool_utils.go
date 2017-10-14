package classfile

import (
	"fmt"
	"unicode/utf16"
)

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}

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
