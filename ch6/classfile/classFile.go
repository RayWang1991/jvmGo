package classfile

import (
	"errors"
	"jvmGo/ch6/utils"
	"jvmGo/ch6/cmn"
)

type ClassFile struct {
	magic          uint32 // magic number, always be 0xCAFEBABE
	minorVersion   uint16
	majorVersion   uint16
	constantPool   ConstantPool
	accessFlag     uint16
	thisClass      uint16
	superClass     uint16 // Object doesn't have one
	interfaces     []uint16
	staticFields   []*FieldInfo
	sSlotN         uint
	instanceFields []*FieldInfo
	methods        []*MethodInfo
	attributes     []AttrInfo
}

// getter for minorVersion
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

// getter for majorVersion
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

// getter for accessFlags
func (cf *ClassFile) AccessFlag() uint16 {
	return cf.accessFlag
}

// getter for ConstantPool
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

// getter for class name
func (cf *ClassFile) ClassName() string {
	cp := cf.constantPool
	return cp[cf.thisClass].(*ClassInfo).ClassName(cp)
}

// getter for super class name
func (cf *ClassFile) SuperClassName() string {
	if cf.ClassName() == utils.CLASSNAME_Object {
		return ""
	}
	cp := cf.constantPool
	return cp[cf.superClass].(*ClassInfo).ClassName(cp)
}

// getter for interface names
// cache if needed
func (cf *ClassFile) InterfaceNames() []string {
	cp := cf.constantPool
	res := make([]string, 0, len(cf.interfaces))
	for _, i := range cf.interfaces {
		res = append(res, cp[i].(*ClassInfo).ClassName(cp))
	}
	return res
}

// getter for Fields
func (cf *ClassFile) StatFields() []*FieldInfo {
	return cf.staticFields
}

func (cf *ClassFile) InstFields() []*FieldInfo {
	return cf.instanceFields
}

func (cf *ClassFile) StatSlotN() uint {
	return cf.sSlotN
}

// getter for Methods
func (cf *ClassFile) MethodInfo() []*MethodInfo {
	return cf.methods
}

const magicNumber = 0xCAFEBABE

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) error {
	magic := reader.ReadUint32()
	if magic != magicNumber {
		return errors.New(utils.ClassMagicError)
	}
	cf.magic = magic
	return nil
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) error {
	cf.minorVersion = reader.ReadUint16()
	cf.majorVersion = reader.ReadUint16()
	if cf.majorVersion >= 45 && cf.majorVersion <= 52 {
		return nil
	} else {
		return errors.New(utils.ClassVersionError)
	}
}

// read access flag and check it
func (cf *ClassFile) readAccessFlag(reader *ClassReader) error {
	cf.accessFlag = reader.ReadUint16()
	return nil
}

// read class index in the constant pool
func (cf *ClassFile) readClassIndex(reader *ClassReader) {
	cf.thisClass = reader.ReadUint16()
}

// read super class index in the constant pool
func (cf *ClassFile) readSuperClassIndex(reader *ClassReader) {
	cf.superClass = reader.ReadUint16()
}

// read interfaces
func (cf *ClassFile) readInterfaces(reader *ClassReader) {
	n := reader.ReadUint16()
	cf.interfaces = make([]uint16, n)
	for i := range cf.interfaces {
		cf.interfaces[i] = reader.ReadUint16()
	}
}

// read fields and methods
func (cf *ClassFile) readFieldsAndMethods(reader *ClassReader) {
	nf := reader.ReadUint16()
	statNum := uint16(0)
	slts := uint(0)
	fields := make([]*FieldInfo, 0, nf)
	for i := uint16(0); i < nf; i++ {
		field := &FieldInfo{
			MemberInfo{cp: cf.constantPool},
			0,
		}
		field.ReadInfo(reader)
		var slotNum uint8
		if cmn.IsStatic(field.accessFlags) {
			statNum++
			slotNum = cmn.SlotNum(field.Description())
			slts += uint(slotNum)
		} else {
			slotNum = cmn.SlotNum(field.Description())
		}
		field.slotNum = slotNum
		fields = append(fields, field)
	}
	cf.sSlotN = slts
	cf.instanceFields = make([]*FieldInfo, 0, nf-statNum)
	cf.staticFields = make([]*FieldInfo, 0, statNum)
	for _, f := range fields {
		if cmn.IsStatic(f.accessFlags) {
			cf.staticFields = append(cf.staticFields, f)
		} else {
			cf.instanceFields = append(cf.instanceFields, f)
		}
	}

	nm := reader.ReadUint16()
	methods := make([]*MethodInfo, 0, nm)
	for i := uint16(0); i < nm; i++ {
		method := &MethodInfo{MemberInfo{cp: cf.constantPool}}
		method.ReadInfo(reader)
		methods = append(methods, method)
	}
	cf.methods = methods
}

// read sequentially
func NewClassFile(reader *ClassReader) (cf *ClassFile, err error) {
	cf = &ClassFile{}
	err = cf.readAndCheckMagic(reader)
	if err != nil {
		return nil, err
	}
	err = cf.readAndCheckVersion(reader)
	if err != nil {
		return nil, err
	}
	//defer func() {
	//	r := recover()
	//	if e, ok := r.(error); ok {
	//		err = e
	//	} else {
	//		err = fmt.Errorf("parsing FromClass File: %v", r)
	//	}
	//}()
	cf.constantPool = NewConstantPool(reader)
	err = cf.readAccessFlag(reader)
	if err != nil {
		return nil, err
	}
	cf.readClassIndex(reader)
	cf.readSuperClassIndex(reader)
	cf.readInterfaces(reader)
	cf.readFieldsAndMethods(reader)
	return
}
