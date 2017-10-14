package instructions

import "fmt"

type codeType int

const (
	noL     codeType = iota // no extra length
	u8L                     // u8 length
	u16L                    // u16 length
	u32L                    // u32 length
	u8_2L                   // u8*2 length
	tables                  // tableswitch
	lookups                 // lookupswitch
)

var inst2Type = map[uint8]codeType{
	// Constants
	OPCODE_nop:         noL,
	OPCODE_aconst_null: noL,
	OPCODE_iconst_m1:   noL,
	OPCODE_iconst_0:    noL,
	OPCODE_iconst_1:    noL,
	OPCODE_iconst_2:    noL,
	OPCODE_iconst_3:    noL,
	OPCODE_iconst_4:    noL,
	OPCODE_iconst_5:    noL,
	OPCODE_lconst_0:    noL,
	OPCODE_lconst_1:    noL,
	OPCODE_fconst_0:    noL,
	OPCODE_fconst_1:    noL,
	OPCODE_fconst_2:    noL,
	OPCODE_dconst_0:    noL,
	OPCODE_dconst_1:    noL,
	OPCODE_bipush:      u8L,
	OPCODE_sipush:      u16L,
	//OPCODE_ldc:         u8L,
	//OPCODE_ldc_w:       u16L,
	//OPCODE_ldc2_w:      u16L,

	// Loads
	OPCODE_iload:   u8L,
	OPCODE_lload:   u8L,
	OPCODE_fload:   u8L,
	OPCODE_dload:   u8L,
	OPCODE_aload:   u8L,
	OPCODE_iload_0: noL,
	OPCODE_iload_1: noL,
	OPCODE_iload_2: noL,
	OPCODE_iload_3: noL,
	OPCODE_lload_0: noL,
	OPCODE_lload_1: noL,
	OPCODE_lload_2: noL,
	OPCODE_lload_3: noL,
	OPCODE_fload_0: noL,
	OPCODE_fload_1: noL,
	OPCODE_fload_2: noL,
	OPCODE_fload_3: noL,
	OPCODE_dload_0: noL,
	OPCODE_dload_1: noL,
	OPCODE_dload_2: noL,
	OPCODE_dload_3: noL,
	OPCODE_aload_0: noL,
	OPCODE_aload_1: noL,
	OPCODE_aload_2: noL,
	OPCODE_aload_3: noL,
	//OPCODE_iaload:  TODO iaload,
	//OPCODE_laload:  TODO laload,
	//OPCODE_faload:  TODO faload,
	//OPCODE_daload:  TODO daload,
	//OPCODE_aaload:  TODO aaload,
	//OPCODE_baload:  TODO baload,
	//OPCODE_caload:  TODO caload,
	//OPCODE_saload:  TODO saload,

	// Stores
	OPCODE_istore:   u8L,
	OPCODE_lstore:   u8L,
	OPCODE_fstore:   u8L,
	OPCODE_dstore:   u8L,
	OPCODE_astore:   u8L,
	OPCODE_istore_0: noL,
	OPCODE_istore_1: noL,
	OPCODE_istore_2: noL,
	OPCODE_istore_3: noL,
	OPCODE_lstore_0: noL,
	OPCODE_lstore_1: noL,
	OPCODE_lstore_2: noL,
	OPCODE_lstore_3: noL,
	OPCODE_fstore_0: noL,
	OPCODE_fstore_1: noL,
	OPCODE_fstore_2: noL,
	OPCODE_fstore_3: noL,
	OPCODE_dstore_0: noL,
	OPCODE_dstore_1: noL,
	OPCODE_dstore_2: noL,
	OPCODE_dstore_3: noL,
	OPCODE_astore_0: noL,
	OPCODE_astore_1: noL,
	OPCODE_astore_2: noL,
	OPCODE_astore_3: noL,
	//OPCODE_iastore:  iastore,
	//OPCODE_lastore:  lastore,
	//OPCODE_fastore:  fastore,
	//OPCODE_dastore:  dastore,
	//OPCODE_aastore:  aastore,
	//OPCODE_bastore:  bastore,
	//OPCODE_castore:  castore,
	//OPCODE_sastore:  sastore,

	// Stack
	OPCODE_pop:     noL,
	OPCODE_pop2:    noL,
	OPCODE_dup:     noL,
	OPCODE_dup_x1:  noL,
	OPCODE_dup_x2:  noL,
	OPCODE_dup2:    noL,
	OPCODE_dup2_x1: noL,
	OPCODE_dup2_x2: noL,
	OPCODE_swap:    noL,

	//	Math
	OPCODE_iadd:  noL,
	OPCODE_ladd:  noL,
	OPCODE_fadd:  noL,
	OPCODE_dadd:  noL,
	OPCODE_isub:  noL,
	OPCODE_lsub:  noL,
	OPCODE_fsub:  noL,
	OPCODE_dsub:  noL,
	OPCODE_imul:  noL,
	OPCODE_lmul:  noL,
	OPCODE_fmul:  noL,
	OPCODE_dmul:  noL,
	OPCODE_idiv:  noL,
	OPCODE_ldiv:  noL,
	OPCODE_fdiv:  noL,
	OPCODE_ddiv:  noL,
	OPCODE_irem:  noL,
	OPCODE_lrem:  noL,
	OPCODE_frem:  noL,
	OPCODE_drem:  noL,
	OPCODE_ineg:  noL,
	OPCODE_lneg:  noL,
	OPCODE_fneg:  noL,
	OPCODE_dneg:  noL,
	OPCODE_ishl:  noL,
	OPCODE_lshl:  noL,
	OPCODE_ishr:  noL,
	OPCODE_lshr:  noL,
	OPCODE_iushr: noL,
	OPCODE_lushr: noL,
	OPCODE_iand:  noL,
	OPCODE_land:  noL,
	OPCODE_ior:   noL,
	OPCODE_lor:   noL,
	OPCODE_ixor:  noL,
	OPCODE_lxor:  noL,
	OPCODE_iinc:  u8_2L,

	// Conversions
	OPCODE_i2l: noL,
	OPCODE_i2f: noL,
	OPCODE_i2d: noL,
	OPCODE_l2i: noL,
	OPCODE_l2f: noL,
	OPCODE_l2d: noL,
	OPCODE_f2i: noL,
	OPCODE_f2l: noL,
	OPCODE_f2d: noL,
	OPCODE_d2i: noL,
	OPCODE_d2l: noL,
	OPCODE_d2f: noL,
	OPCODE_i2b: noL,
	OPCODE_i2c: noL,
	OPCODE_i2s: noL,

	// Comparisons
	OPCODE_lcmp:      noL,
	OPCODE_fcmpl:     noL,
	OPCODE_fcmpg:     noL,
	OPCODE_dcmpl:     noL,
	OPCODE_dcmpg:     noL,
	OPCODE_ifeq:      u16L,
	OPCODE_ifne:      u16L,
	OPCODE_iflt:      u16L,
	OPCODE_ifge:      u16L,
	OPCODE_ifgt:      u16L,
	OPCODE_ifle:      u16L,
	OPCODE_if_icmpeq: u16L,
	OPCODE_if_icmpne: u16L,
	OPCODE_if_icmplt: u16L,
	OPCODE_if_icmpge: u16L,
	OPCODE_if_icmpgt: u16L,
	OPCODE_if_icmple: u16L,
	OPCODE_if_acmpeq: u16L,
	OPCODE_if_acmpne: u16L,

	// Control
	OPCODE_ggoto:        u16L,
	OPCODE_jsr:          u16L,
	OPCODE_ret:          u8L,
	OPCODE_tableswitch:  tables,
	OPCODE_lookupswitch: lookups,
	//OPCODE_ireturn:      ireturn,
	//OPCODE_lreturn:      lreturn,
	//OPCODE_freturn:      freturn,
	//OPCODE_dreturn:      dreturn,
	//OPCODE_areturn:      areturn,
	//OPCODE_rreturn:      rreturn,

	// References
	//OPCODE_getstatic:       getstatic,
	//OPCODE_putstatic:       putstatic,
	//OPCODE_getfield:        getfield,
	//OPCODE_popfield:        popfield,
	//OPCODE_invokevirtual:   invokevirtual,
	//OPCODE_invokespecial:   invokespecial,
	//OPCODE_invokestatic:    invokestatic,
	//OPCODE_invokeinterface: invokeinterface,
	//OPCODE_invokedynamic:   invokedynamic,
	//OPCODE_new:             new,
	//OPCODE_newarray:        newarray,
	//OPCODE_anewarray:       anewarray,
	//OPCODE_arraylength:     arraylength,
	//OPCODE_athrow:          athrow,
	//OPCODE_checkcast:       checkcast,
	//OPCODE_instanceof:      instanceof,
	//OPCODE_monitorenter:    monitorenter,
	//OPCODE_monitorexit:     monitorexit,

	// Extended
	//OPCODE_wide:           wide,
	//OPCODE_multianewarray: multianewarray,
	OPCODE_ifnull:    u16L,
	OPCODE_ifnonnull: u16L,
	OPCODE_ggoto_w:   u32L,
	OPCODE_jsr_w:     u32L,

	// Reserved is not shown
}

//code reader is the rawCode holder for byte rawCode
type codeReader struct {
	rawCode []byte // raw code
	rawI    int
	code    []byte // real code
}

func NewCodeReader(rawCode []byte) *codeReader {
	return &codeReader{
		rawCode: rawCode,
		code:    make([]byte, 0, len(rawCode)),
	}
}

// getter
func (c *codeReader) Code() []byte {
	return c.code
}

func (c *codeReader) readU8() byte {
	r := c.rawCode[c.rawI]
	c.rawI++
	return r
}

func (c *codeReader) readU16() uint16 {
	high, low := c.rawCode[c.rawI], c.rawCode[c.rawI+1]
	c.rawI += 2
	return uint16(high)<<8 | uint16(low)
}

func (c *codeReader) readI32() int32 {
	b0 := int32(c.rawCode[c.rawI])
	b1 := int32(c.rawCode[c.rawI+1])
	b2 := int32(c.rawCode[c.rawI+2])
	b3 := int32(c.rawCode[c.rawI+3])
	c.rawI += 4
	return b0<<24 | b1<<16 | b2<<8 | b3
}

func (c *codeReader) readI32s(n int) []int32 {
	arr := make([]int32, n)
	for i := 0; i < n; i++ {
		arr[i] = c.readI32()
	}
	return arr
}

func (c *codeReader) skipPadding() {
	var i int
	pc := c.rawI
	for i = pc; i < 4+pc; i++ {
		if i%4 == 0 {
			break
		}
	}
	c.rawI = i
}

func (c *codeReader) ReadCode() {
	n := len(c.rawCode)
	for c.rawI < n {
		code := c.readU8()
		t, ok := inst2Type[code]
		if !ok {
			fmt.Errorf("undefined code %x\n", code) // maybe valid
		}
		c.code = append(c.code, code)
		switch t {
		case noL:
		case u8L:
			c.readU8()
		case u8_2L, u16L:
			c.readU16()
		case u32L:
			c.readI32()
		case tables:
			c.skipPadding()
			c.readI32()         // default
			low := c.readI32()  // low
			high := c.readI32() // high
			c.readI32s(int(high - low + 1))
		case lookups:
			c.skipPadding()
			c.readI32()      // default
			n := c.readI32() // high
			c.readI32s(int(n) * 2)
		}
	}
}
