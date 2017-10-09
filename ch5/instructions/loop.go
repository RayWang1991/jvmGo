package instructions

import (
	"jvmGo/ch5/rtdata"
)

// loop for current frame
func loop(t *rtdata.Thread) {
	f := t.CurrentFrame()
	for {
		code := f.ReadU8() // read next opcode
		fn := InstFnc(code)
		fn(f)
	}
}

func debug(f *rtdata.Frame){
}
