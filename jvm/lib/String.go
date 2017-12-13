package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_String, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern()
// ()Ljava/lang/String;
func intern(f *rtdt.Frame) {
	this := f.LocalVar.GetRef(0)
	internStr := marea.GetInternString(this)
	f.OperandStack.PushRef(internStr)
}
