package lib

import (
	"jvmGo/jvm/rtdt"
	"unsafe"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_Object, "hashCode", "()I", hashCode)
	register(utils.CLASSNAME_Object, "registerNatives", "()V", registerNatives)
	register(utils.CLASSNAME_Object, "clone", "()Ljava/lang/Object;", clone)
	register(utils.CLASSNAME_Object, "getClass", "()Ljava/lang/Class;", getClass)
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

// protected native Ljava/lang/Object; clone()
// ()Ljava/lang/Object;
func clone(f *rtdt.Frame) {
	this := f.LocalVar.GetRef(0)
	if this == nil {
		f.OperandStack.PushRef(nil)
		return
	}
	clz := this.Class()
	//check is cloneable
	var found = false
	for _, inf := range clz.InterfaceNames() {
		if inf == utils.CLASSNAME_Cloneable {
			found = true
			break
		}
	}
	if !found {
		panic(utils.CloneNotSupportedException)
	}

	copied := this.Copy()
	f.OperandStack.PushRef(copied)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(f *rtdt.Frame) {
	this := f.LocalVar.GetRef(0)
	if this == nil {
		panic(utils.NullPointerException)
	}
	class := this.Class()
	clzObj := class.GetClassObject()
	f.OperandStack.PushRef(clzObj)
}
