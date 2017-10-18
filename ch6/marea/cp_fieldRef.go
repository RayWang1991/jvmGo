package marea

import (
	cf "jvmGo/ch6/classfile"
	"jvmGo/ch6/utils"
)

func NewFieldRef(cp cf.ConstantPool, i *cf.FieldRefInfo, cls *Class) *FieldRef {
	return &FieldRef{MemberRef: *NewRef(cp, &(i.RefInfo), cls)}
}

type FieldRef struct {
	MemberRef
	f *Field
}

func (r *FieldRef) GetField() *Field {
	if r.f != nil {
		return r.f
	}
	c := r.Ref()
	f := c.LookUpField(r.name)
	if f == nil {
		panic(utils.NoSuchFieldError)
	}
	if !isAccessableField(r.from, f) {
		panic(utils.IllegalAccessError)
	}
	r.f = f
	return f
}
