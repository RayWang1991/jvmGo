package marea

import cf "jvmGo/jvm/classfile"

type MethodType string

func NewMethodType(cp cf.ConstantPool, info *cf.MethodTypeInfo) MethodType {
	return MethodType(info.Desc(cp))
}
