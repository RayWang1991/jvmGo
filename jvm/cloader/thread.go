package cloader

import (
	"fmt"
	"jvmGo/jvm/cmn"
	"jvmGo/jvm/instructions"
	"jvmGo/jvm/rtdt"
)

var loaderThread *rtdt.Thread

func init() {
	loaderThread = rtdt.NewThread(1024)
}

func GetLoaderThread() *rtdt.Thread {
	return loaderThread
}

// loop for current frame
func loop(t *rtdt.Thread) {
	fmt.Println("start loader loop")
	// print all class to be init
	for c := t.CurrentFrame(); c != nil; c = c.GetNext() {
		fmt.Println(c.Method().Class().ClassName())
	}
	var f = t.CurrentFrame()
	for ; f != nil; f = t.CurrentFrame() {
		//fmt.Printf("enter func %s %s\n", f.Method().Name(), f.Method().Desc())
		//fmt.Print(classfile.CodeInst(f.Method().Code()).String())
		code := f.ReadU8() // read next opcode
		fmt.Printf("pc:%-4d code:%s\n", f.GetPC()-1, cmn.InstStr(code))
		fn := instructions.InstFnc(code)
		fn(f)
		//fmt.Printf("Vars:%v\n", t.CurrentFrame().LocalVar)
		//fmt.Printf("OperandStack:%v\n", t.CurrentFrame().OperandStack)
	}
	fmt.Println("init done")
}
