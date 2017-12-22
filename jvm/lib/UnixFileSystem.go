package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/marea"
	"os"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_UnixFileSystem, "getBooleanAttributes0", "(Ljava/io/File;)I", getBooleanAttributes0)
}

// public native int getBooleanAttributes(File f);
// (Ljava/io/File;)I
func getBooleanAttributes0(f *rtdt.Frame) {
	file := f.LocalVar.GetRef(1)
	jPathF := file.Class().InstField("path")
	jPath := file.GetRef(jPathF.VarIdx())
	path := marea.GetGoString(jPath)

	var ba = int32(0)
	info, err := os.Stat(path)
	if err == nil { // todo
		ba |= 1
	}
	if info != nil && info.IsDir() {
		ba |= 4
	}
	f.OperandStack.PushInt(ba)
}
