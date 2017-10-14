package main

import (
	"fmt"
	"jvmGo/ch6/classfile"
	"jvmGo/ch6/instructions"
	"jvmGo/ch6/rtdt"
	"jvmGo/ch6/cmn"
)

func interpret(m *classfile.MethodInfo) {
	codeAttr := m.GetCodeAttr()
	maxStackDep := uint(codeAttr.MaxStack())
	maxLvals := uint(codeAttr.MaxLocals())
	code := codeAttr.Code()
	thread := rtdt.NewThread(1024)
	frame := rtdt.NewFrame(maxLvals, maxStackDep, code)
	thread.PushFrame(frame)
	defer catchErr(thread)
	loop(thread)
}
func catchErr(t *rtdt.Thread) {
	if r := recover(); r != nil {
		fmt.Printf("Vars:%v\n", t.CurrentFrame().LocalVar)
		fmt.Printf("OperandStack:%v\n", t.CurrentFrame().OperandStack)
	}
}

// loop for current frame
func loop(t *rtdt.Thread) {
	f := t.CurrentFrame()
	for {
		code := f.ReadU8() // read next opcode
		fmt.Printf("pc:%-4d code:%s\n", f.GetPC()-1, cmn.InstStr(code))
		fn := instructions.InstFnc(code)
		fn(f)
	}
}
