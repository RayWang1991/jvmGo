package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/cmn"
)

func init() {
	register(utils.CLASSNAME_Class, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	register(utils.CLASSNAME_Class, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	register(utils.CLASSNAME_Class, "getName0", "()Ljava/lang/String;", getName0)
	register(utils.CLASSNAME_Class, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	register(utils.CLASSNAME_Class, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
	register(utils.CLASSNAME_Class, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructores0)
	register(utils.CLASSNAME_Class, "isPrimitive", "()Z", isPrimitive)
	register(utils.CLASSNAME_Class, "isAssignableFrom", "(Ljava/lang/Class;)Z", isAssignableFrom)
	register(utils.CLASSNAME_Class, "isInterface", "()Z", isInterface)
	register(utils.CLASSNAME_Class, "getModifiers", "()I", getModifiers)
	register(utils.CLASSNAME_Class, "getSuperclass", "()Ljava/lang/Class;", getSuperclass)
	register(utils.CLASSNAME_Class, "isArray", "()Z", isArray)
	register(utils.CLASSNAME_Class, "getComponentType", "()Ljava/lang/Class;", getComponentType)
	register(utils.CLASSNAME_Class, "getEnclosingMethod0", "()[Ljava/lang/Object;", getEnclosingMethod0)
	register(utils.CLASSNAME_Class, "getDeclaringClass0", "()Ljava/lang/Class;", getDeclaringClass0)
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
	//debug
	utils.DIstrPrintf("name %s\n", name)
	f.OperandStack.PushRef(str)
}

// private static native Class<?> forName0(String name, boolean initialize, ClassLoader loader, Class<?> caller)
// (Ljava/lang/String;ZLjava/lang/ClassLoader/;Ljava/lang/Class;)Ljava/lang/Class;
func forName0(f *rtdt.Frame) {
	jName := f.LocalVar.GetRef(0)
	goName := cmn.ToSlash(marea.GetGoString(jName))

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
	utils.DIstrPrintf("[FIELD NAMES] class %s len%d", this.GetClzClass().ClassName(), len(pickedFields))
	for i, field := range pickedFields {
		fieldObj := marea.NewObject(fieldClass)
		fieldArray.ArrGetRefs()[i] = fieldObj

		ops := rtdt.NewOperandStack(8)
		ops.PushRef(fieldObj)                                  // this
		ops.PushRef(fieldClzObj)                               // declaring class
		ops.PushRef(marea.GetJavaString(field.Name(), loader)) // java name
		typeName := cmn.ToClassName(field.Desc())
		//debug
		utils.DIstrPrintf("[Type] %s\n", typeName)
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

// public native boolean isAssignableFrom(Class<?> cls);
// (Ljava/lang/Class;)Z
func isAssignableFrom(f *rtdt.Frame) {
	thisClzObj := f.LocalVar.GetRef(0)
	thisClz := thisClzObj.GetClzClass()
	otherClzOjb := f.LocalVar.GetRef(1)
	t := otherClzOjb.GetClzClass()

	// check if is primitive class
	for _, p := range cmn.PrimitiveNames {
		if thisClz.ClassName() == p {
			if thisClz == t {
				f.OperandStack.PushInt(1)
			} else {
				f.OperandStack.PushInt(0)
			}
			return
		}
	}
	if marea.IsDescandent(t, thisClz) {
		f.OperandStack.PushInt(1)
	} else {
		f.OperandStack.PushInt(0)
	}
}

// public native boolean isInterface();
// ()Z
func isInterface(f *rtdt.Frame) {
	thisClzObj := f.LocalVar.GetRef(0)
	clz := thisClzObj.GetClzClass()
	if clz.IsInterface() {
		f.OperandStack.PushInt(1)
	} else {
		f.OperandStack.PushInt(0)
	}
}

// private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Constructor;
func getDeclaredConstructores0(f *rtdt.Frame) {
	thisClzObj := f.LocalVar.GetRef(0)
	clz := thisClzObj.GetClzClass()
	publicOnly := f.LocalVar.GetInt(1) > 0

	constructors := make([]*marea.Method, 0, 4) // todo, 4 is bigger enough for most class
	for _, m := range clz.MethodMap() {
		if m.Name() == utils.METHODNAME_Init && (!publicOnly || m.IsPublic()) {
			constructors = append(constructors, m)
		}
	}

	loader := f.Method().Class().DefineLoader()
	consClz := loader.Load(utils.CLASSNAME_Constructor)
	consArrClz := loader.Load("[" + utils.CLASSNAME_Constructor)
	consArr := marea.NewArrayA(consArrClz, int32(len(constructors)))
	f.OperandStack.PushRef(consArr)
	//clzClz := loader.Load(utils.CLASSNAME_Class)
	clzArrClz := loader.Load("[" + utils.CLASSNAME_Class)
	thread := f.Thread()
	initM := consClz.GetMethodDirect(utils.METHODNAME_Init,
		"(Ljava/lang/Class;"+
			"[Ljava/lang/Class;"+
			"[Ljava/lang/Class;"+
			"II"+
			"Ljava/lang/String;"+
			"[B[B)V")

	//Constructor(Class<T> declaringClass,
	//	Class<?>[] parameterTypes,
	//	Class<?>[] checkedExceptions,
	//	int modifiers,
	//	int slot,
	//	String signature,
	//	byte[] annotations,
	//	byte[] parameterAnnotations)

	for i, cons := range constructors {
		consObj := marea.NewObject(consClz)
		consArr.ArrGetRefs()[i] = consObj

		ops := rtdt.NewOperandStack(9)
		ops.PushRef(consObj)    // this
		ops.PushRef(thisClzObj) // declaring class

		// parameter types
		pts := cons.ParameterTypes()
		ptArray := marea.NewArrayA(clzArrClz, int32(len(pts)))
		for i, p := range pts {
			ptArray.ArrGetRefs()[i] = p.GetClassObject()
		}
		ops.PushRef(ptArray) // parameter Types

		// exception types
		eps := cons.ExceptionTypes()
		epArray := marea.NewArray(clzArrClz, int32(len(eps)))
		for i, e := range eps {
			epArray.ArrGetRefs()[i] = e.GetClassObject()
		}
		ops.PushRef(epArray)             // checkExceptions
		ops.PushInt(int32(cons.Flags())) // modifiers
		ops.PushInt(int32(0))            // todo slotid
		ops.PushRef(nil)                 // todo signature unsupported
		ops.PushRef(nil)                 // todo anotations
		desc := cons.Desc()
		bs := []byte(desc)
		descBA := marea.NewArrayB(loader.Load("[B"), int32(len(bs)))
		is := descBA.ArrGetBytes()
		for j, b := range bs {
			is[j] = int8(b)
		}
		ops.PushRef(descBA) // todo parameterAnnotations, record desc here

		df := dummyFrame(ops, thread)
		thread.PushFrame(df)
		callMethod(initM, df)
	}
}

//public native int getModifiers();
//()I
func getModifiers(f *rtdt.Frame) {
	cls := f.LocalVar.GetRef(0)
	flags := int32(cls.GetClzClass().GetFlags())
	f.OperandStack.PushInt(flags)
}

//public native Class<? super T> getSuperclass();
//()Ljava/lang/Class;
func getSuperclass(f *rtdt.Frame) {
	clsObj := f.LocalVar.GetRef(0)
	clz := clsObj.GetClzClass()
	superClz := clz.Superclass()
	if superClz == nil {
		f.OperandStack.PushRef(nil)
	} else {
		supClzObj := clz.Superclass().GetClassObject()
		f.OperandStack.PushRef(supClzObj)
	}
}

//private native Field[]       getDeclaredFields0(boolean publicOnly);

//private native Method[]      getDeclaredMethods0(boolean publicOnly);

//private native Class<?>[]   getDeclaredClasses0();

// public native boolean isArray();
// ()Z
func isArray(f *rtdt.Frame) {
	this := f.LocalVar.GetRef(0)
	cls := this.GetClzClass()
	var res int32
	if cmn.IsArray(cls.ClassName()) {
		res = 1
	} else {
		res = 0
	}
	f.OperandStack.PushInt(res)
}

//public native Class<?> getComponentType();
// ()Ljava/lang/Class;
func getComponentType(f *rtdt.Frame) {
	this := f.LocalVar.GetRef(0)
	cls := this.GetClzClass()
	// must be array type
	eleNRaw := cmn.ElementName(cls.ClassName())
	eleN := cmn.ToClassName(eleNRaw)
	eleC := f.Method().Class().DefineLoader().Load(eleN)
	eleCObj := eleC.GetClassObject()
	f.OperandStack.PushRef(eleCObj)
}

//private native Object[] getEnclosingMethod0();
// ()[Ljava/lang/Object;
func getEnclosingMethod0(f *rtdt.Frame) {
	//todo
	f.OperandStack.PushRef(nil)
}


// private native Class<?> getDeclaringClass0();
// ()Ljava/lang/Class;
func getDeclaringClass0(f *rtdt.Frame){
	//todo
	f.OperandStack.PushRef(nil)
}