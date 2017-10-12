package marea

import "jvmGo/ch6/classfile"

type InvokeDynamic struct {
	bstmI uint16 // index to bootstrap_methods
	name  string
	tp    string
}

func NewInvokeDynamic(cp classfile.ConstantPool, info *classfile.InvokeDynamicInfo) *InvokeDynamic {
	invoke := &InvokeDynamic{}
	invoke.name, invoke.tp = info.NameType(cp)
	invoke.bstmI = info.BstmI()
}
