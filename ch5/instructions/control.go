package instructions

import "jvmGo/ch5/rtdata"

func ggoto(f *rtdata.Frame) {
	b := f.ReadI16()
	f.SetPC(f.GetPC() + int32(b))
}

func ggoto_w(f *rtdata.Frame) {
	b := f.ReadI32()
	f.SetPC(f.GetPC() + b)
}

// jump subroutine
func jsr(f *rtdata.Frame) {
	addr := f.ReadI16()
	

}
