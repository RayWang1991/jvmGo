package main

import (
	"fmt"
	"jvmGo/jvm/instructions"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/cmn"
)

var mainThread *rtdt.Thread

func GetMainThread() *rtdt.Thread {
	return mainThread
}

func interpretMain(m *marea.Method, args []string) {
	thread := mainThread
	frame := thread.CurrentFrame()

	// create args array
	loader := m.Class().DefineLoader()
	arrStrClass := loader.Load("[java/lang/String")
	arrArgs := marea.NewArrayA(arrStrClass, int32(len(args)))
	arr := arrArgs.ArrGetRefs()
	for i := range arr {
		arr[i] = marea.GetJavaString(args[i], loader)
	}
	frame.LocalVar[0].Ref = arrArgs

	//thread.PushFrame(frame)
	//defer catchErr(thread)
	loop(thread)
}
func catchErr(t *rtdt.Thread) {
	if r := recover(); r != nil {
		if t.CurrentFrame() != nil {
			fmt.Printf("Vars:%s\n", t.CurrentFrame().LocalVar)
			fmt.Printf("OperandStack:%v\n", t.CurrentFrame().OperandStack)
		} else {
			panic(r)
		}
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
		t.SetPC(f.GetPC()) // back up pc in case of roll back
		code := f.ReadU8() // read next opcode/**/
		fmt.Printf("[MAIN]pc:%-4d code:%s class:%s method:%s\n",
			f.GetPC()-1, cmn.InstStr(code), f.Method().Class().ClassName(), f.Method().Name())
		//fmt.Printf("Vars:%s\n", t.CurrentFrame().LocalVar)
		fn := instructions.InstFnc(code)
		fmt.Printf("Vars:%v\n", t.CurrentFrame().LocalVar)
		fmt.Printf("OperandStack:%v\n", t.CurrentFrame().OperandStack)
		fn(f)
	}
}
