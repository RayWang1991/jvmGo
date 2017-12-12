package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/marea"
)

func init() {
	register(utils.CLASSNAME_Thread, "currentThread", "()Ljava/lang/Thread;", currentThread)
}

// public static native currentThread ()Ljava/lang/Thread;
func currentThread(f *rtdt.Frame) {
	loader := f.Method().Class().DefineLoader()
	threadC := loader.Load(utils.CLASSNAME_Thread)
	currentT := marea.NewObject(threadC)
	groupC := loader.Load(utils.CLASSNAME_ThreadGroup)
	groupO := marea.NewObject(groupC)
	fieldG := threadC.InstField("group")
	fieldP := threadC.InstField("priority")
	currentT.SetRef(groupO, fieldG.VarIdx())
	currentT.SetInt(1, fieldP.VarIdx())
	f.OperandStack.PushRef(currentT)
}
