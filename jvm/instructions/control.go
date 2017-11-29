package instructions

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
)

func ggoto(f *rtdt.Frame) {
	b := f.ReadI16()
	branchI16(f, b)
}

func ggoto_w(f *rtdt.Frame) {
	b := f.ReadI32()
	branchI32(f, b)
}

// jump subroutine
func jsr(f *rtdt.Frame) {
	b := f.ReadI16()
	// get next code's pc index and use it as address
	f.OperandStack.PushInt(f.GetPC())
	branchI16(f, b)
}

func jsr_w(f *rtdt.Frame) {
	b := f.ReadI32()
	// get next code's pc index and use it as address
	f.OperandStack.PushInt(f.GetPC())
	branchI32(f, b)
}

func ret(f *rtdt.Frame) {
	i := f.ReadU8()
	l := f.LocalVar.GetInt(uint(i))
	f.SetPC(l)
}

// skip the padding until the pc is multiple of 4
// returns the pc skipped
func skipPadding(f *rtdt.Frame) {
	var i int32
	pc := f.GetPC()
	for i = pc; i < 4+pc; i++ {
		if i%4 == 0 {
			break
		}
	}
	f.SetPC(i)
}

func tableswitch(f *rtdt.Frame) {
	pc := f.GetPC() - 1
	skipPadding(f)
	def := f.ReadI32()
	low := f.ReadI32()
	high := f.ReadI32()
	addrs := f.ReadI32s(int(high - low + 1))
	index := f.OperandStack.PopInt()
	if index < low || index > high {
		f.SetPC(pc + def)
	} else {
		f.SetPC(pc + addrs[index-low])
	}
}

func lookupswitch(f *rtdt.Frame) {
	pc := f.GetPC() - 1
	skipPadding(f)
	def := f.ReadI32()
	n := f.ReadI32()
	pairs := make([][2]int32, n) // key-offset pairs
	for i := 0; i < int(n); i++ {
		pairs[i][0] = f.ReadI32()
		pairs[i][1] = f.ReadI32()
	}
	key := f.OperandStack.PopInt()
	// binary search
	low, high := pairs[0][0], pairs[n-1][0]
	var mid int32
	for low <= high {
		mid = low + (high-low)/2
		if key == mid {
			// hit
			break
		} else if mid < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low <= high {
		// hit
		f.SetPC(pc + pairs[mid][1])
	} else {
		// not hit
		f.SetPC(pc + def)
	}
}

func re(f *rtdt.Frame) {
	utils.DIstrPrintf("[return] %s %s\n", f.Method().Name(), f.Method().Class().ClassName())
}

func top(t *rtdt.Thread) {
	f := t.CurrentFrame()
	if f == nil {
		return
	}
	utils.DIstrPrintf("[re entrer] %s %s\n", f.Method().Name(), f.Method().Class().ClassName())
}

// return family
// debug
func rreturn(f *rtdt.Frame) {
	re(f)
	t := f.Thread()
	t.PopFrame()
	top(t)
}

func areturn(f *rtdt.Frame) {
	re(f)
	ret := f.OperandStack.PopRef()
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushRef(ret)
	top(t)
}

func ireturn(f *rtdt.Frame) {
	re(f)
	ret := f.OperandStack.PopInt()
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushInt(ret)
	top(t)
}

func lreturn(f *rtdt.Frame) {
	re(f)
	ret := f.OperandStack.PopLong()
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushLong(ret)
	top(t)
}

func freturn(f *rtdt.Frame) {
	re(f)
	ret := f.OperandStack.PopFloat()
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushFloat(ret)
	top(t)
}

func dreturn(f *rtdt.Frame) {
	re(f)
	ret := f.OperandStack.PopDouble()
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushDouble(ret)
	top(t)
}
