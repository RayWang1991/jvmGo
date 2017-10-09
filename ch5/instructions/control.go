package instructions

import "jvmGo/ch5/rtdata"

func ggoto(f *rtdata.Frame) {
	b := f.ReadI16()
	branchI16(f, b)
}

func ggoto_w(f *rtdata.Frame) {
	b := f.ReadI32()
	branchI32(f, b)
}

// jump subroutine
func jsr(f *rtdata.Frame) {
	b := f.ReadI16()
	// get next code's pc index and use it as address
	f.OperandStack.PushInt(f.GetPC())
	branchI16(f, b)
}

func ret(f *rtdata.Frame) {
	i := f.ReadU8()
	l := f.LocalVar.GetInt(uint(i))
	f.SetPC(l)
}
