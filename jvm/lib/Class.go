package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/cmn"
	"fmt"
)

func init() {
	register(utils.CLASSNAME_Class, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	register(utils.CLASSNAME_Class, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	register(utils.CLASSNAME_Class, "getName0", "()Ljava/lang/String;", getName0)
	register(utils.CLASSNAME_Class, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	register(utils.CLASSNAME_Class, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
	register(utils.CLASSNAME_Class, "isPrimitive", "()Z", isPrimitive)
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

// private native String getName0()
// ()Ljava/lang/String;
func getName0(f *rtdt.Frame) {
	ref := f.LocalVar.GetRef(0)
	name := ref.GetClzClass().DotedName()
	loader := f.Method().Class().DefineLoader()
	str := marea.GetJavaString(name, loader)
	fmt.Printf("name %s\n", name)
	f.OperandStack.PushRef(str)
}

// private static native Class<?> forName0(String name, boolean initialize, ClassLoader loader, Class<?> caller)
// (Ljava/lang/String;ZLjava/lang/ClassLoader/;Ljava/lang/Class;)Ljava/lang/Class;
func forName0(f *rtdt.Frame) {
	jName := f.LocalVar.GetRef(0)
	goName := cmn.ToSlash(marea.GetGoString(jName))
	//debug
	fmt.Printf("stack %s\n", f.OperandStack)
	fmt.Printf("local var %s\n", f.LocalVar)
	fmt.Printf("name %s class %s\n", goName, jName.Class().ClassName())

	// TODO, initialize and loader is not used
	//initialize := f.LocalVar.GetInt(1)

	loader := f.Method().Class().DefineLoader()
	class := loader.Load(goName)
	loader.Initiate(class)
	clObj := class.GetClassObject()
	f.OperandStack.PushRef(clObj)
}

// field

// private native Field[] getDeclaredFields0(boolean publicOnly)
// getDeclaredFields0 (Z)[Ljava/lang/reflect/Field;
func getDeclaredFields0(f *rtdt.Frame) {
	thread := f.Thread()
	this := f.LocalVar.GetRef(0)
	publicOnly := f.LocalVar.GetInt(1) > 0
	//
	fieldMap := this.GetClzClass().FieldMap()
	pickedFields := make([]*marea.Field, 0, len(fieldMap))
	for _, f := range fieldMap {
		if !publicOnly || publicOnly && f.IsPublic() {
			pickedFields = append(pickedFields, f)
		}
	}
	//
	loader := f.Method().Class().DefineLoader()
	fieldClass := loader.Load(utils.CLASSNAME_Field)
	fieldArrClass := loader.Load("[" + utils.CLASSNAME_Field)
	fieldClzObj := fieldClass.GetClassObject()
	//todo bug
	fieldArray := marea.NewArrayA(fieldArrClass, int32(len(pickedFields)))
	f.OperandStack.PushRef(fieldArray)

	//Field(
	//  Class<?> declaringClass,
	//	String name,
	//	Class<?> type,
	//  int modifiers,
	//	int slot,
	//	String signature,
	//	byte[] annotations)

	fieldConstructor := fieldClass.Method(utils.METHODNAME_Init,
		"(Ljava/lang/Class;Ljava/lang/String;Ljava/lang/Class;IILjava/lang/String;[B)V")
	//debug, print all names
	fmt.Printf("[FIELD NAMES] class %s len%d", this.GetClzClass().ClassName(), len(pickedFields))
	for i, field := range pickedFields {
		fieldObj := marea.NewObject(fieldClass)
		fieldArray.ArrGetRefs()[i] = fieldObj

		ops := rtdt.NewOperandStack(8)
		ops.PushRef(fieldObj)                                  // this
		ops.PushRef(fieldClzObj)                               // declaring class
		ops.PushRef(marea.GetJavaString(field.Name(), loader)) // java name
		//debug
		typeName := cmn.ToClassName(field.Desc())
		fmt.Printf("[Type] %s\n", typeName)
		ops.PushRef(loader.Load(typeName).GetClassObject())    // type class
		ops.PushInt(int32(field.Flags()))                      // modifiers
		ops.PushInt(int32(field.VarIdx()))                     // slotid
		ops.PushRef(marea.GetJavaString(field.Desc(), loader)) // singature
		ops.PushRef(nil)                                       // annotations

		df := dummyFrame(ops, thread)
		thread.PushFrame(df)
		callMethod(fieldConstructor, df)
	}
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(f *rtdt.Frame) {
	thisClz := f.LocalVar.GetRef(0)
	if thisClz == nil {
		panic(utils.NullPointerException)
	}
	name := thisClz.GetClzClass().ClassName()
	isP := cmn.IsPrimitiveType(name)
	if isP {
		f.OperandStack.PushInt(1)
	} else {
		f.OperandStack.PushInt(0)
	}
}

//private native Field[]       getDeclaredFields0(boolean publicOnly);

//private native Method[]      getDeclaredMethods0(boolean publicOnly);

//private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);

//private native Class<?>[]   getDeclaredClasses0();
