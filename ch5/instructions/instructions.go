package instructions

// instruction set
const (
	OPCODE_aaload          = 0x50
	OPCODE_aastore         = 0x53
	OPCODE_aconst_null     = 0x1
	OPCODE_aload           = 0x19
	OPCODE_aload_0         = 0x2a
	OPCODE_aload_1         = 0x2b
	OPCODE_aload_2         = 0x2c
	OPCODE_aload_3         = 0x2d
	OPCODE_anewarray       = 0xbd
	OPCODE_areturn         = 0xb0
	OPCODE_arraylength     = 0xbe
	OPCODE_astore          = 0x3a
	OPCODE_astore_0        = 0x4b
	OPCODE_astore_1        = 0x4c
	OPCODE_astore_2        = 0x4d
	OPCODE_astore_3        = 0x4e
	OPCODE_athrow          = 0xbf
	OPCODE_baload          = 0x33
	OPCODE_bastore         = 0x54
	OPCODE_bipush          = 0x10
	OPCODE_caload          = 0x34
	OPCODE_castore         = 0x55
	OPCODE_checkcast       = 0xc0
	OPCODE_d2f             = 0x90
	OPCODE_d2i             = 0x8e
	OPCODE_d2l             = 0x8f
	OPCODE_dadd            = 0x63
	OPCODE_daload          = 0x31
	OPCODE_dastore         = 0x52
	OPCODE_dcmpl           = 0x97
	OPCODE_dcmpg           = 0x98
	OPCODE_dconst_0        = 0xe
	OPCODE_dconst_1        = 0xf
	OPCODE_ddiv            = 0x6f
	OPCODE_dload           = 0x18
	OPCODE_dload_0         = 0x26
	OPCODE_dload_1         = 0x27
	OPCODE_dload_2         = 0x28
	OPCODE_dload_3         = 0x29
	OPCODE_dmul            = 0x6b
	OPCODE_dneg            = 0x77
	OPCODE_drem            = 0x73
	OPCODE_dreturn         = 0xaf
	OPCODE_dstore          = 0x39
	OPCODE_dstore_0        = 0x71
	OPCODE_dstore_1        = 0x72
	OPCODE_dstore_2        = 0x73
	OPCODE_dstore_3        = 0x74
	OPCODE_dsub            = 0x67
	OPCODE_dup             = 0x59
	OPCODE_dup_x1          = 0x5a
	OPCODE_dup_x2          = 0x5b
	OPCODE_dup2            = 0x5c
	OPCODE_dup2_x1         = 0x5d
	OPCODE_dup2_x2         = 0x5e
	OPCODE_f2d             = 0x8d
	OPCODE_f2i             = 0x8b
	OPCODE_f2l             = 0x8c
	OPCODE_fadd            = 0x62
	OPCODE_faload          = 0x30
	OPCODE_fastore         = 0x51
	OPCODE_fcmpl           = 0x95
	OPCODE_fcmpg           = 0x96
	OPCODE_fcosnt_0        = 0xb
	OPCODE_fcosnt_1        = 0xc
	OPCODE_fcosnt_2        = 0xd
	OPCODE_fdiv            = 0x6e
	OPCODE_fload           = 0x17
	OPCODE_fload_0         = 0x22
	OPCODE_fload_1         = 0x23
	OPCODE_fload_2         = 0x24
	OPCODE_fload_3         = 0x25
	OPCODE_fmul            = 0x6a
	OPCODE_fneg            = 0x76
	OPCODE_frem            = 0x72
	OPCODE_freturn         = 0xae
	OPCODE_fstore          = 0x38
	OPCODE_fstore_0        = 0x43
	OPCODE_fstore_1        = 0x44
	OPCODE_fstore_2        = 0x45
	OPCODE_fstore_3        = 0x64
	OPCODE_fsub            = 0x66
	OPCODE_getfield        = 0xb4
	OPCODE_getstatic       = 0xb2
	OPCODE_ggoto           = 0xa7
	OPCODE_ggoto_w         = 0xc8
	OPCODE_i2b             = 0x91
	OPCODE_i2c             = 0x92
	OPCODE_i2d             = 0x87
	OPCODE_i2f             = 0x86
	OPCODE_i2l             = 0x85
	OPCODE_i2s             = 0x93
	OPCODE_iadd            = 0x60
	OPCODE_iaload          = 0x2e
	OPCODE_iand            = 0x7e
	OPCODE_iastore         = 0x4f
	OPCODE_iconst_m1       = 0x2
	OPCODE_iconst_0        = 0x3
	OPCODE_iconst_1        = 0x4
	OPCODE_iconst_2        = 0x5
	OPCODE_iconst_3        = 0x6
	OPCODE_iconst_4        = 0x7
	OPCODE_iconst_5        = 0x8
	OPCODE_idiv            = 0x6c
	OPCODE_if_acmpeq       = 0xa5
	OPCODE_if_acmpne       = 0xa6
	OPCODE_if_icmpeq       = 0x9f
	OPCODE_if_icmpne       = 0xa0
	OPCODE_if_icmplt       = 0xa1
	OPCODE_if_icmpge       = 0xa2
	OPCODE_if_icmpgt       = 0xa3
	OPCODE_if_icmple       = 0xa4
	OPCODE_ifeq            = 0x99
	OPCODE_ifne            = 0x9a
	OPCODE_iflt            = 0x9b
	OPCODE_ifge            = 0x9c
	OPCODE_ifgt            = 0x9d
	OPCODE_ifle            = 0x9e
	OPCODE_ifnonnul        = 0xc7
	OPCODE_ifnull          = 0xc6
	OPCODE_iinc            = 0x84
	OPCODE_iload           = 0x15
	OPCODE_iload_0         = 0x1a
	OPCODE_iload_1         = 0x1b
	OPCODE_iload_2         = 0x1c
	OPCODE_iload_3         = 0x1d
	OPCODE_imul            = 0x68
	OPCODE_ineg            = 0x74
	OPCODE_instanceof      = 0xc1
	OPCODE_invokedynamic   = 0xba
	OPCODE_invokeinterface = 0xb9
	OPCODE_invokespecial   = 0xb7
	OPCODE_invokestatic    = 0xb8
	OPCODE_invokevirtual   = 0xb6
	OPCODE_ior             = 0x80
	OPCODE_irem            = 0x70
	OPCODE_ireturn         = 0xac
	OPCODE_ishl            = 0x78
	OPCODE_ishr            = 0x7a
	OPCODE_istore          = 0x36
	OPCODE_istore_0        = 0x3b
	OPCODE_istore_1        = 0x3c
	OPCODE_istore_2        = 0x3d
	OPCODE_istore_3        = 0x3e
	OPCODE_isub            = 0x64
	OPCODE_iushr           = 0x7c
	OPCODE_ixor            = 0x82
	OPCODE_jsr             = 0xa8
	OPCODE_jsr_w           = 0xc9
	OPCODE_l2d             = 0x8a
	OPCODE_l2f             = 0x89
	OPCODE_l2i             = 0x88
	OPCODE_ladd            = 0x61
	OPCODE_laload          = 0x2f
	OPCODE_land            = 0x7f
	OPCODE_lastore         = 0x50
	OPCODE_lcmp            = 0x94
	OPCODE_lconst_0        = 0x9
	OPCODE_lconst_1        = 0xa
	OPCODE_ldc             = 0x12
	OPCODE_ldc_w           = 0x13
	OPCODE_ldc2_w          = 0x14
	OPCODE_ldiv            = 0x6d
	OPCODE_lload           = 0x16
	OPCODE_lload_0         = 0x1e
	OPCODE_lload_1         = 0x1f
	OPCODE_lload_2         = 0x20
	OPCODE_lload_3         = 0x21
	OPCODE_lmul            = 0x69
	OPCODE_lneg            = 0x75
	OPCODE_lookupswitch    = 0xab
	OPCODE_lor             = 0x81
	OPCODE_lrem            = 0x71
	OPCODE_lreturn         = 0xad
	OPCODE_lshl            = 0x79
	OPCODE_lshr            = 0x7b
	OPCODE_lstore          = 0x37
	OPCODE_lstore_0        = 0x3f
	OPCODE_lstore_1        = 0x40
	OPCODE_lstore_2        = 0x41
	OPCODE_lstore_3        = 0x42
	OPCODE_lsub            = 0x65
	OPCODE_lushr           = 0x7d
	OPCODE_lxor            = 0x83
	OPCODE_monitorenter    = 0xc2
	OPCODE_monitorexit     = 0xc3
	OPCODE_multianewarray  = 0xc5
	OPCODE_new             = 0xbb
	OPCODE_newarray        = 0xbc
	OPCODE_nop             = 0x0
	OPCODE_pop             = 0x57
	OPCODE_pop2            = 0x58
	OPCODE_popfield        = 0xb5
	OPCODE_putstatic       = 0xb3
	OPCODE_ret             = 0xa9
	OPCODE_rreturn         = 0xb1
	OPCODE_saload          = 0x35
	OPCODE_sastore         = 0x56
	OPCODE_sipush          = 0x11
	OPCODE_swap            = 0x5f
	OPCODE_tableswitch     = 0xaa
	OPCODE_wide            = 0xc4
)
