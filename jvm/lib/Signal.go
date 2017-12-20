package lib

import (
	"jvmGo/jvm/rtdt"
	"fmt"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_Signal, "findSignal", "(Ljava/lang/String;)I", findSignal)
}

// private static native int findSignal(String var0);
// (Ljava/lang/String;)I

func findSignal(f *rtdt.Frame) {
	jstr := f.LocalVar.GetRef(0)
	gostr := marea.GetGoString(jstr)
	fmt.Printf("FINDSIGNAL %s\n", gostr)
	f.OperandStack.PushInt(0)
}

