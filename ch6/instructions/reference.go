package instructions

import (
	"jvmGo/ch6/rtdt"
	"jvmGo/ch6/utils"
	"jvmGo/ch6/marea"
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
	// TODO
	m := f.Method()
	c := m.Class()
	cp := c.ConstantPool()
	ref := cp.GetClassRef(idx)
	//ref := f.Method().Class().ConstantPool().GetClassRef(idx)
	return ref.Ref()
}

func getFieldRefU16(f *rtdt.Frame) *marea.Field {
	inx := f.ReadU16()
	from := f.Method().Class()
	ref := from.ConstantPool().GetFieldRef(inx)
	return ref.GetField()
}
