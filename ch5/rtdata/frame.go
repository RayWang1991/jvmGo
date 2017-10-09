package rtdata

type Frame struct {
	LocalVar     LocalVars // the local var represents the local variables, the length is given by compiler
	OperandStack *OperandStack
	Next         *Frame
	// code
	pc   int32 // for code read and execute
	code []byte
}

// code & pc
func (c *Frame) GetPC() int32 {
	return c.pc
}

func (c *Frame) SetPC(pc int32) {
	c.pc = pc
}

func (c *Frame) ReadU8() byte {
	r := c.code[c.pc]
	c.pc++
	return r
}

// read two code(u8) as u16
func (c *Frame) ReadU16() uint16 {
	high, low := c.code[c.pc], c.code[c.pc+1]
	c.pc += 2
	return uint16(high)<<8 | uint16(low)
}

func (c *Frame) ReadI16() int16 {
	return int16(c.ReadU16())
}

func (c *Frame) ReadI32() int32 {
	b0 := int32(c.code[c.pc])
	b1 := int32(c.code[c.pc+1])
	b2 := int32(c.code[c.pc+2])
	b3 := int32(c.code[c.pc+3])
	c.pc += 4
	return b0<<24 | b1<<16 | b2<<8 | b3
}
