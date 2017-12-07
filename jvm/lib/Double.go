package lib

import "jvmGo/jvm/rtdt"
import "math"

const jlDouble = "java/lang/Double"

func init() {
	register(jlDouble, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	register(jlDouble, "longBitsToDouble", "(J)D", longBitsToDouble)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *rtdt.Frame) {
	value := frame.LocalVar.GetDouble(0)
	bits := math.Float64bits(value) // todo
	frame.OperandStack.PushLong(int64(bits))
}

// public static native double longBitsToDouble(long bits);
// (J)D
func longBitsToDouble(frame *rtdt.Frame) {
	bits := frame.LocalVar.GetLong(0)
	value := math.Float64frombits(uint64(bits)) // todo
	frame.OperandStack.PushDouble(value)
}
