package instructions

import "jvmGo/ch5/rtdata"

// instruction set
const (
	// Constants
	OPCODE_nop         = 0x0
	OPCODE_aconst_null = 0x1
	OPCODE_iconst_m1   = 0x2
	OPCODE_iconst_0    = 0x3
	OPCODE_iconst_1    = 0x4
	OPCODE_iconst_2    = 0x5
	OPCODE_iconst_3    = 0x6
	OPCODE_iconst_4    = 0x7
	OPCODE_iconst_5    = 0x8
	OPCODE_lconst_0    = 0x9
	OPCODE_lconst_1    = 0xa
	OPCODE_fconst_0    = 0xb
	OPCODE_fconst_1    = 0xc
	OPCODE_fconst_2    = 0xd
	OPCODE_dconst_0    = 0xe
	OPCODE_dconst_1    = 0xf
	OPCODE_bipush      = 0x10
	OPCODE_sipush      = 0x11
	OPCODE_ldc         = 0x12
	OPCODE_ldc_w       = 0x13
	OPCODE_ldc2_w      = 0x14

	// Loads
	OPCODE_iload   = 0x15
	OPCODE_lload   = 0x16
	OPCODE_fload   = 0x17
	OPCODE_dload   = 0x18
	OPCODE_aload   = 0x19
	OPCODE_iload_0 = 0x1a
	OPCODE_iload_1 = 0x1b
	OPCODE_iload_2 = 0x1c
	OPCODE_iload_3 = 0x1d
	OPCODE_lload_0 = 0x1e
	OPCODE_lload_1 = 0x1f
	OPCODE_lload_2 = 0x20
	OPCODE_lload_3 = 0x21
	OPCODE_fload_0 = 0x22
	OPCODE_fload_1 = 0x23
	OPCODE_fload_2 = 0x24
	OPCODE_fload_3 = 0x25
	OPCODE_dload_0 = 0x26
	OPCODE_dload_1 = 0x27
	OPCODE_dload_2 = 0x28
	OPCODE_dload_3 = 0x29
	OPCODE_aload_0 = 0x2a
	OPCODE_aload_1 = 0x2b
	OPCODE_aload_2 = 0x2c
	OPCODE_aload_3 = 0x2d
	OPCODE_iaload  = 0x2e
	OPCODE_laload  = 0x2f
	OPCODE_faload  = 0x30
	OPCODE_daload  = 0x31
	OPCODE_aaload  = 0x32
	OPCODE_baload  = 0x33
	OPCODE_caload  = 0x34
	OPCODE_saload  = 0x35

	// Stores
	OPCODE_istore   = 0x36
	OPCODE_lstore   = 0x37
	OPCODE_fstore   = 0x38
	OPCODE_dstore   = 0x39
	OPCODE_astore   = 0x3a
	OPCODE_istore_0 = 0x3b
	OPCODE_istore_1 = 0x3c
	OPCODE_istore_2 = 0x3d
	OPCODE_istore_3 = 0x3e
	OPCODE_lstore_0 = 0x3f
	OPCODE_lstore_1 = 0x40
	OPCODE_lstore_2 = 0x41
	OPCODE_lstore_3 = 0x42
	OPCODE_fstore_0 = 0x43
	OPCODE_fstore_1 = 0x44
	OPCODE_fstore_2 = 0x45
	OPCODE_fstore_3 = 0x46
	OPCODE_dstore_0 = 0x47
	OPCODE_dstore_1 = 0x48
	OPCODE_dstore_2 = 0x49
	OPCODE_dstore_3 = 0x4a
	OPCODE_astore_0 = 0x4b
	OPCODE_astore_1 = 0x4c
	OPCODE_astore_2 = 0x4d
	OPCODE_astore_3 = 0x4e
	OPCODE_iastore  = 0x4f
	OPCODE_lastore  = 0x50
	OPCODE_fastore  = 0x51
	OPCODE_dastore  = 0x52
	OPCODE_aastore  = 0x53
	OPCODE_bastore  = 0x54
	OPCODE_castore  = 0x55
	OPCODE_sastore  = 0x56

	// Stack
	OPCODE_pop     = 0x57
	OPCODE_pop2    = 0x58
	OPCODE_dup     = 0x59
	OPCODE_dup_x1  = 0x5a
	OPCODE_dup_x2  = 0x5b
	OPCODE_dup2    = 0x5c
	OPCODE_dup2_x1 = 0x5d
	OPCODE_dup2_x2 = 0x5e
	OPCODE_swap    = 0x5f

	//	Math
	OPCODE_iadd  = 0x60
	OPCODE_ladd  = 0x61
	OPCODE_fadd  = 0x62
	OPCODE_dadd  = 0x63
	OPCODE_isub  = 0x64
	OPCODE_lsub  = 0x65
	OPCODE_fsub  = 0x66
	OPCODE_dsub  = 0x67
	OPCODE_imul  = 0x68
	OPCODE_lmul  = 0x69
	OPCODE_fmul  = 0x6a
	OPCODE_dmul  = 0x6b
	OPCODE_idiv  = 0x6c
	OPCODE_ldiv  = 0x6d
	OPCODE_fdiv  = 0x6e
	OPCODE_ddiv  = 0x6f
	OPCODE_irem  = 0x70
	OPCODE_lrem  = 0x71
	OPCODE_frem  = 0x72
	OPCODE_drem  = 0x73
	OPCODE_ineg  = 0x74
	OPCODE_lneg  = 0x75
	OPCODE_fneg  = 0x76
	OPCODE_dneg  = 0x77
	OPCODE_ishl  = 0x78
	OPCODE_lshl  = 0x79
	OPCODE_ishr  = 0x7a
	OPCODE_lshr  = 0x7b
	OPCODE_iushr = 0x7c
	OPCODE_lushr = 0x7d
	OPCODE_iand  = 0x7e
	OPCODE_land  = 0x7f
	OPCODE_ior   = 0x80
	OPCODE_lor   = 0x81
	OPCODE_ixor  = 0x82
	OPCODE_lxor  = 0x83
	OPCODE_iinc  = 0x84

	// Conversions
	OPCODE_i2l = 0x85
	OPCODE_i2f = 0x86
	OPCODE_i2d = 0x87
	OPCODE_l2i = 0x88
	OPCODE_l2f = 0x89
	OPCODE_l2d = 0x8a
	OPCODE_f2i = 0x8b
	OPCODE_f2l = 0x8c
	OPCODE_f2d = 0x8d
	OPCODE_d2i = 0x8e
	OPCODE_d2l = 0x8f
	OPCODE_d2f = 0x90
	OPCODE_i2b = 0x91
	OPCODE_i2c = 0x92
	OPCODE_i2s = 0x93

	// Comparisons
	OPCODE_lcmp      = 0x94
	OPCODE_fcmpl     = 0x95
	OPCODE_fcmpg     = 0x96
	OPCODE_dcmpl     = 0x97
	OPCODE_dcmpg     = 0x98
	OPCODE_ifeq      = 0x99
	OPCODE_ifne      = 0x9a
	OPCODE_iflt      = 0x9b
	OPCODE_ifge      = 0x9c
	OPCODE_ifgt      = 0x9d
	OPCODE_ifle      = 0x9e
	OPCODE_if_icmpeq = 0x9f
	OPCODE_if_icmpne = 0xa0
	OPCODE_if_icmplt = 0xa1
	OPCODE_if_icmpge = 0xa2
	OPCODE_if_icmpgt = 0xa3
	OPCODE_if_icmple = 0xa4
	OPCODE_if_acmpeq = 0xa5
	OPCODE_if_acmpne = 0xa6

	// Control
	OPCODE_ggoto        = 0xa7
	OPCODE_jsr          = 0xa8
	OPCODE_ret          = 0xa9
	OPCODE_tableswitch  = 0xaa
	OPCODE_lookupswitch = 0xab
	OPCODE_ireturn      = 0xac
	OPCODE_lreturn      = 0xad
	OPCODE_freturn      = 0xae
	OPCODE_dreturn      = 0xaf
	OPCODE_areturn      = 0xb0
	OPCODE_rreturn      = 0xb1

	// References
	OPCODE_getstatic       = 0xb2
	OPCODE_putstatic       = 0xb3
	OPCODE_getfield        = 0xb4
	OPCODE_popfield        = 0xb5
	OPCODE_invokevirtual   = 0xb6
	OPCODE_invokespecial   = 0xb7
	OPCODE_invokestatic    = 0xb8
	OPCODE_invokeinterface = 0xb9
	OPCODE_invokedynamic   = 0xba
	OPCODE_new             = 0xbb
	OPCODE_newarray        = 0xbc
	OPCODE_anewarray       = 0xbd
	OPCODE_arraylength     = 0xbe
	OPCODE_athrow          = 0xbf
	OPCODE_checkcast       = 0xc0
	OPCODE_instanceof      = 0xc1
	OPCODE_monitorenter    = 0xc2
	OPCODE_monitorexit     = 0xc3

	// Extended
	OPCODE_wide           = 0xc4
	OPCODE_multianewarray = 0xc5
	OPCODE_ifnull         = 0xc6
	OPCODE_ifnonnull      = 0xc7
	OPCODE_ggoto_w        = 0xc8
	OPCODE_jsr_w          = 0xc9

	// Reserved is not shown
)

var code2StringMap = map[uint8]string{
	OPCODE_aaload:          "aaload",
	OPCODE_aastore:         "aastore",
	OPCODE_aconst_null:     "aconst_null",
	OPCODE_aload:           "aload",
	OPCODE_aload_0:         "aload_0",
	OPCODE_aload_1:         "aload_1",
	OPCODE_aload_2:         "aload_2",
	OPCODE_aload_3:         "aload_3",
	OPCODE_anewarray:       "anewarray",
	OPCODE_areturn:         "areturn",
	OPCODE_arraylength:     "arraylength",
	OPCODE_astore:          "astore",
	OPCODE_astore_0:        "astore_0",
	OPCODE_astore_1:        "astore_1",
	OPCODE_astore_2:        "astore_2",
	OPCODE_astore_3:        "astore_3",
	OPCODE_athrow:          "athrow",
	OPCODE_baload:          "baload",
	OPCODE_bastore:         "bastore",
	OPCODE_bipush:          "bipush",
	OPCODE_caload:          "caload",
	OPCODE_castore:         "castore",
	OPCODE_checkcast:       "checkcast",
	OPCODE_d2f:             "d2f",
	OPCODE_d2i:             "d2i",
	OPCODE_d2l:             "d2l",
	OPCODE_dadd:            "dadd",
	OPCODE_daload:          "daload",
	OPCODE_dastore:         "dastore",
	OPCODE_dcmpl:           "dcmpl",
	OPCODE_dcmpg:           "dcmpg",
	OPCODE_dconst_0:        "dconst_0",
	OPCODE_dconst_1:        "dconst_1",
	OPCODE_ddiv:            "ddiv",
	OPCODE_dload:           "dload",
	OPCODE_dload_0:         "dload_0",
	OPCODE_dload_1:         "dload_1",
	OPCODE_dload_2:         "dload_2",
	OPCODE_dload_3:         "dload_3",
	OPCODE_dmul:            "dmul",
	OPCODE_dneg:            "dneg",
	OPCODE_drem:            "drem",
	OPCODE_dreturn:         "dreturn",
	OPCODE_dstore:          "dstore",
	OPCODE_dstore_0:        "dstore_0",
	OPCODE_dstore_1:        "dstore_1",
	OPCODE_dstore_2:        "dstore_2",
	OPCODE_dstore_3:        "dstore_3",
	OPCODE_dsub:            "dsub",
	OPCODE_dup:             "dup",
	OPCODE_dup_x1:          "dup_x1",
	OPCODE_dup_x2:          "dup_x2",
	OPCODE_dup2:            "dup2",
	OPCODE_dup2_x1:         "dup2_x1",
	OPCODE_dup2_x2:         "dup2_x2",
	OPCODE_f2d:             "f2d",
	OPCODE_f2i:             "f2i",
	OPCODE_f2l:             "f2l",
	OPCODE_fadd:            "fadd",
	OPCODE_faload:          "faload",
	OPCODE_fastore:         "fastore",
	OPCODE_fcmpl:           "fcmpl",
	OPCODE_fcmpg:           "fcmpg",
	OPCODE_fconst_0:        "fconst_0",
	OPCODE_fconst_1:        "fconst_1",
	OPCODE_fconst_2:        "fconst_2",
	OPCODE_fdiv:            "fdiv",
	OPCODE_fload:           "fload",
	OPCODE_fload_0:         "fload_0",
	OPCODE_fload_1:         "fload_1",
	OPCODE_fload_2:         "fload_2",
	OPCODE_fload_3:         "fload_3",
	OPCODE_fmul:            "fmul",
	OPCODE_fneg:            "fneg",
	OPCODE_frem:            "frem",
	OPCODE_freturn:         "freturn",
	OPCODE_fstore:          "fstore",
	OPCODE_fstore_0:        "fstore_0",
	OPCODE_fstore_1:        "fstore_1",
	OPCODE_fstore_2:        "fstore_2",
	OPCODE_fstore_3:        "fstore_3",
	OPCODE_fsub:            "fsub",
	OPCODE_getfield:        "getfield",
	OPCODE_getstatic:       "getstatic",
	OPCODE_ggoto:           "ggoto",
	OPCODE_ggoto_w:         "ggoto_w",
	OPCODE_i2b:             "i2b",
	OPCODE_i2c:             "i2c",
	OPCODE_i2d:             "i2d",
	OPCODE_i2f:             "i2f",
	OPCODE_i2l:             "i2l",
	OPCODE_i2s:             "i2s",
	OPCODE_iadd:            "iadd",
	OPCODE_iaload:          "iaload",
	OPCODE_iand:            "iand",
	OPCODE_iastore:         "iastore",
	OPCODE_iconst_m1:       "iconst_m1",
	OPCODE_iconst_0:        "iconst_0",
	OPCODE_iconst_1:        "iconst_1",
	OPCODE_iconst_2:        "iconst_2",
	OPCODE_iconst_3:        "iconst_3",
	OPCODE_iconst_4:        "iconst_4",
	OPCODE_iconst_5:        "iconst_5",
	OPCODE_idiv:            "idiv",
	OPCODE_if_acmpeq:       "if_acmpeq",
	OPCODE_if_acmpne:       "if_acmpne",
	OPCODE_if_icmpeq:       "if_icmpeq",
	OPCODE_if_icmpne:       "if_icmpne",
	OPCODE_if_icmplt:       "if_icmplt",
	OPCODE_if_icmpge:       "if_icmpge",
	OPCODE_if_icmpgt:       "if_icmpgt",
	OPCODE_if_icmple:       "if_icmple",
	OPCODE_ifeq:            "ifeq",
	OPCODE_ifne:            "ifne",
	OPCODE_iflt:            "iflt",
	OPCODE_ifge:            "ifge",
	OPCODE_ifgt:            "ifgt",
	OPCODE_ifle:            "ifle",
	OPCODE_ifnonnull:       "ifnonnull",
	OPCODE_ifnull:          "ifnull",
	OPCODE_iinc:            "iinc",
	OPCODE_iload:           "iload",
	OPCODE_iload_0:         "iload_0",
	OPCODE_iload_1:         "iload_1",
	OPCODE_iload_2:         "iload_2",
	OPCODE_iload_3:         "iload_3",
	OPCODE_imul:            "imul",
	OPCODE_ineg:            "ineg",
	OPCODE_instanceof:      "instanceof",
	OPCODE_invokedynamic:   "invokedynamic",
	OPCODE_invokeinterface: "invokeinterface",
	OPCODE_invokespecial:   "invokespecial",
	OPCODE_invokestatic:    "invokestatic",
	OPCODE_invokevirtual:   "invokevirtual",
	OPCODE_ior:             "ior",
	OPCODE_irem:            "irem",
	OPCODE_ireturn:         "ireturn",
	OPCODE_ishl:            "ishl",
	OPCODE_ishr:            "ishr",
	OPCODE_istore:          "istore",
	OPCODE_istore_0:        "istore_0",
	OPCODE_istore_1:        "istore_1",
	OPCODE_istore_2:        "istore_2",
	OPCODE_istore_3:        "istore_3",
	OPCODE_isub:            "isub",
	OPCODE_iushr:           "iushr",
	OPCODE_ixor:            "ixor",
	OPCODE_jsr:             "jsr",
	OPCODE_jsr_w:           "jsr_w",
	OPCODE_l2d:             "l2d",
	OPCODE_l2f:             "l2f",
	OPCODE_l2i:             "l2i",
	OPCODE_ladd:            "ladd",
	OPCODE_laload:          "laload",
	OPCODE_land:            "land",
	OPCODE_lastore:         "lastore",
	OPCODE_lcmp:            "lcmp",
	OPCODE_lconst_0:        "lconst_0",
	OPCODE_lconst_1:        "lconst_1",
	OPCODE_ldc:             "ldc",
	OPCODE_ldc_w:           "ldc_w",
	OPCODE_ldc2_w:          "ldc2_w",
	OPCODE_ldiv:            "ldiv",
	OPCODE_lload:           "lload",
	OPCODE_lload_0:         "lload_0",
	OPCODE_lload_1:         "lload_1",
	OPCODE_lload_2:         "lload_2",
	OPCODE_lload_3:         "lload_3",
	OPCODE_lmul:            "lmul",
	OPCODE_lneg:            "lneg",
	OPCODE_lookupswitch:    "lookupswitch",
	OPCODE_lor:             "lor",
	OPCODE_lrem:            "lrem",
	OPCODE_lreturn:         "lreturn",
	OPCODE_lshl:            "lshl",
	OPCODE_lshr:            "lshr",
	OPCODE_lstore:          "lstore",
	OPCODE_lstore_0:        "lstore_0",
	OPCODE_lstore_1:        "lstore_1",
	OPCODE_lstore_2:        "lstore_2",
	OPCODE_lstore_3:        "lstore_3",
	OPCODE_lsub:            "lsub",
	OPCODE_lushr:           "lushr",
	OPCODE_lxor:            "lxor",
	OPCODE_monitorenter:    "monitorenter",
	OPCODE_monitorexit:     "monitorexit",
	OPCODE_multianewarray:  "multianewarray",
	OPCODE_new:             "new",
	OPCODE_newarray:        "newarray",
	OPCODE_nop:             "nop",
	OPCODE_pop:             "pop",
	OPCODE_pop2:            "pop2",
	OPCODE_popfield:        "popfield",
	OPCODE_putstatic:       "putstatic",
	OPCODE_ret:             "ret",
	OPCODE_rreturn:         "rreturn",
	OPCODE_saload:          "saload",
	OPCODE_sastore:         "sastore",
	OPCODE_sipush:          "sipush",
	OPCODE_swap:            "swap",
	OPCODE_tableswitch:     "tableswitch",
	OPCODE_wide:            "wide",
}

var code2FuncMap = map[uint8]func(*rtdata.Frame){
	// Constants
	OPCODE_nop:         nop,
	OPCODE_aconst_null: aconst_null,
	OPCODE_iconst_m1:   iconst_m1,
	OPCODE_iconst_0:    iconst_0,
	OPCODE_iconst_1:    iconst_1,
	OPCODE_iconst_2:    iconst_2,
	OPCODE_iconst_3:    iconst_3,
	OPCODE_iconst_4:    iconst_4,
	OPCODE_iconst_5:    iconst_5,
	OPCODE_lconst_0:    lconst_0,
	OPCODE_lconst_1:    lconst_1,
	OPCODE_fconst_0:    fconst_0,
	OPCODE_fconst_1:    fconst_1,
	OPCODE_fconst_2:    fconst_2,
	OPCODE_dconst_0:    dconst_0,
	OPCODE_dconst_1:    dconst_1,
	OPCODE_bipush:      bipush,
	OPCODE_sipush:      sipush,
	//OPCODE_ldc:         ldc,
	//OPCODE_ldc_w:       ldc_w,
	//OPCODE_ldc2_w:      ldc2_w,

	// Loads
	OPCODE_iload:   iload,
	OPCODE_lload:   lload,
	OPCODE_fload:   fload,
	OPCODE_dload:   dload,
	OPCODE_aload:   aload,
	OPCODE_iload_0: iload_0,
	OPCODE_iload_1: iload_1,
	OPCODE_iload_2: iload_2,
	OPCODE_iload_3: iload_3,
	OPCODE_lload_0: lload_0,
	OPCODE_lload_1: lload_1,
	OPCODE_lload_2: lload_2,
	OPCODE_lload_3: lload_3,
	OPCODE_fload_0: fload_0,
	OPCODE_fload_1: fload_1,
	OPCODE_fload_2: fload_2,
	OPCODE_fload_3: fload_3,
	OPCODE_dload_0: dload_0,
	OPCODE_dload_1: dload_1,
	OPCODE_dload_2: dload_2,
	OPCODE_dload_3: dload_3,
	OPCODE_aload_0: aload_0,
	OPCODE_aload_1: aload_1,
	OPCODE_aload_2: aload_2,
	OPCODE_aload_3: aload_3,
	//OPCODE_iaload:  iaload,
	//OPCODE_laload:  laload,
	//OPCODE_faload:  faload,
	//OPCODE_daload:  daload,
	//OPCODE_aaload:  aaload,
	//OPCODE_baload:  baload,
	//OPCODE_caload:  caload,
	//OPCODE_saload:  saload,

	// Stores
	OPCODE_istore:   istore,
	OPCODE_lstore:   lstore,
	OPCODE_fstore:   fstore,
	OPCODE_dstore:   dstore,
	OPCODE_astore:   astore,
	OPCODE_istore_0: istore_0,
	OPCODE_istore_1: istore_1,
	OPCODE_istore_2: istore_2,
	OPCODE_istore_3: istore_3,
	OPCODE_lstore_0: lstore_0,
	OPCODE_lstore_1: lstore_1,
	OPCODE_lstore_2: lstore_2,
	OPCODE_lstore_3: lstore_3,
	OPCODE_fstore_0: fstore_0,
	OPCODE_fstore_1: fstore_1,
	OPCODE_fstore_2: fstore_2,
	OPCODE_fstore_3: fstore_3,
	OPCODE_dstore_0: dstore_0,
	OPCODE_dstore_1: dstore_1,
	OPCODE_dstore_2: dstore_2,
	OPCODE_dstore_3: dstore_3,
	OPCODE_astore_0: astore_0,
	OPCODE_astore_1: astore_1,
	OPCODE_astore_2: astore_2,
	OPCODE_astore_3: astore_3,
	//OPCODE_iastore:  iastore,
	//OPCODE_lastore:  lastore,
	//OPCODE_fastore:  fastore,
	//OPCODE_dastore:  dastore,
	//OPCODE_aastore:  aastore,
	//OPCODE_bastore:  bastore,
	//OPCODE_castore:  castore,
	//OPCODE_sastore:  sastore,

	// Stack
	OPCODE_pop:     pop,
	OPCODE_pop2:    pop2,
	OPCODE_dup:     dup,
	OPCODE_dup_x1:  dup_x1,
	OPCODE_dup_x2:  dup_x2,
	OPCODE_dup2:    dup2,
	OPCODE_dup2_x1: dup2_x1,
	OPCODE_dup2_x2: dup2_x2,
	OPCODE_swap:    swap,

	//	Math
	OPCODE_iadd:  iadd,
	OPCODE_ladd:  ladd,
	OPCODE_fadd:  fadd,
	OPCODE_dadd:  dadd,
	OPCODE_isub:  isub,
	OPCODE_lsub:  lsub,
	OPCODE_fsub:  fsub,
	OPCODE_dsub:  dsub,
	OPCODE_imul:  imul,
	OPCODE_lmul:  lmul,
	OPCODE_fmul:  fmul,
	OPCODE_dmul:  dmul,
	OPCODE_idiv:  idiv,
	OPCODE_ldiv:  ldiv,
	OPCODE_fdiv:  fdiv,
	OPCODE_ddiv:  ddiv,
	OPCODE_irem:  irem,
	OPCODE_lrem:  lrem,
	OPCODE_frem:  frem,
	OPCODE_drem:  drem,
	OPCODE_ineg:  ineg,
	OPCODE_lneg:  lneg,
	OPCODE_fneg:  fneg,
	OPCODE_dneg:  dneg,
	OPCODE_ishl:  ishl,
	OPCODE_lshl:  lshl,
	OPCODE_ishr:  ishr,
	OPCODE_lshr:  lshr,
	OPCODE_iushr: iushr,
	OPCODE_lushr: lushr,
	OPCODE_iand:  iand,
	OPCODE_land:  land,
	OPCODE_ior:   ior,
	OPCODE_lor:   lor,
	OPCODE_ixor:  ixor,
	OPCODE_lxor:  lxor,
	OPCODE_iinc:  iinc,


	// Conversions
	OPCODE_i2l: i2l,
	OPCODE_i2f: i2f,
	OPCODE_i2d: i2d,
	OPCODE_l2i: l2i,
	OPCODE_l2f: l2f,
	OPCODE_l2d: l2d,
	OPCODE_f2i: f2i,
	OPCODE_f2l: f2l,
	OPCODE_f2d: f2d,
	OPCODE_d2i: d2i,
	OPCODE_d2l: d2l,
	OPCODE_d2f: d2f,
	OPCODE_i2b: i2b,
	OPCODE_i2c: i2c,
	OPCODE_i2s: i2s,

	// Comparisons
	OPCODE_lcmp:      lcmp,
	OPCODE_fcmpl:     fcmpl,
	OPCODE_fcmpg:     fcmpg,
	OPCODE_dcmpl:     dcmpl,
	OPCODE_dcmpg:     dcmpg,
	OPCODE_ifeq:      ifeq,
	OPCODE_ifne:      ifne,
	OPCODE_iflt:      iflt,
	OPCODE_ifge:      ifge,
	OPCODE_ifgt:      ifgt,
	OPCODE_ifle:      ifle,
	OPCODE_if_icmpeq: if_icmpeq,
	OPCODE_if_icmpne: if_icmpne,
	OPCODE_if_icmplt: if_icmplt,
	OPCODE_if_icmpge: if_icmpge,
	OPCODE_if_icmpgt: if_icmpgt,
	OPCODE_if_icmple: if_icmple,
	OPCODE_if_acmpeq: if_acmpeq,
	OPCODE_if_acmpne: if_acmpne,

	// Control
	OPCODE_ggoto:        ggoto,
	OPCODE_jsr:          jsr,
	OPCODE_ret:          ret,
	OPCODE_tableswitch:  tableswitch,
	OPCODE_lookupswitch: lookupswitch,
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
	OPCODE_ifnull:         ifnull,
	OPCODE_ifnonnull:      ifnonnull,
	OPCODE_ggoto_w:        ggoto_w,
	OPCODE_jsr_w:          jsr_w,

	// Reserved is not shown
}

func InstStr(i uint8) string {
	return code2StringMap[i]
}

func InstStrHit(i uint8) (string, bool) {
	res, ok := code2StringMap[i]
	return res, ok
}

func InstFnc(i uint8) func(*rtdata.Frame) {
	return code2FuncMap[i]
}

func InstFncHit(i uint8) (func(*rtdata.Frame), bool) {
	res, ok := code2FuncMap[i]
	return res, ok
}
