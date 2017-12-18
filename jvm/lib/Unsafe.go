package lib

import (
	"jvmGo/jvm/rtdt"
	"unsafe"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/marea"
	"fmt"
)

func init() {
	register(utils.CLASSNAME_Unsafe, "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	register(utils.CLASSNAME_Unsafe, "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	register(utils.CLASSNAME_Unsafe, "addressSize", "()I", addressSize)
	register(utils.CLASSNAME_Unsafe, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J", objectFieldOffset)
	register(utils.CLASSNAME_Unsafe, "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", compareAndSwapObject)
	register(utils.CLASSNAME_Unsafe, "compareAndSwapInt", "(Ljava/lang/Object;JII)Z", compareAndSwapInt)
	register(utils.CLASSNAME_Unsafe, "getIntVolatile", "(Ljava/lang/Object;J)I", getIntVolatile)
}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtdt.Frame) {
	stack := frame.OperandStack
	stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtdt.Frame) {
	stack := frame.OperandStack
	stack.PushInt(1) // todo
}

// public native int addressSize();
// ()I
func addressSize(frame *rtdt.Frame) {
	// vars := frame.LocalVars()
	// vars.GetRef(0) // this

	stack := frame.OperandStack
	res := int32(unsafe.Sizeof(int(1)))
	//fmt.Println(res)
	stack.PushInt(res)
}

// public native long objectFieldOffset(Field field);
// (Ljava/lang/reflect/Field;)J
func objectFieldOffset(f *rtdt.Frame) {
	jField := f.LocalVar.GetRef(1)
	slotIdx := jField.GetInsFieldSlotIdx("slot")
	slotN := jField.GetInt(slotIdx)
	f.OperandStack.PushLong(int64(slotN))
}

// public final native boolean compareAndSwapObject(Object o, long offset, Object expected, Object x)
// (Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z
func compareAndSwapObject(f *rtdt.Frame) {
	vars := f.LocalVar
	fmt.Println(vars)
	obj := vars.GetRef(1)
	offset := vars.GetLong(2)
	expected := vars.GetRef(4)
	newVar := vars.GetRef(5)

	switch data := obj.Data().(type) {
	case marea.Vars: // normal objetRef()
		if data.GetRef(uint(int32(offset))) == expected {
			data.SetRef(newVar, uint(int32(offset)))
			f.OperandStack.PushInt(1)
		} else {
			f.OperandStack.PushInt(0)
		}
	case []*marea.Object: // ref array
		if data[offset] == expected {
			data[offset] = newVar
			f.OperandStack.PushInt(1)
		} else {
			f.OperandStack.PushInt(0)
		}
	default: //todo
		panic(fmt.Errorf("compareAndSwapObject, wrong type %T!", data))
	}
}

// public final native boolean compareAndSwapInt(object ar1, long var2 ,int var4, int var5);
// (Ljava/lang/Object;JII)Z
func compareAndSwapInt(f *rtdt.Frame) {
	vars := f.LocalVar
	obj := vars.GetRef(1)
	offset := vars.GetLong(2)
	off := uint(uint32(offset))
	expected := vars.GetInt(4)
	newVar := vars.GetInt(5)

	switch data := obj.Data().(type) {
	case marea.Vars:
		if data.GetInt(off) == expected {
			data.SetInt(newVar, off)
			f.OperandStack.PushInt(1)
		} else {
			f.OperandStack.PushInt(0)
		}
	case []int32:
		if data[offset] == expected {
			data[offset] = newVar
			f.OperandStack.PushInt(1)
		} else {
			f.OperandStack.PushInt(0)
		}
	default:
		panic(fmt.Errorf("compareAndSwapInt, wrong type %T!", data))
	}
}

// public native init getIntVolatile(Object obj, long l);
// (Ljava/lang/Object;J)I
func getIntVolatile(f *rtdt.Frame) {
	obj := f.LocalVar.GetRef(1)
	offset := f.LocalVar.GetLong(2)
	off := uint(int32(offset))
	//debug
	fields := obj.Data()
	fmt.Printf("LONG %d INT %d\n", offset, off)
	fmt.Printf("OBJ %s\n", obj.Class().ClassName())
	switch fields := fields.(type) {
	case marea.Vars:
		//debug
		fmt.Printf("num %d\n", len(fields))
		n := fields[off].Num
		f.OperandStack.PushInt(n)
	case []int32:
		n := fields[off]
		f.OperandStack.PushInt(n)
	default:
		panic(fmt.Errorf("getIntVolatile, wrong type %T!", fields))
	}
}
