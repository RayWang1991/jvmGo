package instructions

// code reader is the code holder for byte code
type codeReader struct {
	code []byte
	pc   int
}

func (c *codeReader) ReadU8() byte {
	r := c.code[c.pc]
	c.pc++
	return r
}

func (c *codeReader) ReadU16() uint16 {
	high, low := c.code[c.pc], c.code[c.pc+1]
	c.pc += 2
	return uint16(high)<<8 | uint16(low)
}

func (c *codeReader) ReadI32() int32 {
	b0 := int32(c.code[c.pc])
	b1 := int32(c.code[c.pc+1])
	b2 := int32(c.code[c.pc+2])
	b3 := int32(c.code[c.pc+3])
	c.pc += 4
	return b0<<24 | b1<<16 | b2<<8 | b3
}
