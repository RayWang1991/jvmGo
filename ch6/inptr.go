package main

import (
	"fmt"
	"jvmGo/ch6/cmn"
	"jvmGo/ch6/instructions"
	"jvmGo/ch6/marea"
	"jvmGo/ch6/rtdt"
)

func interpret(m *marea.Method) {
	thread := rtdt.NewThread(1024)
	frame := rtdt.NewFrame(m, thread)
	thread.PushFrame(frame)
	//defer catchErr(thread)
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
	fmt.Println()
	var f = t.CurrentFrame()
	for ; f != nil; f = t.CurrentFrame() {
		if f == nil {
			fmt.Println("exit")
			break
		}
		//fmt.Printf("enter func %s %s\n", f.Method().Name(), f.Method().Desc())
		//fmt.Print(classfile.CodeInst(f.Method().Code()).String())
		code := f.ReadU8() // read next opcode/**/
		fmt.Printf("pc:%-4d code:%s\n", f.GetPC()-1, cmn.InstStr(code))
		fn := instructions.InstFnc(code)
		fn(f)
		//fmt.Printf("Vars:%v\n", t.CurrentFrame().LocalVar)
		//fmt.Printf("OperandStack:%v\n", t.CurrentFrame().OperandStack)
	}
}
