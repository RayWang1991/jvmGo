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
	var mid, t, low, high int32
	low, high = 0, n-1
	for low <= high {
		mid = low + (high-low)/2
		t = pairs[mid][0]
		if key == t {
			// hit
			break
		} else if t < key {
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
	utils.DIstrPrintf("[re enter] %s %s\n", f.Method().Name(), f.Method().Class().ClassName())
}

// return family
func rreturn(f *rtdt.Frame) {
	t := f.Thread()
	t.PopFrame()
}

func areturn(f *rtdt.Frame) {
	ret := f.OperandStack.PopRef()
	//debug
	utils.DIstrPrintf("%10s\n", ret)
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushRef(ret)
}

func ireturn(f *rtdt.Frame) {
	ret := f.OperandStack.PopInt()
	//debug
	utils.DIstrPrintf("%10d\n", ret)
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushInt(ret)
}

func lreturn(f *rtdt.Frame) {
	ret := f.OperandStack.PopLong()
	//debug
	utils.DIstrPrintf("%10d\n", ret)
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushLong(ret)
}

func freturn(f *rtdt.Frame) {
	ret := f.OperandStack.PopFloat()
	//debug
	utils.DIstrPrintf("%10f\n", ret)
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushFloat(ret)
}

func dreturn(f *rtdt.Frame) {
	ret := f.OperandStack.PopDouble()
	//debug
	utils.DIstrPrintf("%10f\n", ret)
	t := f.Thread()
	t.PopFrame()
	t.CurrentFrame().OperandStack.PushDouble(ret)
}
