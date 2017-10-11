package instructions

import "jvmGo/ch6/rtdata"

// duplicate the top slot of the operand stack
func dup(f *rtdata.Frame) {
	s := f.OperandStack.Top()
	ns := s.Copy()
	f.OperandStack.PushSlot(ns)
}

func dup2(f *rtdata.Frame) {
	v1 := f.OperandStack.Top()
	v2 := f.OperandStack.GetSlot(1)
	v1d := v1.Copy()
	v2d := v2.Copy()
	f.OperandStack.PushSlot(v2d)
	f.OperandStack.PushSlot(v1d)
}

// duplicate the top operand stack value and insert two values down
// TODO simplify
func dup_x1(f *rtdata.Frame) {
	v1 := f.OperandStack.PopSlot()
	v2 := f.OperandStack.PopSlot()
	v1d := v1.Copy()
	f.OperandStack.PushSlot(v1d)
	f.OperandStack.PushSlot(v2)
	f.OperandStack.PushSlot(v1)
}

func dup_x2(f *rtdata.Frame) {
	v1 := f.OperandStack.PopSlot()
	v2 := f.OperandStack.PopSlot()
	v3 := f.OperandStack.PopSlot()
	v1d := v1.Copy()
	f.OperandStack.PushSlot(v1d)
	f.OperandStack.PushSlot(v3)
	f.OperandStack.PushSlot(v2)
	f.OperandStack.PushSlot(v1)
}

func dup2_x1(f *rtdata.Frame) {
	v1 := f.OperandStack.PopSlot()
	v2 := f.OperandStack.PopSlot()
	v3 := f.OperandStack.PopSlot()
	v1d := v1.Copy()
	v2d := v2.Copy()
	f.OperandStack.PushSlot(v2d)
	f.OperandStack.PushSlot(v1d)
	f.OperandStack.PushSlot(v3)
	f.OperandStack.PushSlot(v2)
	f.OperandStack.PushSlot(v1)
}

func dup2_x2(f *rtdata.Frame) {
	v1 := f.OperandStack.PopSlot()
	v2 := f.OperandStack.PopSlot()
	v3 := f.OperandStack.PopSlot()
	v4 := f.OperandStack.PopSlot()
	v1d := v1.Copy()
	v2d := v2.Copy()
	f.OperandStack.PushSlot(v2d)
	f.OperandStack.PushSlot(v1d)
	f.OperandStack.PushSlot(v4)
	f.OperandStack.PushSlot(v3)
	f.OperandStack.PushSlot(v2)
	f.OperandStack.PushSlot(v1)
}

func pop(f *rtdata.Frame) {
	f.OperandStack.PopSlot()
}

func pop2(f *rtdata.Frame) {
	f.OperandStack.PopSlot()
	f.OperandStack.PopSlot()
}

func swap(f *rtdata.Frame) {
	v1 := f.OperandStack.PopSlot()
	v2 := f.OperandStack.PopSlot()
	f.OperandStack.PushSlot(v1)
	f.OperandStack.PushSlot(v2)
}
