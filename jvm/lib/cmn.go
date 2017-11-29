package lib

import (
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
)

func callMethod(m *marea.Method, f *rtdt.Frame) {
	t := f.Thread()
	nf := rtdt.NewFrame(m, t)
	t.PushFrame(nf)
	i := m.ArgSlotNum()
	if m.IsStatic() {
		i--
	}
	for ; i >= 0; i-- {
		slot := f.OperandStack.PopSlot()
		nf.LocalVar.SetSlot(slot, uint(i))
	}
}
