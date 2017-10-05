package instructions

import "jvmGo/ch5/rtdata"

// Loads
// iload
func iload(f *rtdata.Frame, i uint8) int32 {
	return f.LocalVar.GetInt(uint(i))
}
func iload_0(f *rtdata.Frame) int32 {
	return f.LocalVar.GetInt(0)
}
func iload_1(f *rtdata.Frame) int32 {
	return f.LocalVar.GetInt(1)
}
func iload_2(f *rtdata.Frame) int32 {
	return f.LocalVar.GetInt(2)
}
func iload_3(f *rtdata.Frame) int32 {
	return f.LocalVar.GetInt(3)
}

//lload
func lload(f *rtdata.Frame, i uint8) int64 {
	return f.LocalVar.GetLong(uint(i))
}
func lload_0(f *rtdata.Frame) int64 {
	return f.LocalVar.GetLong(uint(0))
}
func lload_1(f *rtdata.Frame) int64 {
	return f.LocalVar.GetLong(uint(1))
}
func lload_2(f *rtdata.Frame) int64 {
	return f.LocalVar.GetLong(uint(2))
}
func lload_3(f *rtdata.Frame) int64 {
	return f.LocalVar.GetLong(uint(3))
}

// fload
func fload(f *rtdata.Frame, i uint8) float32 {
	return f.LocalVar.GetFloat(uint(i))
}
func fload_0(f *rtdata.Frame) float32 {
	return f.LocalVar.GetFloat(uint(0))
}
func fload_1(f *rtdata.Frame) float32 {
	return f.LocalVar.GetFloat(uint(1))
}
func fload_2(f *rtdata.Frame) float32 {
	return f.LocalVar.GetFloat(uint(2))
}
func fload_3(f *rtdata.Frame) float32 {
	return f.LocalVar.GetFloat(uint(3))
}

// dload
func dload(f *rtdata.Frame, i uint8) float64 {
	return f.LocalVar.GetDouble(uint(i))
}
func dload_0(f *rtdata.Frame) float64 {
	return f.LocalVar.GetDouble(uint(0))
}
func dload_1(f *rtdata.Frame) float64 {
	return f.LocalVar.GetDouble(uint(1))
}
func dload_2(f *rtdata.Frame) float64 {
	return f.LocalVar.GetDouble(uint(2))
}
func dload_3(f *rtdata.Frame) float64 {
	return f.LocalVar.GetDouble(uint(3))
}

func aload(f *rtdata.Frame, i uint8) *rtdata.Object {
	return f.LocalVar.GetRef(uint(i))
}
func aload_0(f *rtdata.Frame) *rtdata.Object {
	return f.LocalVar.GetRef(uint(0))
}
func aload_1(f *rtdata.Frame) *rtdata.Object {
	return f.LocalVar.GetRef(uint(1))
}
func aload_2(f *rtdata.Frame) *rtdata.Object {
	return f.LocalVar.GetRef(uint(2))
}
func aload_3(f *rtdata.Frame) *rtdata.Object {
	return f.LocalVar.GetRef(uint(3))
}

// TODO load int,float,refs... from array

// Stores
func istore(f *rtdata.Frame, i uint8) {
	r := f.OperandStack.PopInt()
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

func dstore(f *rtdata.Frame, i uint8) {
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

func fstore(f *rtdata.Frame, i uint8) {
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

func astore(f *rtdata.Frame, i uint8) {
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

func lstore(f *rtdata.Frame, i uint8) {
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