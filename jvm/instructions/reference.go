package instructions

import (
	"fmt"
	"jvmGo/ch6/marea"
	"jvmGo/ch6/rtdt"
	"jvmGo/ch6/utils"
)

func getfield(frame *rtdt.Frame) {
	// get *Field
	field := getFieldRefU16(frame)
	if field.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	stack := frame.OperandStack
	obj := stack.PopNonnilRef()

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
	if field.IsFinal() && frame.Method().Name() != "<cinit>" {
		panic(utils.IllegalAccessError)
	}

	class := field.Class()
	stack := frame.OperandStack
	i := field.VarIdx()

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
	//fmt.Println("new obj from class", class.ClassName())
	obj := marea.NewObject(class)
	frame.OperandStack.PushRef(obj)
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

	fmt.Printf("method name:%s desc:%s args:%d ret:%s class:%s\n",
		m.Name(), m.Desc(), m.ArgSlotNum(), m.RetD(), m.Class().ClassName())
	//pop objref
	pos := f.OperandStack.GetSize() - uint(m.ArgSlotNum()) - 1 // must be instance method
	objref := f.OperandStack.GetSlot(uint(pos)).Ref
	fmt.Println("pos", pos)
	fmt.Println("obj", objref)
	fmt.Println("slots", f.OperandStack)

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
	nf := rtdt.NewFrame(realMethod, t)
	t.PushFrame(nf)
	for i := realMethod.ArgSlotNum(); i >= 0; i-- {
		slot := f.OperandStack.PopSlot()
		nf.LocalVar.SetSlot(slot, uint(i))
	}
}

// invokespecial
func invokespecial(f *rtdt.Frame) {
	t := f.Thread()
	// locate the method
	ind := f.ReadU16()
	cc := f.Method().Class() //current class
	cp := cc.ConstantPool()
	mr := cp.GetMethodRef(ind)

	//fmt.Println("invokespecial")
	//fmt.Printf("method name:%s desc:%s class:%s\n", mr.Name(), mr.Desc(), mr.ClassName())

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

	//fmt.Printf("real called method %s %s %s\n", m.Name(), m.Desc(), m.Class().ClassName())
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
	nf := rtdt.NewFrame(m, t)
	t.PushFrame(nf)
	for i := m.ArgSlotNum(); i >= 0; i-- {
		slot := f.OperandStack.PopSlot()
		nf.LocalVar.SetSlot(slot, uint(i))
	}
	//fmt.Println("New Frame's locals", nf.LocalVar)
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

	// call method
	nf := rtdt.NewFrame(m, t)
	t.PushFrame(nf)
	for i := m.ArgSlotNum() - 1; i >= 0; i-- {
		slot := f.OperandStack.PopSlot()
		nf.LocalVar.SetSlot(slot, uint(i))
	}
}

func invokeinterface(f *rtdt.Frame) {
	t := f.Thread()
	// locate the method
	ind := f.ReadU16()
	f.ReadU16()              // for count and 0,historical
	cc := f.Method().Class() //current class
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
	fmt.Println("pos", pos)
	fmt.Println("obj", objref)
	fmt.Println("slots", f.OperandStack)

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
	nf := rtdt.NewFrame(realMethod, t)
	t.PushFrame(nf)
	for i := realMethod.ArgSlotNum(); i >= 0; i-- {
		slot := f.OperandStack.PopSlot()
		nf.LocalVar.SetSlot(slot, uint(i))
	}
}

// TODO invokedynamic
