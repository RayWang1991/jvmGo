package marea

import (
	"jvmGo/jvm/classfile"
	"jvmGo/jvm/cmn"
)

type Field struct {
	ClassMember
	vIdx uint  // index to vars
	sn   uint8 // slotNum
}

func NewField(from *Class, info *classfile.FieldInfo) *Field {
	f := &Field{}
	f.class = from
	f.name = info.Name()
	f.flags = info.AccFlags()
	f.desc = info.Description()
	f.sn = info.SlotNum()
	return f
}

// setter
func (f *Field) SetVarIdx(i uint) {
	f.vIdx = i
}

// getter
func (f *Field) VarIdx() uint {
	return f.vIdx
}

func (f *Field) SlotNum() uint8 {
	return f.sn
}

// access methods
func (f *Field) IsPublic() bool {
	return cmn.IsPublic(f.flags)
}

func (f *Field) IsPrivate() bool {
	return cmn.IsPrivate(f.flags)
}

func (f *Field) IsProtected() bool {
	return cmn.IsProtected(f.flags)
}

func (f *Field) IsStatic() bool {
	return cmn.IsStatic(f.flags)
}

func (f *Field) IsFinal() bool {
	return cmn.IsFinal(f.flags)
}

func (f *Field) IsVolatile() bool {
	return cmn.IsVolatile(f.flags)
}

func (f *Field) IsEnum() bool {
	return cmn.IsEnum(f.flags)
}
