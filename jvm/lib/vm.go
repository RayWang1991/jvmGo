package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_System, utils.METHODNAME_InitializeSystemClass, "()V", initialize)
}

func initialize(frame *rtdt.Frame) {
	panic("")
	loader := frame.Method().Class().DefineLoader()
	sysC := loader.Load(utils.CLASSNAME_System)
	initM := sysC.GetMethodDirect(utils.METHODNAME_InitializeSystemClass, "()V")
	callMethod(initM, frame)
}
