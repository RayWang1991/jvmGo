package marea

import cf "jvmGo/ch6/classfile"

type MethodType string

func NewMethodType(cp cf.ConstantPool, info *cf.MethodTypeInfo) MethodType {
	return MethodType(info.Desc(cp))
}
