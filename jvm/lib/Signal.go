package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_Signal, "findSignal", "(Ljava/lang/String;)I", findSignal)
	register(utils.CLASSNAME_Signal, "handle0", "(IJ)J", handle0)
}

// private static native int findSignal(String var0);
// (Ljava/lang/String;)I

func findSignal(f *rtdt.Frame) {
	jstr := f.LocalVar.GetRef(0)
	gostr := marea.GetGoString(jstr)
	utils.DIstrPrintf("FINDSIGNAL %s\n", gostr)
	f.OperandStack.PushInt(0)
	//todo
}

// private static native long handle0(int i, long l);
// (IJ)J
func handle0(f *rtdt.Frame) {
	f.OperandStack.PushLong(0)
	//todo
}
