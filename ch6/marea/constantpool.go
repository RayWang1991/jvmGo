package marea

import (
	cf  "jvmGo/ch6/classfile"
	"fmt"
)

type Constant interface{}

type ConstantPool []Constant

func NewConstantPool(cfcp cf.ConstantPool) ConstantPool {
	cp := make([]Constant, len(cfcp))
	for i, c := range cfcp {
		switch c := c.(type) {
		// literals
		case *cf.IntegerInfo:
			cp[i] = c.Value()
		case *cf.FloatInfo:
			cp[i] = c.Value()
		case *cf.DoubleInfo:
			cp[i] = c.Value()
		case *cf.LongInfo:
			cp[i] = c.Value()
		case *cf.Utf8Info:
			cp[i] = c.String()
		case *cf.StringInfo:
			cp[i] = c.String(cfcp)
		case *cf.NameTypeInfo:
			cp[i] = NewNameTypeInfo(cfcp, c)
		case *cf.ClassInfo:
			cp[i] = NewClassInfo(cfcp, c)
		case *cf.FieldRefInfo:
			cp[i] = NewFieldRef(cfcp, c)
		case *cf.MethodRefInfo:
			cp[i] = NewMethodRef(cfcp, c)
		case *cf.InterfaceMethodRefInfo:
			cp[i] = NewInterfaceMethodRef(cfcp, c)
		case *cf.MethodTypeInfo:
			cp[i] = NewMethodType(cfcp, c)
		case *cf.MethodHandleInfo:
			cp[i] = NewMethodHandle(cfcp, c)
		case *cf.InvokeDynamicInfo:
			cp[i] = NewInvokeDynamic(cfcp, c)
		default:
			panic(fmt.Errorf("unsupported type %T", c))
		}
	}
	return cp
}
