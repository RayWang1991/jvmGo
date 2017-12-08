package lib

import (
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/cmn"
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

func dummyFrame(op *rtdt.OperandStack, t *rtdt.Thread) *rtdt.Frame {
	dummyClass := marea.HackClass("<dummy>")
	m := marea.HackMethod(
		dummyClass,
		cmn.ACC_STATIC,
		"<dummy>",
		"()V",
		[]byte{cmn.OPCODE_rreturn})
	frame := rtdt.HackFrame(m, t, op, nil)
	return frame
}
