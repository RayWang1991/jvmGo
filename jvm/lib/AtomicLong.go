package lib

import (
	"jvmGo/jvm/utils"
	"jvmGo/jvm/rtdt"
)

func init() {
	register(utils.CLASSNAME_AtomicLong, "VMSupportsCS8", "()Z", VMSupportsCS8)
}

func VMSupportsCS8(f *rtdt.Frame) {
	f.OperandStack.PushInt(0) // return false
}
