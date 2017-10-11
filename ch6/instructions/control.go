package instructions

import "jvmGo/ch6/rtdata"

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

func jsr_w(f *rtdata.Frame) {
	b := f.ReadI32()
	// get next code's pc index and use it as address
	f.OperandStack.PushInt(f.GetPC())
	branchI32(f, b)
}

func ret(f *rtdata.Frame) {
	i := f.ReadU8()
	l := f.LocalVar.GetInt(uint(i))
	f.SetPC(l)
}

// skip the padding until the pc is multiple of 4
// returns the pc skipped
func skipPadding(f *rtdata.Frame) {
	var i int32
	pc := f.GetPC()
	for i = pc; i < 4+pc; i++ {
		if i%4 == 0 {
			break
		}
	}
	f.SetPC(i)
}

func tableswitch(f *rtdata.Frame) {
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

func lookupswitch(f *rtdata.Frame) {
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
