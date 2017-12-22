package lib

import (
	"jvmGo/jvm/rtdt"
	"runtime"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_Runtime, "availableProcessors", "()I", availableProcessors)
}

// public native int availableProcessors();
// ()I
func availableProcessors(f *rtdt.Frame) {
	//todo
	num := runtime.NumCPU()
	f.OperandStack.PushInt(int32(num))
}
