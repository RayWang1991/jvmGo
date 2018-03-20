package instructions

import (
	"jvmGo/jvm/lib"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/cmn"
	"fmt"
)

func getfield(frame *rtdt.Frame) {
	// get *Field

	field := getFieldRefU16(frame)
	frame.Method().Class().DefineLoader().Initiate(field.Class())
	if field.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	stack := frame.OperandStack
	obj := stack.PopNonnilRef()

	//utils.DIstrPrintf("GET FIELD %s %s %s\n", field.Name(), field.Desc(), field.Class().ClassName())

	i := field.VarIdx()
	//debug
	utils.DIstrPrintf("GET FIELD %s %s %s ", field.Name(), field.Desc(), field.Class().ClassName())
	switch field.Desc() {
	case "B", "C", "I", "S", "Z":
		v := obj.GetInt(i)
		stack.PushInt(v)
		utils.DIstrPrintf("I %d\n", v)
	case "D":
		v := obj.GetDouble(i)
		stack.PushDouble(v)
		utils.DIstrPrintf("D %f\n", v)
	case "J":
		v := obj.GetLong(i)
		stack.PushLong(v)
		utils.DIstrPrintf("J %d\n", v)
	case "F":
		v := obj.GetFloat(i)
		stack.PushFloat(v)
		utils.DIstrPrintf("F %f\n", v)
	default:
		// [,L...;
		v := obj.GetRef(i)
		stack.PushRef(v)
		utils.DIstrPrintf("[L %s\n", v)
	}
}

func getstatic(frame *rtdt.Frame) {
	// get *Field
	field := getFieldRefU16(frame)
	frame.Method().Class().DefineLoader().Initiate(field.Class())
	if !field.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	class := field.Class()
	stack := frame.OperandStack

	utils.DIstrPrintf("GET STATIC %s %s %s\n", field.Name(), field.Desc(), class.ClassName())

	i := field.VarIdx()
	switch field.Desc() {
	case "B", "C", "I", "S", "Z":
		v := class.GetInt(i)
		stack.PushInt(v)
	case "D":
		v := class.GetDouble(i)
		stack.PushDouble(v)
	case "J":
		v := class.GetLong(i)
		stack.PushLong(v)
	case "F":
		v := class.GetFloat(i)
		stack.PushFloat(v)
	default:
		// [,L...;
		v := class.GetRef(i)
		utils.DIstrPrintf("[GET STATIC.REF] %s %s %s %s\n", field.Name(), field.Desc(), class.ClassName(), v)
		stack.PushRef(v)
	}
}

func putfield(frame *rtdt.Frame) {
	// get *Field
	field := getFieldRefU16(frame)
	if field.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}
	if field.IsFinal() && frame.Method().Name() != "<init>" {
		panic(utils.IllegalAccessError)
	}

	stack := frame.OperandStack
	i := field.VarIdx()

	//debug
	utils.DIstrPrintf("PUT FIELD %s %s %s index%d\n", field.Name(), field.Desc(), field.Class().ClassName(), i)

	switch field.Desc() {
	case "B", "C", "I", "S", "Z":
		v := stack.PopInt()
		obj := stack.PopNonnilRef()
		obj.SetInt(v, i)
		utils.DIstrPrintf("I %d\n", v)
	case "D":
		v := stack.PopDouble()
		obj := stack.PopNonnilRef()
		obj.SetDouble(v, i)
		utils.DIstrPrintf("D %f\n", v)
	case "J":
		v := stack.PopLong()
		obj := stack.PopNonnilRef()
		obj.SetLong(v, i)
		utils.DIstrPrintf("J %d\n", v)
	case "F":
		v := stack.PopFloat()
		obj := stack.PopNonnilRef()
		obj.SetFloat(v, i)
		utils.DIstrPrintf("F %f\n", v)
	default:
		// [,L...;
		v := stack.PopRef()
		obj := stack.PopNonnilRef()
		obj.SetRef(v, i)
		utils.DIstrPrintf("[L %s\n", v)
	}
}

func putstatic(frame *rtdt.Frame) {
	// get *Field

	field := getFieldRefU16(frame)
	if !field.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}
	if field.IsFinal() && frame.Method().Name() != cmn.METHOD_CLINIT_NAME {
		utils.DIstrPrintf("%s\n", frame.Method().Name())
		panic(utils.IllegalAccessError)
	}

	class := field.Class()
	stack := frame.OperandStack
	i := field.VarIdx()

	utils.DIstrPrintf("[PUT FIELD STATIC] %s %s %s\n", field.Name(), field.Desc(), class.ClassName())

	switch field.Desc() {
	case "B", "C", "I", "S", "Z":
		v := stack.PopInt()
		class.SetInt(v, i)
	case "D":
		v := stack.PopDouble()
		class.SetDouble(v, i)
	case "J":
		v := stack.PopLong()
		class.SetLong(v, i)
	case "F":
		v := stack.PopFloat()
		class.SetFloat(v, i)
	default:
		// [,L...;
		v := stack.PopRef()
		utils.DIstrPrintf("[PUT FIELD STATIC.REF] %s %s %s %s\n",
			field.Name(), field.Desc(), class.ClassName(), v)
		class.SetRef(v, i)
	}
}

func new(frame *rtdt.Frame) {
	class := getClassRefU16(frame)
	frame.Method().Class().DefineLoader().Initiate(class)
	utils.DIstrPrintf("[NEW]Get class %s\n", class.ClassName())
	if class.IsAbstract() || class.IsInterface() || class.IsArray() {
		panic(utils.InstantiationError)
	}
	utils.DIstrPrintf("new obj from class: %s\n", class.ClassName())
	obj := marea.NewObject(class)
	frame.OperandStack.PushRef(obj)
}

func newarray(frame *rtdt.Frame) {
	const (
		T_BOOLEAN = iota + 4
		T_CHAR
		T_FLOAT
		T_DOUBLE
		T_BYTE
		T_SHORT
		T_INT
		T_LONG
	)
	atype := frame.ReadU8()
	length := frame.OperandStack.PopInt()
	if length < 0 {
		panic(utils.NegativeArraySizeException)
	}
	var obj *marea.Object
	loader := frame.Method().Class().InitLoader() // TODO, may be use other loader
	switch atype {
	case T_BOOLEAN:
		c := loader.Load("[Z")
		obj = marea.NewArrayB(c, length) // use byte array
	case T_BYTE:
		c := loader.Load("[B")
		obj = marea.NewArrayB(c, length)
	case T_CHAR:
		c := loader.Load("[C")
		obj = marea.NewArrayC(c, length)
	case T_FLOAT:
		c := loader.Load("[F")
		obj = marea.NewArrayF(c, length)
	case T_DOUBLE:
		c := loader.Load("[D")
		obj = marea.NewArrayD(c, length)
	case T_SHORT:
		c := loader.Load("[S")
		obj = marea.NewArrayS(c, length)
	case T_INT:
		c := loader.Load("[I")
		obj = marea.NewArrayI(c, length)
	case T_LONG:
		c := loader.Load("[J")
		obj = marea.NewArrayJ(c, length)
	}
	utils.DIstrPrintf("New Primative array %s \n", obj.Class().ClassName())
	frame.OperandStack.PushRef(obj)
}

func anewarray(frame *rtdt.Frame) {
	utils.DIstrPrintf("Enter anew array\n")
	length := frame.OperandStack.PopInt()
	if length < 0 {
		panic(utils.NegativeArraySizeException)
	}
	elec := getClassRefU16(frame)
	arrName := "[" + elec.ClassName()
	arrC := elec.DefineLoader().Load(arrName)
	obj := marea.NewArrayA(arrC, length)
	frame.OperandStack.PushRef(obj)
	utils.DIstrPrintf("Put ref %s\n", arrC.ClassName())
}

func multianewarray(frame *rtdt.Frame) {
	elec := getClassRefU16(frame)
	dim := frame.ReadU8() // dim must >= 1
	if dim < 1 { // 0 - dim array ??? can not be
		panic("0 dim array")
	}

	// preparation
	eleClsArr := make([]*marea.Class, dim) // want element class array
	eleCntArr := make([]int32, dim)        // element count array
	eleName := elec.ClassName()

	maxL := 1
	for i := uint8(0); i < dim; i++ {
		length := frame.OperandStack.PopInt()
		if length < 0 {
			panic(utils.NegativeArraySizeException)
		}
		if i < dim-1 {
			maxL *= int(length)
		}
		eleClsArr[i] = elec
		eleCntArr[dim-1-i] = length
		eleName = eleName[1:]
		if i < dim-1 {
			elec = elec.DefineLoader().Load(eleName)
		}
	}

	// bfs allocate and link all obj arrays, TODO dfs maybe
	workList := make([]*marea.Object, 0, maxL)
	tempList := make([]*marea.Object, 0, maxL)

	ret := marea.NewArrayA(eleClsArr[0], eleCntArr[0])
	workList = append(workList, ret)

	for i := uint8(0); i < dim-1; i++ {
		// for certain array type
		// want num is total
		// current num is cnt (for array length)
		arrCnt := eleCntArr[i]
		nextCnt := eleCntArr[i+1]
		cls := eleClsArr[i+1]
		for _, container := range workList {
			for j := int32(0); j < arrCnt; j++ {
				a := marea.NewArray(cls, nextCnt)
				container.ArrGetRefs()[j] = a
				tempList = append(tempList, a)
			}
		}
		workList = tempList
		tempList = tempList[:0]
	}

	//utils.DIstrPrintf("result element is %s\n", ret.ArrGetRefs()[0].ArrGetRefs()[0].Class().ClassName())
	frame.OperandStack.PushRef(ret)
}

func arraylength(frame *rtdt.Frame) {
	arref := frame.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	frame.OperandStack.PushInt(arref.ArrayLength())
}

func instanceof(frame *rtdt.Frame) {
	obj := frame.OperandStack.PopRef()
	T := getClassRefU16(frame) // T, test class

	if obj == nil {
		frame.OperandStack.PushInt(0)
		return
	}

	S := obj.Class() // S, instance class
	v := marea.IsAssignable(S, T)
	if v {
		frame.OperandStack.PushInt(1)
	} else {
		frame.OperandStack.PushInt(0)
	}
}

func checkcast(frame *rtdt.Frame) {
	obj := frame.OperandStack.Top().Ref
	// must consume u16

	T := getClassRefU16(frame)
	if obj == nil {
		return
	}

	v := marea.IsAssignable(obj.Class(), T)
	if !v {
		fmt.Printf("objclz %s T %s obj%s\n",
			obj.Class().ClassName(), T.ClassName(), obj.Data().([]*marea.Object))
		panic(utils.ClassCastException)
	}
}

func getClassRefU16(f *rtdt.Frame) *marea.Class {
	idx := f.ReadU16()
	m := f.Method()
	c := m.Class()
	cp := c.ConstantPool()
	ref := cp.GetClassRef(idx)
	return ref.Ref()
}

func getFieldRefU16(f *rtdt.Frame) *marea.Field {
	inx := f.ReadU16()
	from := f.Method().Class()
	ref := from.ConstantPool().GetFieldRef(inx)
	return ref.GetField()
}

// invoke family
func invokevirtual(f *rtdt.Frame) {
	utils.DIstrPrintf("[INVOKE VIRTUAL]\n")
	t := f.Thread()
	// locate the method
	ind := f.ReadU16()
	cc := f.Method().Class() //current class
	cp := cc.ConstantPool()
	mr := cp.GetMethodRef(ind)
	m := mr.GetMethod()

	//debug
	utils.DIstrPrintf("method name:%s desc:%s args:%d ret:%s class:%s current %s\n",
		m.Name(), m.Desc(), m.ArgSlotNum(), m.RetD(), m.Class().ClassName(), cc.ClassName())
	if m.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	//pop objref
	pos := uint(m.ArgSlotNum())
	objref := f.OperandStack.GetSlot(uint(pos)).Ref
	utils.DIstrPrintf("pos %d\n", pos)
	utils.DIstrPrintf("obj %v\n", objref)
	utils.DIstrPrintf("slots %v\n", f.OperandStack)

	if objref == nil {
		//debug
		fmt.Printf("Local %s Stack %s\n", f.LocalVar, f.OperandStack)

		panic(utils.NullPointerException)
	}

	if m.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}
	//if m.IsProtected() && marea.IsDescandent(cc, m.Class()) &&
	//	objref.Class().PackageName() != cc.PackageName() && !marea.IsDescandent(objref.Class(), cc) {
	//	panic(utils.IllegalAccessError)
	//}

	//debug
	utils.DIstrPrintf("[REAL] name:%s desc:%s call:%s from:%s\n",
		m.Name(), m.Desc(), objref.Class().ClassName(), cc.ClassName())
	realMethod := marea.LookUpMethodVirtual(objref.Class(), cc, m.Name(), m.Desc())

	if realMethod.IsProtected() && marea.IsDescandent(cc, realMethod.Class()) &&
		objref.Class().PackageName() != cc.PackageName() && !marea.IsDescandent(objref.Class(), cc) {
		//todo cloneable for array class obj, use native clone in Object class
		if !cmn.IsArray(objref.Class().ClassName()) || realMethod.Name() != "clone" {
			fmt.Printf("OBJ ref is %s %s\n", objref.Class().ClassName(), objref.GetClzClass().ClassName())
			panic(utils.IllegalAccessError)
		}
	}

	if nil == realMethod || realMethod.IsAbstract() {
		panic(utils.AbstractMethodError)
	}

	// call method
	utils.DIstrPrintf("raw %s %s %s, REAL call %s %s %s\n",
		m.Name(), m.Desc(), m.Class().ClassName(),
		realMethod.Name(), realMethod.Desc(), realMethod.Class().ClassName())
	callMethod(realMethod, t)
}

// invokespecial
func invokespecial(f *rtdt.Frame) {
	t := f.Thread()
	// locate the method
	ind := f.ReadU16()
	cc := f.Method().Class() //current class
	cp := cc.ConstantPool()
	mr := cp.GetMethodRef(ind)

	//utils.DIstrPrintln("invokespecial")
	//utils.DIstrPrintf("method name:%s desc:%s class:%s\n", mr.Name(), mr.Desc(), mr.ClassName())

	var m *marea.Method
	if mr.Name() != "<init>" && cc.IsSuper() && cc.Superclass() == mr.Ref() {
		m = cc.Superclass().LookUpMethod(mr.Name(), mr.Desc()) //look up method recusively in mr.ref(current class's super class)
	} else { // "<init>" or private method
		m = mr.Ref().LookUpMethodDirectly(mr.Name(), mr.Desc())
	}

	if m == nil {
		panic(utils.NoSuchMethodError)
	}

	if m.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	if m.IsAbstract() {
		panic(utils.AbstractMethodError)
	}

	//debug
	utils.DIstrPrintf("method name:%s desc:%s args:%d ret:%s class:%s\n",
		m.Name(), m.Desc(), m.ArgSlotNum(), m.RetD(), m.Class().ClassName())
	//utils.DIstrPrintf("real called method %s %s %s\n", m.Name(), m.Desc(), m.Class().ClassName())
	pos := m.ArgSlotNum()
	objref := f.OperandStack.GetSlot(uint(pos)).Ref

	if objref == nil {
		panic(utils.NullPointerException)
	}
	if m.IsProtected() && marea.IsDescandent(cc, m.Class()) &&
		m.Class().PackageName() != cc.PackageName() && !marea.IsDescandent(objref.Class(), cc) {
		//fmt.Printf("objref %s cc %s, callingMethod %s", objref.Class().ClassName(), cc.ClassName(), m.Class().ClassName())
		panic(utils.IllegalAccessError)
	}

	// call method
	utils.DIstrPrintf("[INVOKE SPECIAL], raw %s %s %s, REAL call %s %s %s\n",
		mr.Name(), mr.Desc(), mr.ClassName(),
		m.Name(), m.Desc(), m.Class().ClassName())
	callMethod(m, t)
}

func invokestatic(f *rtdt.Frame) {
	t := f.Thread()
	// locate the method
	ind := f.ReadU16()
	cc := f.Method().Class() //current class
	cp := cc.ConstantPool()
	mr := cp.GetMethodRef(ind)
	mClz := mr.Ref()
	f.Method().Class().DefineLoader().Initiate(mClz)
	m := mClz.LookUpMethodDirectly(mr.Name(), mr.Desc())

	//debug
	utils.DIstrPrintf("method name:%s desc:%s args:%d ret:%s class:%s\n",
		m.Name(), m.Desc(), m.ArgSlotNum(), m.RetD(), m.Class().ClassName())

	if m == nil {
		panic(utils.NoSuchMethodError)
	}

	if !m.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	// call method
	utils.DIstrPrintf("[INVOKE STATIC], %s %s %s\n",
		m.Name(), m.Desc(), m.Class().ClassName())
	callMethod(m, t)
}

func invokeinterface(f *rtdt.Frame) {
	utils.DIstrPrintf("[INVOKE INTERFACE]\n")
	t := f.Thread()
	// locate the method
	ind := f.ReadU16()
	f.ReadU16()              // for count and 0,historical
	cc := f.Method().Class() // current class
	cp := cc.ConstantPool()
	mr := cp.GetInterfaceMethodRef(ind)

	m := mr.GetMethod()
	//debug
	utils.DIstrPrintf("METHOD NAME %s class %s\n", m.Name(), m.Class().ClassName())

	if m.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	//pop objref
	pos := uint(m.ArgSlotNum()) // must be instance method
	objref := f.OperandStack.GetSlot(uint(pos)).Ref
	utils.DIstrPrintf("pos %d\n", pos)
	utils.DIstrPrintf("obj %s\n", objref)
	utils.DIstrPrintf("slots %s\n", f.OperandStack)

	if objref == nil {
		panic(utils.NullPointerException)
	}
	// debug
	utils.DIstrPrintf("CLASS %s\n", objref.Class().ClassName())

	realMethod := marea.LookUpMethodVirtual(objref.Class(), cc, m.Name(), m.Desc())
	if realMethod.IsAbstract() {
		panic(utils.AbstractMethodError)
	}

	if realMethod.IsAbstract() {
		fmt.Printf("%s %s\n", realMethod.Class().ClassName(), realMethod.Name())
		panic(utils.AbstractMethodError)
	}

	if !realMethod.IsPublic() {
		panic(utils.IllegalAccessError)
	}

	// call method
	utils.DIstrPrintf(" raw m %s %s %s, REAL call %s %s %s\n",
		m.Name(), m.Desc(), m.Class().ClassName(),
		realMethod.Name(), realMethod.Desc(), realMethod.Class().ClassName())
	callMethod(realMethod, t)
}

// TODO invokedynamic

func callMethod(m *marea.Method, t *rtdt.Thread) {
	utils.DIstrPrintf("[CALL real] %s %s\n", m.Name(), m.Class().ClassName())
	if m.IsNative() {
		utils.DIstrPrintf("[NATIVE] %s %s\n", m.Name(), m.Desc())
		m.SetMaxLocalVars(uint16(m.ArgSlotNum() + 1))
		m.SetMaxStackDep(16) //TODO, variable
		// inject return code
		switch m.RetD()[0] {
		case 'V':
			m.SetCode([]byte{cmn.OPCODE_rreturn})
		case 'D':
			m.SetCode([]byte{cmn.OPCODE_dreturn})
		case 'F':
			m.SetCode([]byte{cmn.OPCODE_freturn})
		case 'J':
			m.SetCode([]byte{cmn.OPCODE_lreturn})
		case 'L', '[':
			m.SetCode([]byte{cmn.OPCODE_areturn})
		default:
			m.SetCode([]byte{cmn.OPCODE_ireturn})
		}
		setUpCallingFrame(t, m)
		lib.CallNative(m, t)
		return
	}
	setUpCallingFrame(t, m)
}

func setUpCallingFrame(t *rtdt.Thread, m *marea.Method) {
	f := t.CurrentFrame()
	nf := rtdt.NewFrame(m, t)
	t.PushFrame(nf)
	i := m.ArgSlotNum()
	if m.IsStatic() {
		i--
	}
	for ; i >= 0; i-- {
		slot := f.OperandStack.PopSlot()
		nf.LocalVar.SetSlot(slot, uint(i))
	}
	//debug
	f = nf
	utils.DIstrPrintf("stack %s\n locals %s\n", f.OperandStack, f.LocalVar)

	if utils.CallTraceFlag {
		fmt.Printf("%s.%s()\n", m.Class().ClassName(), m.Name())
	}

	if f.Method().Name() == "fillInStackTrace" {
		t.PrintStack()
	}

	//todo hack
	if f.Method().Name() == "loadLibrary" && f.Method().Class().ClassName() == "java/lang/System" {
		f.Method().SetCode([]byte{cmn.OPCODE_rreturn})
	}
}

// TODO monitorenter
func monitorenter(f *rtdt.Frame) {
	f.OperandStack.PopNonnilRef()
}

// TODO monitorexit
func monitorexit(f *rtdt.Frame) {
	f.OperandStack.PopNonnilRef()
}
