package lib

import (
	"math"
	"jvmGo/jvm/rtdt"
)

//TODO
const jlFloat = "java/lang/Float"

func init() {
	register(jlFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	register(jlFloat, "intBitsToFloat", "(I)F", intBitsToFloat)
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *rtdt.Frame) {
	value := frame.LocalVar.GetFloat(0)
	bits := math.Float32bits(value) // todo
	frame.OperandStack.PushInt(int32(bits))
}

// public static native float intBitsToFloat(int bits);
// (I)F
func intBitsToFloat(frame *rtdt.Frame) {
	bits := frame.LocalVar.GetInt(0)
	value := math.Float32frombits(uint32(bits)) // todo
	frame.OperandStack.PushFloat(value)
}
