package marea

import cf "jvmGo/ch6/classfile"

type ClassInfo string

func NewClassInfo(cp cf.ConstantPool, info *cf.ClassInfo) ClassInfo {
	return ClassInfo(info.ClassName(cp))
}
