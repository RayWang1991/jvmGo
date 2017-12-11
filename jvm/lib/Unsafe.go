package lib

import (
	"jvmGo/jvm/rtdt"
	"unsafe"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_Unsafe, "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	register(utils.CLASSNAME_Unsafe, "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	register(utils.CLASSNAME_Unsafe, "addressSize", "()I", addressSize)
}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtdt.Frame) {
	stack := frame.OperandStack
	stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtdt.Frame) {
	stack := frame.OperandStack
	stack.PushInt(1) // todo
}

// public native int addressSize();
// ()I
func addressSize(frame *rtdt.Frame) {
	// vars := frame.LocalVars()
	// vars.GetRef(0) // this

	stack := frame.OperandStack
	res := int32(unsafe.Sizeof(int(1)))
	//fmt.Println(res)
	stack.PushInt(res)
}
