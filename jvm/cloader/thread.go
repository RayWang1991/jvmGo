package cloader

import (
	"fmt"
	"jvmGo/jvm/cmn"
	"jvmGo/jvm/instructions"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/utils"
)

func call(m *marea.Method) {
	t := rtdt.NewThread(1024)
	t.PushFrame(rtdt.NewFrame(m, t))
	utils.DLoaderPrintf("[INIT] start %s cls %s\n", m.Name(), m.Class().ClassName())
	//for c := t.CurrentFrame(); c != nil; c = c.GetNext() {
	//	fmt.Println(c.Method().Class().ClassName())
	//}
	var f = t.CurrentFrame()
	for ; f != nil; f = t.CurrentFrame() {
		code := f.ReadU8() // read next opcode
		//debug
		fmt.Printf("[LOAD]pc:%-4d code:%s class:%s method:%s\n",
			f.GetPC()-1, cmn.InstStr(code), f.Method().Class().ClassName(), f.Method().Name())
		fn := instructions.InstFnc(code)
		fn(f)
		//fmt.Printf("Vars:%v\n", t.CurrentFrame().LocalVar)
		//fmt.Printf("OperandStack:%v\n", t.CurrentFrame().OperandStack)
	}
	utils.DLoaderPrintf("[INIT] done %s cls %s\n", m.Name(), m.Class().ClassName())
}
