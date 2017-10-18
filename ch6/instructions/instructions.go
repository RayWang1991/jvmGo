package instructions

import (
	"jvmGo/ch6/rtdt"
	"jvmGo/ch6/cmn"
)

var code2FuncMap = map[uint8]func(*rtdt.Frame){
	// Constants
	cmn.OPCODE_nop:         nop,
	cmn.OPCODE_aconst_null: aconst_null,
	cmn.OPCODE_iconst_m1:   iconst_m1,
	cmn.OPCODE_iconst_0:    iconst_0,
	cmn.OPCODE_iconst_1:    iconst_1,
	cmn.OPCODE_iconst_2:    iconst_2,
	cmn.OPCODE_iconst_3:    iconst_3,
	cmn.OPCODE_iconst_4:    iconst_4,
	cmn.OPCODE_iconst_5:    iconst_5,
	cmn.OPCODE_lconst_0:    lconst_0,
	cmn.OPCODE_lconst_1:    lconst_1,
	cmn.OPCODE_fconst_0:    fconst_0,
	cmn.OPCODE_fconst_1:    fconst_1,
	cmn.OPCODE_fconst_2:    fconst_2,
	cmn.OPCODE_dconst_0:    dconst_0,
	cmn.OPCODE_dconst_1:    dconst_1,
	cmn.OPCODE_bipush:      bipush,
	cmn.OPCODE_sipush:      sipush,
	//cmn.OPCODE_ldc:         ldc,
	//cmn.OPCODE_ldc_w:       ldc_w,
	//cmn.OPCODE_ldc2_w:      ldc2_w,

	// Loads
	cmn.OPCODE_iload:   iload,
	cmn.OPCODE_lload:   lload,
	cmn.OPCODE_fload:   fload,
	cmn.OPCODE_dload:   dload,
	cmn.OPCODE_aload:   aload,
	cmn.OPCODE_iload_0: iload_0,
	cmn.OPCODE_iload_1: iload_1,
	cmn.OPCODE_iload_2: iload_2,
	cmn.OPCODE_iload_3: iload_3,
	cmn.OPCODE_lload_0: lload_0,
	cmn.OPCODE_lload_1: lload_1,
	cmn.OPCODE_lload_2: lload_2,
	cmn.OPCODE_lload_3: lload_3,
	cmn.OPCODE_fload_0: fload_0,
	cmn.OPCODE_fload_1: fload_1,
	cmn.OPCODE_fload_2: fload_2,
	cmn.OPCODE_fload_3: fload_3,
	cmn.OPCODE_dload_0: dload_0,
	cmn.OPCODE_dload_1: dload_1,
	cmn.OPCODE_dload_2: dload_2,
	cmn.OPCODE_dload_3: dload_3,
	cmn.OPCODE_aload_0: aload_0,
	cmn.OPCODE_aload_1: aload_1,
	cmn.OPCODE_aload_2: aload_2,
	cmn.OPCODE_aload_3: aload_3,
	//cmn.OPCODE_iaload:  iaload,
	//cmn.OPCODE_laload:  laload,
	//cmn.OPCODE_faload:  faload,
	//cmn.OPCODE_daload:  daload,
	//cmn.OPCODE_aaload:  aaload,
	//cmn.OPCODE_baload:  baload,
	//cmn.OPCODE_caload:  caload,
	//cmn.OPCODE_saload:  saload,

	// Stores
	cmn.OPCODE_istore:   istore,
	cmn.OPCODE_lstore:   lstore,
	cmn.OPCODE_fstore:   fstore,
	cmn.OPCODE_dstore:   dstore,
	cmn.OPCODE_astore:   astore,
	cmn.OPCODE_istore_0: istore_0,
	cmn.OPCODE_istore_1: istore_1,
	cmn.OPCODE_istore_2: istore_2,
	cmn.OPCODE_istore_3: istore_3,
	cmn.OPCODE_lstore_0: lstore_0,
	cmn.OPCODE_lstore_1: lstore_1,
	cmn.OPCODE_lstore_2: lstore_2,
	cmn.OPCODE_lstore_3: lstore_3,
	cmn.OPCODE_fstore_0: fstore_0,
	cmn.OPCODE_fstore_1: fstore_1,
	cmn.OPCODE_fstore_2: fstore_2,
	cmn.OPCODE_fstore_3: fstore_3,
	cmn.OPCODE_dstore_0: dstore_0,
	cmn.OPCODE_dstore_1: dstore_1,
	cmn.OPCODE_dstore_2: dstore_2,
	cmn.OPCODE_dstore_3: dstore_3,
	cmn.OPCODE_astore_0: astore_0,
	cmn.OPCODE_astore_1: astore_1,
	cmn.OPCODE_astore_2: astore_2,
	cmn.OPCODE_astore_3: astore_3,
	//cmn.OPCODE_iastore:  iastore,
	//cmn.OPCODE_lastore:  lastore,
	//cmn.OPCODE_fastore:  fastore,
	//cmn.OPCODE_dastore:  dastore,
	//cmn.OPCODE_aastore:  aastore,
	//cmn.OPCODE_bastore:  bastore,
	//cmn.OPCODE_castore:  castore,
	//cmn.OPCODE_sastore:  sastore,

	// Stack
	cmn.OPCODE_pop:     pop,
	cmn.OPCODE_pop2:    pop2,
	cmn.OPCODE_dup:     dup,
	cmn.OPCODE_dup_x1:  dup_x1,
	cmn.OPCODE_dup_x2:  dup_x2,
	cmn.OPCODE_dup2:    dup2,
	cmn.OPCODE_dup2_x1: dup2_x1,
	cmn.OPCODE_dup2_x2: dup2_x2,
	cmn.OPCODE_swap:    swap,

	//	Math
	cmn.OPCODE_iadd:  iadd,
	cmn.OPCODE_ladd:  ladd,
	cmn.OPCODE_fadd:  fadd,
	cmn.OPCODE_dadd:  dadd,
	cmn.OPCODE_isub:  isub,
	cmn.OPCODE_lsub:  lsub,
	cmn.OPCODE_fsub:  fsub,
	cmn.OPCODE_dsub:  dsub,
	cmn.OPCODE_imul:  imul,
	cmn.OPCODE_lmul:  lmul,
	cmn.OPCODE_fmul:  fmul,
	cmn.OPCODE_dmul:  dmul,
	cmn.OPCODE_idiv:  idiv,
	cmn.OPCODE_ldiv:  ldiv,
	cmn.OPCODE_fdiv:  fdiv,
	cmn.OPCODE_ddiv:  ddiv,
	cmn.OPCODE_irem:  irem,
	cmn.OPCODE_lrem:  lrem,
	cmn.OPCODE_frem:  frem,
	cmn.OPCODE_drem:  drem,
	cmn.OPCODE_ineg:  ineg,
	cmn.OPCODE_lneg:  lneg,
	cmn.OPCODE_fneg:  fneg,
	cmn.OPCODE_dneg:  dneg,
	cmn.OPCODE_ishl:  ishl,
	cmn.OPCODE_lshl:  lshl,
	cmn.OPCODE_ishr:  ishr,
	cmn.OPCODE_lshr:  lshr,
	cmn.OPCODE_iushr: iushr,
	cmn.OPCODE_lushr: lushr,
	cmn.OPCODE_iand:  iand,
	cmn.OPCODE_land:  land,
	cmn.OPCODE_ior:   ior,
	cmn.OPCODE_lor:   lor,
	cmn.OPCODE_ixor:  ixor,
	cmn.OPCODE_lxor:  lxor,
	cmn.OPCODE_iinc:  iinc,

	// Conversions
	cmn.OPCODE_i2l: i2l,
	cmn.OPCODE_i2f: i2f,
	cmn.OPCODE_i2d: i2d,
	cmn.OPCODE_l2i: l2i,
	cmn.OPCODE_l2f: l2f,
	cmn.OPCODE_l2d: l2d,
	cmn.OPCODE_f2i: f2i,
	cmn.OPCODE_f2l: f2l,
	cmn.OPCODE_f2d: f2d,
	cmn.OPCODE_d2i: d2i,
	cmn.OPCODE_d2l: d2l,
	cmn.OPCODE_d2f: d2f,
	cmn.OPCODE_i2b: i2b,
	cmn.OPCODE_i2c: i2c,
	cmn.OPCODE_i2s: i2s,

	// Comparisons
	cmn.OPCODE_lcmp:      lcmp,
	cmn.OPCODE_fcmpl:     fcmpl,
	cmn.OPCODE_fcmpg:     fcmpg,
	cmn.OPCODE_dcmpl:     dcmpl,
	cmn.OPCODE_dcmpg:     dcmpg,
	cmn.OPCODE_ifeq:      ifeq,
	cmn.OPCODE_ifne:      ifne,
	cmn.OPCODE_iflt:      iflt,
	cmn.OPCODE_ifge:      ifge,
	cmn.OPCODE_ifgt:      ifgt,
	cmn.OPCODE_ifle:      ifle,
	cmn.OPCODE_if_icmpeq: if_icmpeq,
	cmn.OPCODE_if_icmpne: if_icmpne,
	cmn.OPCODE_if_icmplt: if_icmplt,
	cmn.OPCODE_if_icmpge: if_icmpge,
	cmn.OPCODE_if_icmpgt: if_icmpgt,
	cmn.OPCODE_if_icmple: if_icmple,
	cmn.OPCODE_if_acmpeq: if_acmpeq,
	cmn.OPCODE_if_acmpne: if_acmpne,

	// Control
	cmn.OPCODE_ggoto:        ggoto,
	cmn.OPCODE_jsr:          jsr,
	cmn.OPCODE_ret:          ret,
	cmn.OPCODE_tableswitch:  tableswitch,
	cmn.OPCODE_lookupswitch: lookupswitch,
	//cmn.OPCODE_ireturn:      ireturn,
	//cmn.OPCODE_lreturn:      lreturn,
	//cmn.OPCODE_freturn:      freturn,
	//cmn.OPCODE_dreturn:      dreturn,
	//cmn.OPCODE_areturn:      areturn,
	//cmn.OPCODE_rreturn:      rreturn,

	// References
	cmn.OPCODE_getstatic: getstatic,
	cmn.OPCODE_putstatic: putstatic,
	cmn.OPCODE_getfield:  getfield,
	cmn.OPCODE_putfield:  putfield,
	//cmn.OPCODE_invokevirtual:   invokevirtual,
	//cmn.OPCODE_invokespecial:   invokespecial,
	//cmn.OPCODE_invokestatic:    invokestatic,
	//cmn.OPCODE_invokeinterface: invokeinterface,
	//cmn.OPCODE_invokedynamic:   invokedynamic,
	cmn.OPCODE_new: new,
	//cmn.OPCODE_newarray:        newarray,
	//cmn.OPCODE_anewarray:       anewarray,
	//cmn.OPCODE_arraylength:     arraylength,
	//cmn.OPCODE_athrow:          athrow,
	//cmn.OPCODE_checkcast:       checkcast,
	//cmn.OPCODE_instanceof:      instanceof,
	//cmn.OPCODE_monitorenter:    monitorenter,
	//cmn.OPCODE_monitorexit:     monitorexit,

	// Extended
	//cmn.OPCODE_wide:           wide,
	//cmn.OPCODE_multianewarray: multianewarray,
	cmn.OPCODE_ifnull:    ifnull,
	cmn.OPCODE_ifnonnull: ifnonnull,
	cmn.OPCODE_ggoto_w:   ggoto_w,
	cmn.OPCODE_jsr_w:     jsr_w,

	// Reserved is not shown
}

func InstFnc(i uint8) func(*rtdt.Frame) {
	return code2FuncMap[i]
}

func InstFncHit(i uint8) (func(*rtdt.Frame), bool) {
	res, ok := code2FuncMap[i]
	return res, ok
}
