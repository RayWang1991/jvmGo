package lib

import (
	"jvmGo/jvm/rtdt"
	"unsafe"
	"jvmGo/jvm/utils"
	"jvmGo/jvm/marea"
	"fmt"
	"encoding/binary"
)

func init() {
	register(utils.CLASSNAME_Unsafe, "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	register(utils.CLASSNAME_Unsafe, "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	register(utils.CLASSNAME_Unsafe, "addressSize", "()I", addressSize)
	register(utils.CLASSNAME_Unsafe, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J", objectFieldOffset)
	register(utils.CLASSNAME_Unsafe, "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", compareAndSwapObject)
	register(utils.CLASSNAME_Unsafe, "compareAndSwapInt", "(Ljava/lang/Object;JII)Z", compareAndSwapInt)
	register(utils.CLASSNAME_Unsafe, "getIntVolatile", "(Ljava/lang/Object;J)I", getIntVolatile)
	register(utils.CLASSNAME_Unsafe, "allocateMemory", "(J)J", allocateMemory)
	register(utils.CLASSNAME_Unsafe, "putLong", "(JJ)V", putLong)
	register(utils.CLASSNAME_Unsafe, "getByte", "(J)B", getByte)
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

// todo
const _offset = int64(10086)

var _address = _offset
var mem = make([]byte, 2048) // mem is a continuos mem chunk, start size is 2048

func _memIdx(addr int64) int64 {
	return addr - _offset
}

func _addr(idx int64) int64 {
	return idx + _offset
}

// todo test !
// public native long allocateMemory (long var1);
// (J)J
func allocateMemory(f *rtdt.Frame) {
	size := f.LocalVar.GetLong(1)
	addr := _address
	_address += size

	// check if there is enough mem
	cursize := _memIdx(_address) + 1
	if cursize > int64(len(mem)) {

		newLen := int64(2 * len(mem))
		if cursize > newLen {
			newLen = cursize * 2
		}
		// reallocate and copy
		t := mem
		mem = make([]byte, newLen)
		copy(mem, t)
	}

	f.OperandStack.PushLong(addr)
}

// public native void putLong(long var1, long var3);
// (JJ)V
func putLong(f *rtdt.Frame) {
	addr := f.LocalVar.GetLong(1)
	value := f.LocalVar.GetLong(3)
	idx := _memIdx(addr)
	fmt.Printf("slot1 %d slot2 %d addr is %d idx is %d\n", f.LocalVar.GetInt(1), f.LocalVar.GetInt(2), addr, idx)
	bytes := mem[idx:]

	if len(bytes) < 8 {
		panic(utils.IllegalArgumentException) //TODO
	}
	binary.BigEndian.PutUint64(bytes, uint64(value))
}

// public native byte getByte(long var1);
// (J)B
func getByte(f *rtdt.Frame) {
	addr := f.LocalVar.GetLong(1)
	idx := _memIdx(addr)
	bytes := mem[idx:]

	if len(bytes) < 1 {
		panic(utils.IllegalArgumentException) //TODO
	}

	b := int32(bytes[0])
	f.OperandStack.PushInt(b)
}

// public native void freeMemory(long var1);
// (J)V
func freeMemory(f *rtdt.Frame){
	//todo
}