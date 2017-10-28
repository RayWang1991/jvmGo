package instructions

import (
	"jvmGo/jvm/cmn"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
)

// Loads
// iload
func iload(f *rtdt.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetInt(uint(i))
	f.OperandStack.PushInt(v)
}
func iload_0(f *rtdt.Frame) {
	v := f.LocalVar.GetInt(uint(0))
	f.OperandStack.PushInt(v)
}
func iload_1(f *rtdt.Frame) {
	v := f.LocalVar.GetInt(uint(1))
	f.OperandStack.PushInt(v)
}
func iload_2(f *rtdt.Frame) {
	v := f.LocalVar.GetInt(uint(2))
	f.OperandStack.PushInt(v)
}
func iload_3(f *rtdt.Frame) {
	v := f.LocalVar.GetInt(uint(3))
	f.OperandStack.PushInt(v)
}

//lload
func lload(f *rtdt.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetLong(uint(i))
	f.OperandStack.PushLong(v)
}
func lload_0(f *rtdt.Frame) {
	v := f.LocalVar.GetLong(uint(0))
	f.OperandStack.PushLong(v)
}
func lload_1(f *rtdt.Frame) {
	v := f.LocalVar.GetLong(uint(1))
	f.OperandStack.PushLong(v)
}
func lload_2(f *rtdt.Frame) {
	v := f.LocalVar.GetLong(uint(2))
	f.OperandStack.PushLong(v)
}
func lload_3(f *rtdt.Frame) {
	v := f.LocalVar.GetLong(uint(3))
	f.OperandStack.PushLong(v)
}

// fload
func fload(f *rtdt.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetFloat(uint(i))
	f.OperandStack.PushFloat(v)
}
func fload_0(f *rtdt.Frame) {
	v := f.LocalVar.GetFloat(uint(0))
	f.OperandStack.PushFloat(v)
}
func fload_1(f *rtdt.Frame) {
	v := f.LocalVar.GetFloat(uint(1))
	f.OperandStack.PushFloat(v)
}
func fload_2(f *rtdt.Frame) {
	v := f.LocalVar.GetFloat(uint(2))
	f.OperandStack.PushFloat(v)
}
func fload_3(f *rtdt.Frame) {
	v := f.LocalVar.GetFloat(uint(3))
	f.OperandStack.PushFloat(v)
}

// dload
func dload(f *rtdt.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetDouble(uint(i))
	f.OperandStack.PushDouble(v)
}
func dload_0(f *rtdt.Frame) {
	v := f.LocalVar.GetDouble(uint(0))
	f.OperandStack.PushDouble(v)
}
func dload_1(f *rtdt.Frame) {
	v := f.LocalVar.GetDouble(uint(1))
	f.OperandStack.PushDouble(v)
}
func dload_2(f *rtdt.Frame) {
	v := f.LocalVar.GetDouble(uint(2))
	f.OperandStack.PushDouble(v)
}
func dload_3(f *rtdt.Frame) {
	v := f.LocalVar.GetDouble(uint(3))
	f.OperandStack.PushDouble(v)
}

func aload(f *rtdt.Frame) {
	i := f.ReadU8()
	v := f.LocalVar.GetRef(uint(i))
	f.OperandStack.PushRef(v)
}

func aload_0(f *rtdt.Frame) {
	v := f.LocalVar.GetRef(uint(0))
	f.OperandStack.PushRef(v)
}

func aload_1(f *rtdt.Frame) {
	v := f.LocalVar.GetRef(uint(1))
	f.OperandStack.PushRef(v)
}

func aload_2(f *rtdt.Frame) {
	v := f.LocalVar.GetRef(uint(2))
	f.OperandStack.PushRef(v)
}

func aload_3(f *rtdt.Frame) {
	v := f.LocalVar.GetRef(uint(3))
	f.OperandStack.PushRef(v)
}

func iaload(f *rtdt.Frame) {
	indx := f.OperandStack.PopInt()
	arref := f.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	length := arref.ArrayLengthI()
	if indx < 0 || indx > length-1 {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	f.OperandStack.PushInt(arref.ArrGetInts()[indx])
}

func laload(f *rtdt.Frame) {
	indx := f.OperandStack.PopInt()
	arref := f.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	length := arref.ArrayLengthJ()
	if indx < 0 || indx > length-1 {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	f.OperandStack.PushLong(arref.ArrGetLongs()[indx])
}

func faload(f *rtdt.Frame) {
	indx := f.OperandStack.PopInt()
	arref := f.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	length := arref.ArrayLengthF()
	if indx < 0 || indx > length-1 {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	f.OperandStack.PushFloat(arref.ArrGetFloats()[indx])
}

func daload(f *rtdt.Frame) {
	indx := f.OperandStack.PopInt()
	arref := f.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	length := arref.ArrayLengthD()
	if indx < 0 || indx > length-1 {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	f.OperandStack.PushDouble(arref.ArrGetDoubles()[indx])
}

func aaload(f *rtdt.Frame) {
	indx := f.OperandStack.PopInt()
	arref := f.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	length := arref.ArrayLengthA()
	if indx < 0 || indx > length-1 {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	f.OperandStack.PushRef(arref.ArrGetRefs()[indx])
}

func baload(f *rtdt.Frame) {
	indx := f.OperandStack.PopInt()
	arref := f.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	length := arref.ArrayLengthB()
	if indx < 0 || indx > length-1 {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	f.OperandStack.PushInt(int32(arref.ArrGetBytes()[indx]))
}

func caload(f *rtdt.Frame) {
	indx := f.OperandStack.PopInt()
	arref := f.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	length := arref.ArrayLengthC()
	if indx < 0 || indx > length-1 {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	f.OperandStack.PushInt(int32(arref.ArrGetChars()[indx]))
}

func saload(f *rtdt.Frame) {
	indx := f.OperandStack.PopInt()
	arref := f.OperandStack.PopRef()
	if arref == nil {
		panic(utils.NullPointerException)
	}
	length := arref.ArrayLengthS()
	if indx < 0 || indx > length-1 {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	f.OperandStack.PushInt(int32(arref.ArrGetShorts()[indx]))
}

// Stores
func istore(f *rtdt.Frame) {
	r := f.OperandStack.PopInt()
	i := f.ReadU8()
	f.LocalVar.SetInt(r, uint(i))
}
func istore_0(f *rtdt.Frame) {
	r := f.OperandStack.PopInt()
	f.LocalVar.SetInt(r, 0)
}
func istore_1(f *rtdt.Frame) {
	r := f.OperandStack.PopInt()
	f.LocalVar.SetInt(r, 1)
}
func istore_2(f *rtdt.Frame) {
	r := f.OperandStack.PopInt()
	f.LocalVar.SetInt(r, 2)
}
func istore_3(f *rtdt.Frame) {
	r := f.OperandStack.PopInt()
	f.LocalVar.SetInt(r, 3)
}

func dstore(f *rtdt.Frame) {
	i := f.ReadU8()
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, uint(i))
}
func dstore_0(f *rtdt.Frame) {
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, 0)
}
func dstore_1(f *rtdt.Frame) {
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, 1)
}
func dstore_2(f *rtdt.Frame) {
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, 2)
}
func dstore_3(f *rtdt.Frame) {
	r := f.OperandStack.PopDouble()
	f.LocalVar.SetDouble(r, 3)
}

func fstore(f *rtdt.Frame) {
	i := f.ReadU8()
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, uint(i))
}
func fstore_0(f *rtdt.Frame) {
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, 0)
}
func fstore_1(f *rtdt.Frame) {
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, 1)
}
func fstore_2(f *rtdt.Frame) {
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, 2)
}
func fstore_3(f *rtdt.Frame) {
	r := f.OperandStack.PopFloat()
	f.LocalVar.SetFloat(r, 3)
}

func astore(f *rtdt.Frame) {
	i := f.ReadU8()
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, uint(i))
}
func astore_0(f *rtdt.Frame) {
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, 0)
}
func astore_1(f *rtdt.Frame) {
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, 1)
}
func astore_2(f *rtdt.Frame) {
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, 2)
}
func astore_3(f *rtdt.Frame) {
	r := f.OperandStack.PopRef()
	f.LocalVar.SetRef(r, 3)
}

func lstore(f *rtdt.Frame) {
	i := f.ReadU8()
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, uint(i))
}
func lstore_0(f *rtdt.Frame) {
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, 0)
}
func lstore_1(f *rtdt.Frame) {
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, 1)
}
func lstore_2(f *rtdt.Frame) {
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, 2)
}
func lstore_3(f *rtdt.Frame) {
	r := f.OperandStack.PopLong()
	f.LocalVar.SetLong(r, 3)
}

func iastore(f *rtdt.Frame) {
	val := f.OperandStack.PopInt()
	index := f.OperandStack.PopInt()
	arrayRef := f.OperandStack.PopRef()
	if arrayRef == nil {
		panic(utils.NullPointerException)
	}
	arrayLen := arrayRef.ArrayLengthI()
	if index < 0 || index+1 > arrayLen {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	cn := arrayRef.Class().ClassName()
	if cn != "[I" {
		panic(utils.ArrayStoreException)
	}
	arrayRef.ArrGetInts()[index] = val
}

func lastore(f *rtdt.Frame) {
	val := f.OperandStack.PopLong()
	index := f.OperandStack.PopInt()
	arrayRef := f.OperandStack.PopRef()
	if arrayRef == nil {
		panic(utils.NullPointerException)
	}
	arrayLen := arrayRef.ArrayLengthJ()
	if index < 0 || index+1 > arrayLen {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	cn := arrayRef.Class().ClassName()
	if cn != "[J" {
		panic(utils.ArrayStoreException)
	}
	arrayRef.ArrGetLongs()[index] = val
}

func fastore(f *rtdt.Frame) {
	val := f.OperandStack.PopFloat()
	index := f.OperandStack.PopInt()
	arrayRef := f.OperandStack.PopRef()
	if arrayRef == nil {
		panic(utils.NullPointerException)
	}
	arrayLen := arrayRef.ArrayLengthF()
	if index < 0 || index+1 > arrayLen {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	cn := arrayRef.Class().ClassName()
	if cn != "[F" {
		panic(utils.ArrayStoreException)
	}
	arrayRef.ArrGetFloats()[index] = val
}

func dastore(f *rtdt.Frame) {
	val := f.OperandStack.PopDouble()
	index := f.OperandStack.PopInt()
	arrayRef := f.OperandStack.PopRef()
	if arrayRef == nil {
		panic(utils.NullPointerException)
	}
	arrayLen := arrayRef.ArrayLengthD()
	if index < 0 || index+1 > arrayLen {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	cn := arrayRef.Class().ClassName()
	if cn != "[D" {
		panic(utils.ArrayStoreException)
	}
	arrayRef.ArrGetDoubles()[index] = val
}

func aastore(f *rtdt.Frame) {
	val := f.OperandStack.PopRef()
	if val == nil {
		panic(utils.NullPointerException)
	}
	index := f.OperandStack.PopInt()
	arrayRef := f.OperandStack.PopRef()
	if arrayRef == nil {
		panic(utils.NullPointerException)
	}
	arrayLen := arrayRef.ArrayLengthA()
	if index < 0 || index+1 > arrayLen {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	arrC := arrayRef.Class()
	eleN := cmn.ElementName(arrC.ClassName())
	eleC := arrC.DefineLoader().Load(eleN) // must not be primitive type
	if !marea.IsAssignable(val.Class(), eleC) {
		panic(utils.ArrayStoreException)
	}
	arrayRef.ArrGetRefs()[index] = val
}

func bastore(f *rtdt.Frame) {
	val := int8(f.OperandStack.PopInt())
	index := f.OperandStack.PopInt()
	arrayRef := f.OperandStack.PopRef()
	if arrayRef == nil {
		panic(utils.NullPointerException)
	}
	arrayLen := arrayRef.ArrayLengthB()
	if index < 0 || index+1 > arrayLen {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	cn := arrayRef.Class().ClassName()
	if cn != "[B" && cn != "[Z" {
		panic(utils.ArrayStoreException)
	}
	arrayRef.ArrGetBytes()[index] = val
}

func castore(f *rtdt.Frame) {
	val := uint16(f.OperandStack.PopInt())
	index := f.OperandStack.PopInt()
	arrayRef := f.OperandStack.PopRef()
	if arrayRef == nil {
		panic(utils.NullPointerException)
	}
	arrayLen := arrayRef.ArrayLengthC()
	if index < 0 || index+1 > arrayLen {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	cn := arrayRef.Class().ClassName()
	if cn != "[C" {
		panic(utils.ArrayStoreException)
	}
	arrayRef.ArrGetChars()[index] = val
}

func sastore(f *rtdt.Frame) {
	val := int16(f.OperandStack.PopInt())
	index := f.OperandStack.PopInt()
	arrayRef := f.OperandStack.PopRef()
	if arrayRef == nil {
		panic(utils.NullPointerException)
	}
	arrayLen := arrayRef.ArrayLengthS()
	if index < 0 || index+1 > arrayLen {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	cn := arrayRef.Class().ClassName()
	if cn != "[S" {
		panic(utils.ArrayStoreException)
	}
	arrayRef.ArrGetShorts()[index] = val
}
