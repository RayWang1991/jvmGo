package rtdata

type Frame struct {
	localVar     LocalVars // the local var represents the local variables, the length is given by compiler
	operandStack *OperandStack
	next         *Frame
}
