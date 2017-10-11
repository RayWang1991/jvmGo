package instructions

import "jvmGo/ch6/rtdata"

func branchI16(f *rtdata.Frame, off int16) {
	f.SetPC(f.GetPC() + int32(off) - 3)
}

func branchI32(f *rtdata.Frame, off int32) {
	f.SetPC(f.GetPC() + int32(off) - 5)
}
