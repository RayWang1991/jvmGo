package instructions

import "jvmGo/ch6/rtdata"

// int to long
func i2l(f *rtdata.Frame) {
	v := f.OperandStack.PopInt()
	f.OperandStack.PushLong(int64(v))
}

// int to float
func i2f(f *rtdata.Frame) {
	v := f.OperandStack.PopInt()
	f.OperandStack.PushFloat(float32(v))
}

// int to double
func i2d(f *rtdata.Frame) {
	v := f.OperandStack.PopInt()
	f.OperandStack.PushDouble(float64(v))
}

// int to byte
func i2b(f *rtdata.Frame) {
	v := f.OperandStack.PopInt()
	f.OperandStack.PushInt(int32(int8(v)))
}

// int to char
func i2c(f *rtdata.Frame) {
	v := f.OperandStack.PopInt()
	f.OperandStack.PushInt(int32(uint16(v)))
}

// int to short
func i2s(f *rtdata.Frame) {
	v := f.OperandStack.PopInt()
	f.OperandStack.PushInt(int32(int16(v)))
}

// long to int
func l2i(f *rtdata.Frame) {
	v := f.OperandStack.PopLong()
	f.OperandStack.PushInt(int32(v))
}

// long to float
func l2f(f *rtdata.Frame) {
	v := f.OperandStack.PopLong()
	f.OperandStack.PushFloat(float32(v))
}

// long to double
func l2d(f *rtdata.Frame) {
	v := f.OperandStack.PopLong()
	f.OperandStack.PushDouble(float64(v))
}

// float to int
func f2i(f *rtdata.Frame) {
	v := f.OperandStack.PopFloat()
	f.OperandStack.PushInt(int32(v))
}

// float to long
func f2l(f *rtdata.Frame) {
	v := f.OperandStack.PopFloat()
	f.OperandStack.PushLong(int64(v))
}

// float to double
func f2d(f *rtdata.Frame) {
	v := f.OperandStack.PopFloat()
	f.OperandStack.PushDouble(float64(v))
}

// double to int
func d2i(f *rtdata.Frame) {
	v := f.OperandStack.PopDouble()
	f.OperandStack.PushInt(int32(v))
}

// double to long
func d2l(f *rtdata.Frame) {
	v := f.OperandStack.PopDouble()
	f.OperandStack.PushLong(int64(v))
}

// double to float
func d2f(f *rtdata.Frame) {
	v := f.OperandStack.PopDouble()
	f.OperandStack.PushFloat(float32(v))
}
