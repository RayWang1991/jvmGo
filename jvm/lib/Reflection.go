package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"fmt"
)

func init() {
	register(utils.CLASSNAME_Reflection, "getCallerClass", "()Ljava/lang/Class;", getCallerClass)
	register(utils.CLASSNAME_Reflection, "getClassAccessFlags", "(Ljava/lang/Class;)I", getClassAccessFlags)
}

//public static native ()Ljava/lang/Class;
func getCallerClass(f *rtdt.Frame) {

	nf := f.GetNext().GetNext()
	ref := nf.Method().Class().GetClassObject()
	if utils.DebugFlag && utils.IstrDebugFlag {
		for _, f := range ref.GetClzClass().FieldMap() {
			fmt.Printf("FIELD %s %s\n", f.Name(), f.Desc())
		}
	}
	f.OperandStack.PushRef(ref)
}

//public static native int getClassAccessFlags(Class<?> var0);
//(Ljava/lang/Class;)I
func getClassAccessFlags(f *rtdt.Frame) {
	cls := f.LocalVar.GetRef(0)
	flags := int32(cls.GetClzClass().GetFlags())
	f.OperandStack.PushInt(flags)
}
