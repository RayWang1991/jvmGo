package instructions

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/marea"
	"fmt"
)

func nop(f *rtdt.Frame) {
	// do nothing
}

// ref const
// push null to operand stack
func aconst_null(f *rtdt.Frame) {
	f.OperandStack.PushRef(nil) // TODO
}

// int const
// push -1 to operand stack
func iconst_m1(f *rtdt.Frame) {
	f.OperandStack.PushInt(-1)
}

// push 0 to operand stack
func iconst_0(f *rtdt.Frame) {
	f.OperandStack.PushInt(0)
}

// push 1 to operand stack
func iconst_1(f *rtdt.Frame) {
	f.OperandStack.PushInt(1)
}

// push 2 to operand stack
func iconst_2(f *rtdt.Frame) {
	f.OperandStack.PushInt(2)
}

// push 3 to operand stack
func iconst_3(f *rtdt.Frame) {
	f.OperandStack.PushInt(3)
}

// push 4 to operand stack
func iconst_4(f *rtdt.Frame) {
	f.OperandStack.PushInt(4)
}

// push 5 to operand stack
func iconst_5(f *rtdt.Frame) {
	f.OperandStack.PushInt(5)
}

// long const
// push 0 to operand stack
func lconst_0(f *rtdt.Frame) {
	f.OperandStack.PushLong(0)
}

// push 1 to operand stack
func lconst_1(f *rtdt.Frame) {
	f.OperandStack.PushLong(1)
}

// float const
// push 0.0f to operand stack
func fconst_0(f *rtdt.Frame) {
	f.OperandStack.PushFloat(0)
}

// push 1.0f to operand stack
func fconst_1(f *rtdt.Frame) {
	f.OperandStack.PushFloat(1)
}

// push 2.0f to operand stack
func fconst_2(f *rtdt.Frame) {
	f.OperandStack.PushFloat(2)
}

// double const
// push 0.0d to operand stack
func dconst_0(f *rtdt.Frame) {
	f.OperandStack.PushDouble(0)
}

// push 1.0d to operand stack
func dconst_1(f *rtdt.Frame) {
	f.OperandStack.PushDouble(1)
}

// push int8(extended to int32) to operand stack
func bipush(f *rtdt.Frame) {
	b := f.ReadU8()
	f.OperandStack.PushInt(int32(b))
}

// push int16(extended to int32) to operand stack,
// call codeReader.Read16I()method for the s parameter
func sipush(f *rtdt.Frame) {
	s := f.ReadU16()
	f.OperandStack.PushInt(int32(s))
}

// TODO ldc, ldcw, lcd2_w, load constants from constant pool
func ldc_base(f *rtdt.Frame, indx uint16) {
	val := f.Method().Class().ConstantPool().GetConstant(indx)

	switch val := val.(type) {
	case int32:
		f.OperandStack.PushInt(val)
	case float32:
		f.OperandStack.PushFloat(val)
	case string:
		// TODO
		panic("string")
	case marea.MethodType:
		// TODO
		panic("method type")
	case *marea.MethodHandle:
		// TODO
		panic("method handle")
	default:
		panic(fmt.Errorf("unsupported type %T", val))
	}
}

func ldc(f *rtdt.Frame) {
	index := uint16(f.ReadU8())
	ldc_base(f, index)
}

func ldc_w(f *rtdt.Frame) {
	index := f.ReadU16()
	ldc_base(f, index)
}

func ldc2_w(f *rtdt.Frame) {
	index := f.ReadU16()
	val := f.Method().Class().ConstantPool().GetConstant(index)

	switch val := val.(type) {
	case int64:
		f.OperandStack.PushLong(val)
	case float64:
		f.OperandStack.PushDouble(val)
	default:
		panic(fmt.Errorf("unsupported type %T", val))
	}
}
