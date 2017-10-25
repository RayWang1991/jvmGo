package instructions

import "jvmGo/ch6/rtdt"

func branchI16(f *rtdt.Frame, off int16) {
	f.SetPC(f.GetPC() + int32(off) - 3)
}

func branchI32(f *rtdt.Frame, off int32) {
	f.SetPC(f.GetPC() + int32(off) - 5)
}
