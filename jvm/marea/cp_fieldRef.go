package marea

import (
	"fmt"
	cf "jvmGo/jvm/classfile"
	"jvmGo/jvm/utils"
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
	fmt.Println(r.name)
	fmt.Println(c.FieldMap())
	if f == nil {
		panic(utils.NoSuchFieldError)
	}
	if !isAccessableField(r.from, f) {
		panic(utils.IllegalAccessError)
	}
	r.f = f
	return f
}
