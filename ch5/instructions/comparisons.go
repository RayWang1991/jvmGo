package instructions

import "jvmGo/ch5/rtdata"

func lcmp(f *rtdata.Frame) {
	v2 := f.OperandStack.PopLong()
	v1 := f.OperandStack.PopLong()
	var res int32
	if v1 == v2 {
		res = 0
	} else if v1 < v2 {
		res = -1
	} else {
		res = 1
	}
	f.OperandStack.PushInt(res)
}

func fcmpl(f *rtdata.Frame) {
	v2 := f.OperandStack.PopFloat()
	v1 := f.OperandStack.PopFloat()
	var res int32
	if v1 == v2 {
		res = 0
	} else if v1 > v2 {
		res = 1
	} else {
		// v1 < v2 | NaN
		res = -1
	}
	f.OperandStack.PushInt(res)
}

func fcmpg(f *rtdata.Frame) {
	v2 := f.OperandStack.PopFloat()
	v1 := f.OperandStack.PopFloat()
	var res int32
	if v1 == v2 {
		res = 0
	} else if v1 < v2 {
		res = -1
	} else {
		// v1 > v2 | NaN
		res = 1
	}
	f.OperandStack.PushInt(res)
}

func dcml(f *rtdata.Frame) {
	v2 := f.OperandStack.PopDouble()
	v1 := f.OperandStack.PopDouble()
	var res int32
	if v1 == v2 {
		res = 0
	} else if v1 > v2 {
		res = 1
	} else {
		// v1 < v2 | NaN
		res = -1
	}
	f.OperandStack.PushInt(res)
}

func dcmg(f *rtdata.Frame) {
	v2 := f.OperandStack.PopDouble()
	v1 := f.OperandStack.PopDouble()
	var res int32
	if v1 == v2 {
		res = 0
	} else if v1 < v2 {
		res = -1
	} else {
		// v1 > v2 | NaN
		res = 1
	}
	f.OperandStack.PushInt(res)
}

// if<cond>
func ifeq(f *rtdata.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v == 0 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func ifne(f *rtdata.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v != 0 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func iflt(f *rtdata.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v < 0 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func ifle(f *rtdata.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v <= 0 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func ifgt(f *rtdata.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v > 0 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func ifge(f *rtdata.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v >= 0 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

// if_icmp<cond>
func _icmp(f *rtdata.Frame) (int16, int32, int32) {
	b := f.ReadI16()
	v2 := f.OperandStack.PopInt()
	v1 := f.OperandStack.PopInt()
	return b, v1, v2
}

func if_imcpeq(f *rtdata.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 == v2 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func if_imcpne(f *rtdata.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 != v2 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func if_imcplt(f *rtdata.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 < v2 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func if_imcple(f *rtdata.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 <= v2 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func if_imcpgt(f *rtdata.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 > v2 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func if_imcpge(f *rtdata.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 >= v2 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

// if_acmp
func _acmp(f *rtdata.Frame) (int16, *rtdata.Object, *rtdata.Object) {
	b := f.ReadI16()
	v2 := f.OperandStack.PopRef()
	v1 := f.OperandStack.PopRef()
	return b, v1, v2
}

func if_acmpeq(f *rtdata.Frame) {
	b, v1, v2 := _acmp(f)
	if v1 == v2 {
		f.SetPC(f.GetPC() + int32(b))
	}
}

func if_acmpne(f *rtdata.Frame) {
	b, v1, v2 := _acmp(f)
	if v1 != v2 {
		f.SetPC(f.GetPC() + int32(b))
	}
}
