package instructions

import (
	"jvmGo/ch6/marea"
	"jvmGo/ch6/rtdt"
)

func lcmp(f *rtdt.Frame) {
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

func fcmpl(f *rtdt.Frame) {
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

func fcmpg(f *rtdt.Frame) {
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

func dcmpl(f *rtdt.Frame) {
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

func dcmpg(f *rtdt.Frame) {
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
func ifeq(f *rtdt.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v == 0 {
		branchI16(f, b)
	}
}

func ifne(f *rtdt.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v != 0 {
		branchI16(f, b)
	}
}

func iflt(f *rtdt.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v < 0 {
		branchI16(f, b)
	}
}

func ifle(f *rtdt.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v <= 0 {
		branchI16(f, b)
	}
}

func ifgt(f *rtdt.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v > 0 {
		branchI16(f, b)
	}
}

func ifge(f *rtdt.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopInt()
	if v >= 0 {
		branchI16(f, b)
	}
}

// if_icmp<cond>
func _icmp(f *rtdt.Frame) (int16, int32, int32) {
	b := f.ReadI16()
	v2 := f.OperandStack.PopInt()
	v1 := f.OperandStack.PopInt()
	return b, v1, v2
}

func if_icmpeq(f *rtdt.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 == v2 {
		branchI16(f, b)
	}
}

func if_icmpne(f *rtdt.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 != v2 {
		branchI16(f, b)
	}
}

func if_icmplt(f *rtdt.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 < v2 {
		branchI16(f, b)
	}
}

func if_icmple(f *rtdt.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 <= v2 {
		branchI16(f, b)
	}
}

func if_icmpgt(f *rtdt.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 > v2 {
		branchI16(f, b)
	}
}

func if_icmpge(f *rtdt.Frame) {
	b, v1, v2 := _icmp(f)
	if v1 >= v2 {
		branchI16(f, b)
	}
}

// if_acmp
func _acmp(f *rtdt.Frame) (int16, *marea.Object, *marea.Object) {
	b := f.ReadI16()
	v2 := f.OperandStack.PopRef()
	v1 := f.OperandStack.PopRef()
	return b, v1, v2
}

func if_acmpeq(f *rtdt.Frame) {
	b, v1, v2 := _acmp(f)
	if v1 == v2 {
		branchI16(f, b)
	}
}

func if_acmpne(f *rtdt.Frame) {
	b, v1, v2 := _acmp(f)
	if v1 != v2 {
		branchI16(f, b)
	}
}

// extend for null
func ifnull(f *rtdt.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopRef()
	if v == nil {
		branchI16(f, b)
	}
}

func ifnonnull(f *rtdt.Frame) {
	b := f.ReadI16()
	v := f.OperandStack.PopRef()
	if v != nil {
		branchI16(f, b)
	}
}
