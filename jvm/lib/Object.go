package lib

import (
	"jvmGo/jvm/rtdt"
	"unsafe"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_Object, "hashCode", "()I", hashCode)
	register(utils.CLASSNAME_Object, "registerNatives", "()V", registerNatives)
}

func registerNatives(frame *rtdt.Frame) {
	// do nothing
}

// public native int hashCode()
func hashCode(f *rtdt.Frame) {
	ref := f.LocalVar.GetRef(0)
	hash := int32(uintptr(unsafe.Pointer(ref)))
	f.OperandStack.PushInt(hash)
}
