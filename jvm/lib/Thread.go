package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/marea"
)

func init() {
	register(utils.CLASSNAME_Thread, "currentThread", "()Ljava/lang/Thread;", currentThread)
	register(utils.CLASSNAME_Thread, "setPriority0", "(I)V", setPriority0)
	register(utils.CLASSNAME_Thread, "isAlive", "()Z", isAlive)
	register(utils.CLASSNAME_Thread, "start0", "()V", start0)
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

// Todo, thread system
// private native void setPriority0(int newPriority);
// (I)V
func setPriority0(frame *rtdt.Frame) {
	// vars := frame.LocalVars()
	// this := vars.GetThis()
	// newPriority := vars.GetInt(1))
}

// TODO
// public final native boolean isAlive();
// ()Z
func isAlive(frame *rtdt.Frame) {
	//vars := frame.LocalVars()
	//this := vars.GetThis()

	stack := frame.OperandStack
	stack.PushInt(0)
}

// TODO
// public native void start0()
// ()V
func start0(frame *rtdt.Frame) {
}
