package instructions

import (
	"jvmGo/ch5/rtdata"
)

// loop for current frame
func loop(t *rtdata.Thread) {
	f := t.CurrentFrame()
	for {
		code := f.ReadU8()
		excute(f, code)
	}
}

func excute(f *rtdata.Frame, opcode uint8) {
	//switch opcode {
	//case
	//}
}
