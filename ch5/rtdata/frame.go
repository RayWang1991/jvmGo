package rtdata

type Frame struct {
	LocalVar     LocalVars // the local var represents the local variables, the length is given by compiler
	OperandStack *OperandStack
	Next         *Frame
}
