package lib

import (
	"jvmGo/jvm/utils"
	"jvmGo/jvm/rtdt"
	"fmt"
	"jvmGo/jvm/cmn"
	"jvmGo/jvm/marea"
)

func init() {
	register(utils.CLASSNAME_NativeConstructorAccessorImpl, "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", newInstance0)
}

// private static native Object newInstance0(Constructor<?> var0, Object[] var1) throws InstantiationException, IllegalArgumentException, InvocationTargetException;
// (Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;
func newInstance0(frame *rtdt.Frame) {
	constructor := frame.LocalVar.GetRef(0) // constructor
	consClz := constructor.Class()
	//debug
	fmt.Printf("const %s\n", consClz.ClassName())
	clzSlotId := consClz.InstField("clazz").VarIdx()
	descId := consClz.InstField("parameterAnnotations").VarIdx()
	vars := frame.LocalVar.GetRef(1) // Object[]
	var varArray []*marea.Object
	if vars != nil {
		varArray = vars.ArrGetRefs()
	}

	clzObj := constructor.GetRef(clzSlotId)
	class := clzObj.GetClzClass()
	loader := frame.Method().Class().DefineLoader()
	loader.Initiate(class)

	// find the constructor method from name and desc
	descBA := constructor.GetRef(uint(descId))
	is := descBA.ArrGetBytes()
	bs := make([]byte, len(is))
	for j, b := range is {
		bs[j] = byte(b)
	}
	desc := string(bs)
	m := class.GetMethodDirect("<init>", desc)
	fmt.Printf("CONSTM %s %s\n", m.Desc(),class.ClassName())
	//panic("TODO")

	if len(varArray) != len(m.ArgDs()) {
		panic(utils.IllegalArgumentException)
	}
	// check the arguments, todo basic args
	for _, arg := range m.ArgDs() {
		if cmn.IsPrimitiveType(arg) {
			panic("todo primitive type")
		}
	}

	// new instance
	instance := marea.NewObject(class)
	thread := frame.Thread()

	frame.OperandStack.PushRef(instance)
	op := rtdt.NewOperandStack(uint(len(varArray) + 1))
	op.PushRef(instance)
	for _, arg := range varArray { //todo
		op.PushRef(arg)
	}
	df := dummyFrame(op, thread)
	thread.PushFrame(df)
	callMethod(m, df)
}
