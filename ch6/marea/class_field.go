package marea

import (
	"jvmGo/ch6/cmn"
	"jvmGo/ch6/classfile"
)

type Field struct {
	ClassMember
	vId uint16 // index to vars
}

func NewField(info *classfile.FieldInfo) *Field {
	f := &Field{}
	f.name = info.Name()
	f.flags = info.AccFlags()
	f.desc = info.Description()
	return f
}

// setter
func (f *Field) SetVId(id uint16) {
	f.vId = id
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
