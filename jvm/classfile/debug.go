package classfile

import (
	"bytes"
	"fmt"
	"jvmGo/jvm/cmn"
	"strconv"
	"strings"
)

func readerPos(cr *ClassReader) string {
	return fmt.Sprintf("len:%d", cr.length())
}

// Print Debug Message for FromClass File
// TODO using template?
// TODO debug string for attr info
func (cf *ClassFile) PrintDebugMessage() {
	fmt.Printf("magic: %X\n", cf.magic) // magic
	fmt.Printf("version: %d.%d\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("flags: %s\n", cmn.FlagNumToString(cf.accessFlag, cmn.ACC_TYPE_CLASS))
	fmt.Print(cf.constantPool.String())
	fmt.Printf("class: %s\n", cf.ClassName())
	fmt.Printf("super class: %s\n", cf.SuperClassName())
	fmt.Printf("interfaces(%d items): %s \n", len(cf.interfaces), strings.Join(cf.InterfaceNames(), ","))
	// Fields and Methods
	fmt.Printf("Fields static (%d items):\n", len(cf.staticFields))
	for i, f := range cf.staticFields {
		fmt.Print(f.String(fmt.Sprintf("#%d\n", i), cmn.ACC_TYPE_FIELD))
	}
	fmt.Printf("Fields instance (%d items):\n", len(cf.instanceFields))
	for i, f := range cf.instanceFields {
		fmt.Print(f.String(fmt.Sprintf("#%d\n", i), cmn.ACC_TYPE_FIELD))
	}
	fmt.Printf("Methods(%d items):\n", len(cf.methods))
	for i, m := range cf.methods {
		fmt.Print(m.String(fmt.Sprintf("#%d\n", i), cmn.ACC_TYPE_METHOD))
		if cmn.IsNative(m.accessFlags) {
			fmt.Println("Native Method")
		} else {
			codeAttr := m.GetCodeAttr()
			fmt.Print(codeAttr.AttrString())
		}
	}
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

// debug string for constant info in constant pool
func debugString(cp ConstantPool, info ConstInfo) (string, string, string) {
	switch info := info.(type) {
	case *ClassInfo:
		return "FromClass", debugIndex(uint(info.nameIndex)), "// " + cp.GetUTF8(info.nameIndex)
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
		return "String", debugIndex(uint(info.index)), "// " + cp.GetUTF8(info.index)
	case *FieldRefInfo:
		return "Fieldref",
			debugIndex(uint(info.classIndex)) + "." + debugIndex(uint(info.nameTypeIndex)),
			"// " + info.ClassInfo(cp).ClassName(cp) + "." + info.NameTypeInfo(cp).String(cp)
	case *MethodRefInfo:
		return "Methodref",
			debugIndex(uint(info.classIndex)) + "." + debugIndex(uint(info.nameTypeIndex)),
			"// " + info.ClassInfo(cp).ClassName(cp) + "." + info.NameTypeInfo(cp).String(cp)
	case *InterfaceMethodRefInfo:
		return "InterfaceMethodref",
			debugIndex(uint(info.classIndex)) + "." + debugIndex(uint(info.nameTypeIndex)),
			"// " + info.ClassInfo(cp).ClassName(cp) + "." + info.NameTypeInfo(cp).String(cp)
	case *NameTypeInfo:
		return "NameAndType", debugIndex(uint(info.nameIndex)) + ":" +
				debugIndex(uint(info.typeIndex)),
			"// " + info.String(cp)
	case *MethodHandleInfo: // I don't know how to print it
		return "MethodHandle", debugIndex(uint(info.refKind)) + ":" + debugIndex(uint(info.refIndex)), ""
	case *MethodTypeInfo:
		return "MethodType", debugIndex(uint(info.descIndex)), "// " + cp.GetUTF8(info.descIndex)
	case *InvokeDynamicInfo: // I don't know how to print it
		return "InvokeDynamic",
			debugIndex(uint(info.bootstrapMethodAttrIndex)) + ":" + debugIndex(uint(info.nameTypeIndex)), ""
	default:
		return "Unknown", "", ""
	}
}

func debugIndex(i uint) string {
	return "#" + strconv.Itoa(int(i))
}

func (m *MemberInfo) String(title string, acc_type cmn.ACC_TYPE) string {
	buf := &bytes.Buffer{}
	buf.WriteString(title)
	buf.WriteString("name: ")
	buf.WriteString(m.cp.GetUTF8(m.nameIndex))
	buf.WriteByte('\n')
	buf.WriteString("flags: ")
	buf.WriteString(cmn.FlagNumToString(m.accessFlags, acc_type))
	buf.WriteByte('\n')
	buf.WriteString("descriptor: ")
	buf.WriteString(m.cp.GetUTF8(m.descIndex))
	buf.WriteByte('\n')
	return buf.String()
}

func (cp ConstantPool) getConstDebugString(index uint16) string {
	switch x := cp[index].(type) {
	case *IntegerInfo:
		return fmt.Sprintf("int %d", x.val)
	case *LongInfo:
		return fmt.Sprintf("long %dl", x.val)
	case *FloatInfo:
		return fmt.Sprintf("float %ff", x.val)
	case *DoubleInfo:
		return fmt.Sprintf("double %fd", x.val)
	case *Utf8Info:
		return fmt.Sprintf("String %s", x.val)
	default:
		return "Unknown ConstantValue"
	}
}

// debug for code
type CodeInst []byte

func (code CodeInst) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString("code:\n")
	// code
	cr := cmn.NewCodeReader(code)
	cr.ReadCode()
	for i, c := range cr.Code() {
		buf.WriteString(fmt.Sprintf("#%d %s\n", i, cmn.InstStr(c)))
	}
	return buf.String()
}
