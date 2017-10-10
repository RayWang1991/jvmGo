package main

import (
	"jvmGo/ch5/classfile"
	"jvmGo/ch5/rtdata"
	"jvmGo/ch5/instructions"
	"fmt"
)

func interpret(m *classfile.MethodInfo) {
	codeAttr := m.GetCodeAttr()
	maxStackDep := uint(codeAttr.MaxStack())
	maxLvals := uint(codeAttr.MaxLocals())
	code := codeAttr.Code()
	thread := rtdata.NewThread(1024)
	frame := rtdata.NewFrame(maxLvals, maxStackDep, code)
	thread.PushFrame(frame)
	defer catchErr(thread)
	loop(thread)
}
func catchErr(t *rtdata.Thread) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", t.CurrentFrame().LocalVar)
		fmt.Printf("OperandStack:%v\n", t.CurrentFrame().OperandStack)
	}
}

// loop for current frame
func loop(t *rtdata.Thread) {
	f := t.CurrentFrame()
	for {
		code := f.ReadU8() // read next opcode
		fmt.Printf("pc:%-4d code:%s\n", f.GetPC()-1, instructions.InstStr(code))
		fn := instructions.InstFnc(code)
		fn(f)
	}
}
