package instructions

import (
	"jvmGo/jvm/lib"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/cmn"
)

func getfield(frame *rtdt.Frame) {
	// get *Field
	field := getFieldRefU16(frame)
	if field.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	stack := frame.OperandStack
	obj := stack.PopNonnilRef()

	utils.DIstrPrintf("GET FIELD %s %s %s\n", field.Name(), field.Desc(), field.Class().ClassName())

	i := field.VarIdx()
	switch field.Desc() {
	case "B", "C", "I", "S", "Z":
		v := obj.GetInt(i)
		stack.PushInt(v)
	case "D":
		v := obj.GetDouble(i)
		stack.PushDouble(v)
	case "J":
		v := obj.GetLong(i)
		stack.PushLong(v)
	case "F":
		v := obj.GetFloat(i)
		stack.PushFloat(v)
	default:
		// [,L...;
		v := obj.GetRef(i)
		stack.PushRef(v)
	}
}

func getstatic(frame *rtdt.Frame) {
	// get *Field
	field := getFieldRefU16(frame)
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

	utils.DIstrPrintf("PUT FIELD %s %s %s\n", field.Name(), field.Desc(), field.Class().ClassName())

	switch field.Desc() {
	case "B", "C", "I", "S", "Z":
		v := stack.PopInt()
		obj := stack.PopNonnilRef()
		obj.SetInt(v, i)
	case "D":
		v := stack.PopDouble()
		obj := stack.PopNonnilRef()
		obj.SetDouble(v, i)
	case "J":
		v := stack.PopLong()
		obj := stack.PopNonnilRef()
		obj.SetLong(v, i)
	case "F":
		v := stack.PopFloat()
		obj := stack.PopNonnilRef()
		obj.SetFloat(v, i)
	default:
		// [,L...;
		v := stack.PopRef()
		obj := stack.PopNonnilRef()
		obj.SetRef(v, i)
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

	utils.DIstrPrintf("PUT FIELD %s %s %s\n", field.Name(), field.Desc(), class.ClassName())

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
		class.SetRef(v, i)
	}
}

func new(frame *rtdt.Frame) {
	class := getClassRefU16(frame)
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
		c := loader.LoadArrayClass("[Z")
		obj = marea.NewArrayB(c, length) // use byte array
	case T_BYTE:
		c := loader.LoadArrayClass("[B")
		obj = marea.NewArrayB(c, length)
	case T_CHAR:
		c := loader.LoadArrayClass("[C")
		obj = marea.NewArrayC(c, length)
	case T_FLOAT:
		c := loader.LoadArrayClass("[F")
		obj = marea.NewArrayF(c, length)
	case T_DOUBLE:
		c := loader.LoadArrayClass("[D")
		obj = marea.NewArrayD(c, length)
	case T_SHORT:
		c := loader.LoadArrayClass("[S")
		obj = marea.NewArrayS(c, length)
	case T_INT:
		c := loader.LoadArrayClass("[I")
		obj = marea.NewArrayI(c, length)
	case T_LONG:
		c := loader.LoadArrayClass("[J")
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
	arrC := elec.DefineLoader().LoadArrayClass(arrName)
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
	if obj == nil {
		frame.OperandStack.PushInt(1)
		return
	}

	T := getClassRefU16(frame) // T, test class
	S := obj.Class()           // S, instance class
	v := marea.IsAssignable(S, T)
	if v {
		frame.OperandStack.PushInt(1)
	} else {
		frame.OperandStack.PushInt(0)
	}
}

func checkcast(frame *rtdt.Frame) {
	obj := frame.OperandStack.Top().Ref
	if obj == nil {
		return
	}

	T := getClassRefU16(frame)
	v := marea.IsAssignable(obj.Class(), T)
	if !v {
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
	t := f.Thread()
	// locate the method
	ind := f.ReadU16()
	cc := f.Method().Class() //current class
	cp := cc.ConstantPool()
	mr := cp.GetMethodRef(ind)
	m := mr.GetMethod()

	if m.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	if m.IsAbstract() {
		panic(utils.AbstractMethodError)
	}

	utils.DIstrPrintf("method name:%s desc:%s args:%d ret:%s class:%s\n",
		m.Name(), m.Desc(), m.ArgSlotNum(), m.RetD(), m.Class().ClassName())
	//pop objref
	pos := f.OperandStack.GetSize() - uint(m.ArgSlotNum()) - 1 // must be instance method
	objref := f.OperandStack.GetSlot(uint(pos)).Ref
	utils.DIstrPrintf("pos %d\n", pos)
	utils.DIstrPrintf("obj %d\n", objref)
	utils.DIstrPrintf("slots %d\n", f.OperandStack)

	if objref == nil {
		panic(utils.NullPointerException)
	}
	if m.IsProtected() && marea.IsDescandent(cc, m.Class()) &&
		objref.Class().PackageName() != cc.PackageName() && !marea.IsDescandent(objref.Class(), cc) {
		panic(utils.IllegalAccessError)
	}

	realMethod := marea.LookUpMethodVirtual(objref.Class(), cc, m.Name(), m.Desc())
	if realMethod.IsAbstract() {
		panic(utils.AbstractMethodError)
	}

	// call method
	utils.DIstrPrintf("INVOKE VIRTUAL, raw %s %s %s, REAL call %s %s %s\n",
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

	//utils.DIstrPrintf("real called method %s %s %s\n", m.Name(), m.Desc(), m.Class().ClassName())
	pos := f.OperandStack.GetSize() - uint(m.ArgSlotNum()) - 1 // must be instance method
	objref := f.OperandStack.GetSlot(uint(pos)).Ref

	if objref == nil {
		panic(utils.NullPointerException)
	}
	if m.IsProtected() && marea.IsDescandent(cc, m.Class()) &&
		objref.Class().PackageName() != cc.PackageName() && !marea.IsDescandent(objref.Class(), cc) {
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
	m := mr.Ref().LookUpMethodDirectly(mr.Name(), mr.Desc())

	if m == nil {
		panic(utils.NoSuchMethodError)
	}

	if !m.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	if m.Name() == "registerNatives" {
		utils.DIstrPrintf("\n") // debug??
	}

	// call method
	utils.DIstrPrintf("[INVOKE STATIC], %s %s %s\n",
		m.Name(), m.Desc(), m.Class().ClassName())
	callMethod(m, t)
}

func invokeinterface(f *rtdt.Frame) {
	t := f.Thread()
	// locate the method
	ind := f.ReadU16()
	f.ReadU16()              // for count and 0,historical
	cc := f.Method().Class() // current class
	cp := cc.ConstantPool()
	mr := cp.GetInterfaceMethodRef(ind)

	m := mr.GetMethod()

	if m.IsAbstract() {
		panic(utils.AbstractMethodError)
	}

	if m.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	//pop objref
	pos := f.OperandStack.GetSize() - uint(m.ArgSlotNum()) - 1 // must be instance method
	objref := f.OperandStack.GetSlot(uint(pos)).Ref
	//utils.DIstrPrintln("pos", pos)
	//utils.DIstrPrintln("obj", objref)
	//utils.DIstrPrintln("slots", f.OperandStack)

	if objref == nil {
		panic(utils.NullPointerException)
	}

	realMethod := marea.LookUpMethodVirtual(objref.Class(), cc, m.Name(), m.Desc())
	if realMethod.IsAbstract() {
		panic(utils.AbstractMethodError)
	}

	if !realMethod.IsPublic() {
		panic(utils.IllegalAccessError)
	}

	// call method
	utils.DIstrPrintf("INVOKE INTERFACE, raw m %s %s %s, REAL call %s %s %s\n",
		m.Name(), m.Desc(), m.Class().ClassName(),
		realMethod.Name(), realMethod.Desc(), realMethod.Class().ClassName())
	callMethod(realMethod, t)
}

// TODO invokedynamic

func callMethod(m *marea.Method, t *rtdt.Thread) {
	if m.IsNative() {
		lib.CallNative(m)
		return
	}
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
}
