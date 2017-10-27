package marea

import (
	"fmt"
	cf "jvmGo/jvm/classfile"
)

type Constant interface{}

type ConstantPool []Constant

func NewConstantPool(cfcp cf.ConstantPool, cls *Class) ConstantPool {
	cp := make([]Constant, len(cfcp))
	skip := false
	for i, c := range cfcp {
		if skip || i == 0 {
			skip = false
			continue
		}
		switch c := c.(type) {
		// literals
		case *cf.IntegerInfo:
			cp[i] = c.Value()
		case *cf.FloatInfo:
			cp[i] = c.Value()
		case *cf.DoubleInfo:
			cp[i] = c.Value()
			skip = true
		case *cf.LongInfo:
			cp[i] = c.Value()
			skip = true
		case *cf.Utf8Info:
			cp[i] = c.String()
		case *cf.StringInfo:
			cp[i] = c.String(cfcp)
		case *cf.NameTypeInfo:
			cp[i] = NewNameTypeInfo(cfcp, c)
		case *cf.ClassInfo:
			cp[i] = NewClassRef(c.ClassName(cfcp), cls)
		case *cf.FieldRefInfo:
			cp[i] = NewFieldRef(cfcp, c, cls)
		case *cf.MethodRefInfo:
			cp[i] = NewMethodRef(cfcp, c, cls)
		case *cf.InterfaceMethodRefInfo:
			cp[i] = NewInterfaceMethodRef(cfcp, c, cls)
		case *cf.MethodTypeInfo:
			cp[i] = NewMethodType(cfcp, c)
		case *cf.MethodHandleInfo:
			cp[i] = NewMethodHandle(cfcp, c, cls)
		case *cf.InvokeDynamicInfo:
			cp[i] = NewInvokeDynamic(cfcp, c)
		default:
			panic(fmt.Errorf("unsupported type %v", c))
		}
	}
	return cp
}

func (cp ConstantPool) GetString(index uint16) string {
	return cp[index].(string)
}

func (cp ConstantPool) GetInteger(index uint16) int32 {
	return cp[index].(int32)
}

func (cp ConstantPool) GetLong(index uint16) int64 {
	return cp[index].(int64)
}

func (cp ConstantPool) GetFloat(index uint16) float32 {
	return cp[index].(float32)
}

func (cp ConstantPool) GetDouble(index uint16) float64 {
	return cp[index].(float64)
}

func (cp ConstantPool) GetClassRef(index uint16) *ClassRef {
	return cp[index].(*ClassRef)
}

func (cp ConstantPool) GetFieldRef(index uint16) *FieldRef {
	return cp[index].(*FieldRef)
}

func (cp ConstantPool) GetMethodRef(index uint16) *MethodRef {
	return cp[index].(*MethodRef)
}

func (cp ConstantPool) GetInterfaceMethodRef(index uint16) *InterfaceMethodRef {
	return cp[index].(*InterfaceMethodRef)
}

func (cp ConstantPool) GetConstant(index uint16) Constant {
	return cp[index]
}

// wrapper for methodRef and interfaceMethodRef

type MethodGetter interface {
	GetMethod() *Method
	GetRef() *MemberRef
}

func (cp ConstantPool) GetMethodFromGetter(index uint16) *Method {
	r := cp[index].(MethodGetter)
	return r.GetMethod()
}

func (cp ConstantPool) GetMemberRefFromGetter(index uint16) *MemberRef {
	r := cp[index].(MethodGetter)
	return r.GetRef()
}
