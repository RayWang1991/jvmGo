package classfile

import (
	"fmt"
	"strings"
	"bytes"
	"strconv"
)

func readerPos(cr *ClassReader) string {
	return fmt.Sprintf("len:%d", cr.length())
}

// TODO using template?
func (cf *ClassFile) PrintDebugMessage() {
	fmt.Printf("magic: %X\n", cf.magic) // magic
	fmt.Printf("version: %d.%d\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Print(cf.constantPool.String())
	fmt.Printf("class: %s\n", cf.ClassName())
	fmt.Printf("super class: %s\n", cf.SuperClassName())
	fmt.Printf("interfaces(%d items): %s \n", len(cf.interfaces), strings.Join(cf.InterfaceNames(), ","))
	// TODO fields, methods, attrinfos
}

// constant pool entry format: #64 = Methodref       #6.#44      // java/lang/Object."<init>":()V
var cpInfoFmt = "%6s = %-20s%-30s%s\n"

func (cp ConstantPool) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString("Constant pool:\n")
	for i := range cp {
		if i == 0 {
			continue
		}
		arg0, arg1, arg2 := debugString(cp, cp[i])
		buf.WriteString(
			fmt.Sprintf(cpInfoFmt, debugIndex(uint(i)), arg0, arg1, arg2))
	}
	return buf.String()
}

func debugString(cp ConstantPool, info ConstInfo) (string, string, string) {
	switch info := info.(type) {
	case *ClassInfo:
		return "Class", debugIndex(uint(info.nameIndex)), "// " + cp.getUtf8(info.nameIndex)
	case *Utf8Info:
		return "Utf8", info.val, ""
	case *IntegerInfo:
		return "Integer", strconv.Itoa(int(info.val)), ""
	case *FloatInfo:
		return "Float", strconv.FormatFloat(float64(info.val), 'f', -1, 32) + "f", ""
	case *LongInfo:
		return "Long", strconv.Itoa(int(info.val)) + "l", ""
	case *DoubleInfo:
		return "Double", strconv.FormatFloat(float64(info.val), 'f', -1, 64) + "d", ""
	case *StringInfo:
		return "String", debugIndex(uint(info.index)), "// " + cp.getUtf8(info.index)
	case *FieldRefInfo:
		return "Fieldref",
			debugIndex(uint(info.classIndex)) + "." + debugIndex(uint(info.nameTypeIndex)),
			"// " + info.NameTypeInfo(cp).String(cp)
	case *MethodRefInfo:
		return "Methodref",
			debugIndex(uint(info.classIndex)) + "." + debugIndex(uint(info.nameTypeIndex)),
			"// " + info.NameTypeInfo(cp).String(cp)
	case *InterfaceMethodRefInfo:
		return "InterfaceMethodref",
			debugIndex(uint(info.classIndex)) + "." + debugIndex(uint(info.nameTypeIndex)),
			"// " + info.NameTypeInfo(cp).String(cp)
	case *NameTypeInfo:
		return "NameAndType", debugIndex(uint(info.nameIndex)) + ":" +
			debugIndex(uint(info.typeIndex)),
			"// " + info.String(cp)
	case *MethodHandleInfo: // I don't know how to print it
		return "MethodHandle", debugIndex(uint(info.refKind)) + ":" + debugIndex(uint(info.refIndex)), ""
	case *MethodTypeInfo:
		return "MethodType", debugIndex(uint(info.descIndex)), "// " + cp.getUtf8(info.descIndex)
	case *InvokeDynamic_Info: // I don't know how to print it
		return "InvokeDynamic",
			debugIndex(uint(info.bootstrapMethodAttrIndex)) + ":" + debugIndex(uint(info.nameTypeIndex)), ""
	default:
		return "Unknown", "", ""
	}
}

func debugIndex(i uint) string {
	return "#" + strconv.Itoa(int(i))
}

// TODO
func (memb *MemberInfo) String() {
}
