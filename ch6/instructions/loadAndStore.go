package instructions

import "jvmGo/ch6/rtdata"

// Loads
// iload
func iload(f *rtdata.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetInt(uint(i))
	f.OperandStack.PushInt(v)
}
func iload_0(f *rtdata.Frame) {
	v := f.LocalVar.GetInt(uint(0))
	f.OperandStack.PushInt(v)
}
func iload_1(f *rtdata.Frame) {
	v := f.LocalVar.GetInt(uint(1))
	f.OperandStack.PushInt(v)
}
func iload_2(f *rtdata.Frame) {
	v := f.LocalVar.GetInt(uint(2))
	f.OperandStack.PushInt(v)
}
func iload_3(f *rtdata.Frame) {
	v := f.LocalVar.GetInt(uint(3))
	f.OperandStack.PushInt(v)
}

//lload
func lload(f *rtdata.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetLong(uint(i))
	f.OperandStack.PushLong(v)
}
func lload_0(f *rtdata.Frame) {
	v := f.LocalVar.GetLong(uint(0))
	f.OperandStack.PushLong(v)
}
func lload_1(f *rtdata.Frame) {
	v := f.LocalVar.GetLong(uint(1))
	f.OperandStack.PushLong(v)
}
func lload_2(f *rtdata.Frame) {
	v := f.LocalVar.GetLong(uint(2))
	f.OperandStack.PushLong(v)
}
func lload_3(f *rtdata.Frame) {
	v := f.LocalVar.GetLong(uint(3))
	f.OperandStack.PushLong(v)
}

// fload
func fload(f *rtdata.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetFloat(uint(i))
	f.OperandStack.PushFloat(v)
}
func fload_0(f *rtdata.Frame) {
	v := f.LocalVar.GetFloat(uint(0))
	f.OperandStack.PushFloat(v)
}
func fload_1(f *rtdata.Frame) {
	v := f.LocalVar.GetFloat(uint(1))
	f.OperandStack.PushFloat(v)
}
func fload_2(f *rtdata.Frame) {
	v := f.LocalVar.GetFloat(uint(2))
	f.OperandStack.PushFloat(v)
}
func fload_3(f *rtdata.Frame) {
	v := f.LocalVar.GetFloat(uint(3))
	f.OperandStack.PushFloat(v)
}

// dload
func dload(f *rtdata.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetDouble(uint(i))
	f.OperandStack.PushDouble(v)
}
func dload_0(f *rtdata.Frame) {
	v := f.LocalVar.GetDouble(uint(0))
	f.OperandStack.PushDouble(v)
}
func dload_1(f *rtdata.Frame) {
	v := f.LocalVar.GetDouble(uint(1))
	f.OperandStack.PushDouble(v)
}
func dload_2(f *rtdata.Frame) {
	v := f.LocalVar.GetDouble(uint(2))
	f.OperandStack.PushDouble(v)
}
func dload_3(f *rtdata.Frame) {
	v := f.LocalVar.GetDouble(uint(3))
	f.OperandStack.PushDouble(v)
}

func aload(f *rtdata.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetRef(uint(i))
	f.OperandStack.PushRef(v)
}
func aload_0(f *rtdata.Frame) {
	v := f.LocalVar.GetRef(uint(0))
	f.OperandStack.PushRef(v)
}
func aload_1(f *rtdata.Frame) {
	v := f.LocalVar.GetRef(uint(1))
	f.OperandStack.PushRef(v)
}
func aload_2(f *rtdata.Frame) {
	v := f.LocalVar.GetRef(uint(2))
	f.OperandStack.PushRef(v)
}
func aload_3(f *rtdata.Frame) {
	v := f.LocalVar.GetRef(uint(3))
	f.OperandStack.PushRef(v)
}

// TODO load int,float,refs... from array

// Stores
func istore(f *rtdata.Frame) {
	r := f.OperandStack.PopInt()
	i := f.ReadU8()
	f.LocalVar.SetInt(r, uint(i))
}
func istore_0(f *rtdata.Frame) {
	r := f.OperandStack.PopInt()
	f.LocalVar.SetInt(r, 0)
}
func istore_1(f *rtdata.Frame) {
	r := f.OperandStack.PopInt()
	f.LocalVar.SetInt(r, 1)
}
func istore_2(f *rtdata.Frame) {
	r := f.OperandStack.PopInt()
	f.LocalVar.SetInt(r, 2)
}
func istore_3(f *rtdata.Frame) {
	r := f.OperandStack.PopInt()
	f.LocalVar.SetInt(r, 3)
}

func dstore(f *rtdata.Frame) {
	i := f.ReadU8()
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, uint(i))
}
func dstore_0(f *rtdata.Frame) {
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, 0)
}
func dstore_1(f *rtdata.Frame) {
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, 1)
}
func dstore_2(f *rtdata.Frame) {
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, 2)
}
func dstore_3(f *rtdata.Frame) {
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, 3)
}

func fstore(f *rtdata.Frame) {
	i := f.ReadU8()
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, uint(i))
}
func fstore_0(f *rtdata.Frame) {
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, 0)
}
func fstore_1(f *rtdata.Frame) {
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, 1)
}
func fstore_2(f *rtdata.Frame) {
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, 2)
}
func fstore_3(f *rtdata.Frame) {
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, 3)
}

func astore(f *rtdata.Frame) {
	i := f.ReadU8()
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, uint(i))
}
func astore_0(f *rtdata.Frame) {
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, 0)
}
func astore_1(f *rtdata.Frame) {
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, 1)
}
func astore_2(f *rtdata.Frame) {
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, 2)
}
func astore_3(f *rtdata.Frame) {
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, 3)
}

func lstore(f *rtdata.Frame) {
	i := f.ReadU8()
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, uint(i))
}
func lstore_0(f *rtdata.Frame) {
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, 0)
}
func lstore_1(f *rtdata.Frame) {
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, 1)
}
func lstore_2(f *rtdata.Frame) {
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, 2)
}
func lstore_3(f *rtdata.Frame) {
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, 3)
}

// TODO <T>a, fa,la,ia... store, store T to array
