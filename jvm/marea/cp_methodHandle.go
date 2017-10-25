package marea

import cf "jvmGo/jvm/classfile"

type REF_Kind byte

const (
	REF_getField REF_Kind = iota
	REF_getStatic
	REF_putField
	REF_putStatic
	REF_invokeVirtual
	REF_invokeStatic
	REF_newInvokeSpecial
	REF_invokeInterface
)

type MethodHandle struct {
	kind REF_Kind
	ref  *MethodRef // TODO, check if there is one, avoiding dumplication
}

func NewMethodHandle(cp cf.ConstantPool, info *cf.MethodHandleInfo, cls *Class) *MethodHandle {
	mh := &MethodHandle{}
	mh.kind = REF_Kind(info.RefKind())
	refCp := cp[info.RefIndex()].(*cf.MethodRefInfo)
	ref := NewMethodRef(cp, refCp, cls)
	mh.ref = ref
	return mh
}
