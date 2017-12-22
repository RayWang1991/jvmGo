package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_ClassLoader, "findLoadedClass0", "(Ljava/lang/String;)Ljava/lang/Class;", findLoadedClass0)
}

// private native final Class<?> findLoadedClass0(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findLoadedClass0(f *rtdt.Frame) {
	jstr := f.LocalVar.GetRef(1)
	gostr := marea.GetGoString(jstr)
	loader := f.Method().Class().DefineLoader() // todo
	clz := loader.Load(gostr)
	f.OperandStack.PushRef(clz.GetClassObject())
}
