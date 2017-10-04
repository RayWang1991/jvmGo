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
func lload_0(f *rtdata.Frame, i uint8) int64 {
	return f.LocalVar.GetLong(uint(0))
}
func lload_1(f *rtdata.Frame, i uint8) int64 {
	return f.LocalVar.GetLong(uint(1))
}
func lload_2(f *rtdata.Frame, i uint8) int64 {
	return f.LocalVar.GetLong(uint(2))
}
func lload_3(f *rtdata.Frame, i uint8) int64 {
	return f.LocalVar.GetLong(uint(3))
}

// fload
func fload(f *rtdata.Frame, i uint8) float32 {
	return f.LocalVar.GetFloat(uint(i))
}
func fload_0(f *rtdata.Frame, i uint8) float32 {
	return f.LocalVar.GetFloat(uint(0))
}
func fload_1(f *rtdata.Frame, i uint8) float32 {
	return f.LocalVar.GetFloat(uint(1))
}
func fload_2(f *rtdata.Frame, i uint8) float32 {
	return f.LocalVar.GetFloat(uint(2))
}
func fload_3(f *rtdata.Frame, i uint8) float32 {
	return f.LocalVar.GetFloat(uint(3))
}

// dload
func dload(f *rtdata.Frame, i uint8) float64 {
	return f.LocalVar.GetDouble(uint(i))
}
func dload_0(f *rtdata.Frame, i uint8) float64 {
	return f.LocalVar.GetDouble(uint(0))
}
func dload_1(f *rtdata.Frame, i uint8) float64 {
	return f.LocalVar.GetDouble(uint(1))
}
func dload_2(f *rtdata.Frame, i uint8) float64 {
	return f.LocalVar.GetDouble(uint(2))
}
func dload_3(f *rtdata.Frame, i uint8) float64 {
	return f.LocalVar.GetDouble(uint(3))
}

func aload(f *rtdata.Frame, i uint8) *rtdata.Object {
	return f.LocalVar.GetRef(uint(i))
}
func aload_0(f *rtdata.Frame, i uint8) *rtdata.Object {
	return f.LocalVar.GetRef(uint(0))
}
func aload_1(f *rtdata.Frame, i uint8) *rtdata.Object {
	return f.LocalVar.GetRef(uint(1))
}
func aload_2(f *rtdata.Frame, i uint8) *rtdata.Object {
	return f.LocalVar.GetRef(uint(2))
}
func aload_3(f *rtdata.Frame, i uint8) *rtdata.Object {
	return f.LocalVar.GetRef(uint(3))
}


// Stores
