package instructions

import (
	"jvmGo/ch6/rtdt"
	"jvmGo/ch6/utils"
)

//cmn.OPCODE_getstatic:       getstatic,
//cmn.OPCODE_putstatic:       putstatic,
//cmn.OPCODE_getfield:        getfield,
//cmn.OPCODE_putfield:        popfield,
func getfield(f *rtdt.Frame) {
	// get *field
	inx := f.ReadU16()
	class := f.Method().Class()
	ref := class.ConstantPool().GetFieldRef(inx)
	field := ref.GetField()
	if field.IsStatic() {
		panic(utils.IncompatibleClassChangeError)
	}

	// get and push value
	obj := f.OperandStack.PopRef()
	if obj == nil {
		panic(utils.NullPointerException)
	}
	vid := field.VarId()
	
	f.OperandStack.PushRef(field.)
}
