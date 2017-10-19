package main

import (
	"fmt"
	"jvmGo/ch6/instructions"
	"jvmGo/ch6/rtdt"
	"jvmGo/ch6/cmn"
	"jvmGo/ch6/marea"
	"jvmGo/ch6/classfile"
)

func interpret(m *marea.Method) {
	thread := rtdt.NewThread(1024)
	frame := rtdt.NewFrame(m)
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
	fmt.Printf("enter func %s %s\n", f.Method().Name(), f.Method().Desc())
	fmt.Print(classfile.CodeInst(f.Method().Code()).String())
	for {
		code := f.ReadU8() // read next opcode
		fmt.Printf("pc:%-4d code:%s\n", f.GetPC()-1, cmn.InstStr(code))
		fn := instructions.InstFnc(code)
		fn(f)
	}
}
