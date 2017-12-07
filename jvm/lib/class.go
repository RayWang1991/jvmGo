package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/marea"
)

func init() {
	register(utils.CLASSNAME_Class, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	register(utils.CLASSNAME_Class, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *rtdt.Frame) {
	// todo
	frame.OperandStack.PushInt(0)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtdt.Frame) {
	nameObj := frame.LocalVar.GetRef(0)
	name := marea.GetGoString(nameObj)

	loader := frame.Method().Class().DefineLoader()
	class := loader.Load(name).GetClassObject()

	frame.OperandStack.PushRef(class)
}
