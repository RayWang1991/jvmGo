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
	slot := jField.GetInsFieldSlotIdx("slot")
	f.OperandStack.PushLong(int64(slot))
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
		panic("unsupported yet")
	}
}
