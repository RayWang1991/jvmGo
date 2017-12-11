package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_Reflection, "getCallerClass", "()Ljava/lang/Class;", getCallerClass)
}

//public static native ()Ljava/lang/Class;
func getCallerClass(f *rtdt.Frame) {

	nf := f.GetNext().GetNext()
	ref := nf.Method().Class().GetClassObject()
	//debug
	//fmt.Printf("1230%s %s\n", f.Method().Name(), f.Method().Class().GetClassObject().GetClzClass().ClassName())
	//fmt.Printf("1230%s %s\n", f.GetNext().Method().Name(), f.GetNext().Method().Class().GetClassObject().GetClzClass().ClassName())
	//fmt.Printf("1232%s %s\n", nf.Method().Name(), ref.GetClzClass().ClassName())
	f.OperandStack.PushRef(ref)
}
