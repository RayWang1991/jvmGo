package instructions

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"math"
)

// add
func iadd(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	f.OperandStack.PushInt(v1 + v2)
}

func ladd(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopLong()
	f.OperandStack.PushLong(v1 + v2)
}

func fadd(f *rtdt.Frame) {
	v1 := f.OperandStack.PopFloat()
	v2 := f.OperandStack.PopFloat()
	f.OperandStack.PushFloat(v1 + v2)
}

func dadd(f *rtdt.Frame) {
	v1 := f.OperandStack.PopDouble()
	v2 := f.OperandStack.PopDouble()
	f.OperandStack.PushDouble(v1 + v2)
}

// sub
func isub(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	f.OperandStack.PushInt(v2 - v1)
}

func lsub(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopLong()
	f.OperandStack.PushLong(v2 - v1)
}

func fsub(f *rtdt.Frame) {
	v1 := f.OperandStack.PopFloat()
	v2 := f.OperandStack.PopFloat()
	f.OperandStack.PushFloat(v2 - v1)
}

func dsub(f *rtdt.Frame) {
	v1 := f.OperandStack.PopDouble()
	v2 := f.OperandStack.PopDouble()
	f.OperandStack.PushDouble(v2 - v1)
}

// mul
func imul(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	f.OperandStack.PushInt(v2 * v1)
}

func lmul(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopLong()
	f.OperandStack.PushLong(v2 * v1)
}

func fmul(f *rtdt.Frame) {
	v1 := f.OperandStack.PopFloat()
	v2 := f.OperandStack.PopFloat()
	f.OperandStack.PushFloat(v2 * v1)
}

func dmul(f *rtdt.Frame) {
	v1 := f.OperandStack.PopDouble()
	v2 := f.OperandStack.PopDouble()
	f.OperandStack.PushDouble(v2 * v1)
}

// div
// TODO, NAN, Infinite
func idiv(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	if v1 == 0 {
		panic(utils.DivideByZero)
	}
	f.OperandStack.PushInt(v2 / v1)
}

func ldiv(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopLong()
	if v1 == 0 {
		panic(utils.DivideByZero)
	}
	f.OperandStack.PushLong(v2 / v1)
}

func fdiv(f *rtdt.Frame) {
	v1 := f.OperandStack.PopFloat()
	v2 := f.OperandStack.PopFloat()
	if v1 == 0 {
		panic(utils.DivideByZero)
	}
	f.OperandStack.PushFloat(v2 / v1)
}

func ddiv(f *rtdt.Frame) {
	v1 := f.OperandStack.PopDouble()
	v2 := f.OperandStack.PopDouble()
	if v1 == 0 {
		panic(utils.DivideByZero)
	}
	f.OperandStack.PushDouble(v2 / v1)
}

// rem
// TODO NAN, inf
func irem(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	if v1 == 0 {
		panic(utils.DivideByZero)
	}
	f.OperandStack.PushInt(v2 % v1)
}

func lrem(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopLong()
	if v1 == 0 {
		panic(utils.DivideByZero)
	}
	f.OperandStack.PushLong(v2 % v1)
}

func frem(f *rtdt.Frame) {
	v1 := f.OperandStack.PopFloat()
	v2 := f.OperandStack.PopFloat()
	if v1 == 0 {
		panic(utils.DivideByZero)
	}
	f.OperandStack.PushFloat(float32(math.Mod(float64(v2), float64(v1))))
}

func drem(f *rtdt.Frame) {
	v1 := f.OperandStack.PopDouble()
	v2 := f.OperandStack.PopDouble()
	if v1 == 0 {
		panic(utils.DivideByZero)
	}
	f.OperandStack.PushDouble(math.Mod(v2, v1))
}

// neg
func ineg(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	f.OperandStack.PushInt(-v1)
}

func lneg(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	f.OperandStack.PushLong(-v1)
}

func fneg(f *rtdt.Frame) {
	v1 := f.OperandStack.PopFloat()
	f.OperandStack.PushFloat(-v1)
}

func dneg(f *rtdt.Frame) {
	v1 := f.OperandStack.PopDouble()
	f.OperandStack.PushDouble(-v1)
}

// shl
func ishl(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	res := v1 << uint(v2&0x1F)
	f.OperandStack.PushInt(res)
}

func lshl(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopInt()
	res := v1 << uint(v2&0x3F)
	f.OperandStack.PushLong(res)
}

// shr
func ishr(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	res := v1 >> uint(v2&0x1F)
	f.OperandStack.PushInt(res)
}

func lshr(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopInt()
	res := v1 >> uint(v2&0x3F)
	f.OperandStack.PushLong(res)
}

// ushr
func iushr(f *rtdt.Frame) {
	v1 := uint32(f.OperandStack.PopInt())
	v2 := f.OperandStack.PopInt()
	res := v1 >> uint(v2&0x1F)
	f.OperandStack.PushInt(int32(res))
}

func lushr(f *rtdt.Frame) {
	v1 := uint64(f.OperandStack.PopLong())
	v2 := f.OperandStack.PopInt()
	res := v1 >> uint(v2&0x3F)
	f.OperandStack.PushLong(int64(res))
}

// and
func iand(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	f.OperandStack.PushInt(v1 & v2)
}

func land(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopLong()
	f.OperandStack.PushLong(v1 & v2)
}

// or
func ior(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	f.OperandStack.PushInt(v1 | v2)
}

func lor(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopLong()
	f.OperandStack.PushLong(v1 | v2)
}

// xor
func ixor(f *rtdt.Frame) {
	v1 := f.OperandStack.PopInt()
	v2 := f.OperandStack.PopInt()
	f.OperandStack.PushInt(v1 ^ v2)
}

func lxor(f *rtdt.Frame) {
	v1 := f.OperandStack.PopLong()
	v2 := f.OperandStack.PopLong()
	f.OperandStack.PushLong(v1 ^ v2)
}

// inc
func iinc(f *rtdt.Frame) {
	index := f.ReadU8()
	c := int8(f.ReadU8())
	v := f.LocalVar.GetInt(uint(index))
	f.LocalVar.SetInt(v+int32(c), uint(index))
}
