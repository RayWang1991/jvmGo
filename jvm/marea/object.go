package marea

import (
	"math"
	"bytes"
	"fmt"
	"unsafe"
	"jvmGo/jvm/cmn"
)

// An Object represents a reference type (non-array type and array type is different)
type Object struct {
	class *Class
	data  interface{}
}

func (o *Object) Class() *Class {
	return o.class
}

//
func NewArray(class *Class, length int32) *Object {
	switch class.name[1] { // first is '['
	case 'B', 'Z':
		return NewArrayB(class, length)
	case 'C':
		return NewArrayC(class, length)
	case 'S':
		return NewArrayS(class, length)
	case 'I':
		return NewArrayI(class, length)
	case 'J':
		return NewArrayJ(class, length)
	case 'F':
		return NewArrayF(class, length)
	case 'D':
		return NewArrayD(class, length)
	default:
		return NewArrayA(class, length)
	}
}

func NewArrayB(class *Class, length int32) *Object {
	return &Object{
		class: class,
		data:  make([]int8, length),
	}
}

func NewArrayC(class *Class, length int32) *Object {
	return &Object{
		class: class,
		data:  make([]uint16, length),
	}
}

func NewArrayS(class *Class, length int32) *Object {
	return &Object{
		class: class,
		data:  make([]int16, length),
	}
}

func NewArrayI(class *Class, length int32) *Object {
	return &Object{
		class: class,
		data:  make([]int32, length),
	}
}

func NewArrayJ(class *Class, length int32) *Object {
	return &Object{
		class: class,
		data:  make([]int64, length),
	}
}

func NewArrayF(class *Class, length int32) *Object {
	return &Object{
		class: class,
		data:  make([]float32, length),
	}
}

func NewArrayD(class *Class, length int32) *Object {
	return &Object{
		class: class,
		data:  make([]float64, length),
	}
}

func NewArrayA(class *Class, length int32) *Object {
	return &Object{
		class: class,
		data:  make([]*Object, length),
	}
}

func (o *Object) ArrayLength() int32 {
	switch o.class.name[1] { // first is '['
	case 'B', 'Z':
		return o.ArrayLengthB()
	case 'C':
		return o.ArrayLengthC()
	case 'S':
		return o.ArrayLengthS()
	case 'I':
		return o.ArrayLengthI()
	case 'J':
		return o.ArrayLengthJ()
	case 'F':
		return o.ArrayLengthF()
	case 'D':
		return o.ArrayLengthD()
	default:
		return o.ArrayLengthA()
	}
}

func (o *Object) ArrayLengthA() int32 {
	return int32(len(o.data.([]*Object)))
}

func (o *Object) ArrayLengthB() int32 {
	return int32(len(o.data.([]int8)))
}

func (o *Object) ArrayLengthC() int32 {
	return int32(len(o.data.([]uint16)))
}

func (o *Object) ArrayLengthS() int32 {
	return int32(len(o.data.([]int16)))
}

func (o *Object) ArrayLengthI() int32 {
	return int32(len(o.data.([]int32)))
}

func (o *Object) ArrayLengthJ() int32 {
	return int32(len(o.data.([]int64)))
}

func (o *Object) ArrayLengthF() int32 {
	return int32(len(o.data.([]float32)))
}

func (o *Object) ArrayLengthD() int32 {
	return int32(len(o.data.([]float64)))
}

func (o *Object) ArrGetBytes() []int8 {
	return o.data.([]int8)
}

func (o *Object) ArrGetChars() []uint16 {
	return o.data.([]uint16)
}

func (o *Object) ArrGetShorts() []int16 {
	return o.data.([]int16)
}

func (o *Object) ArrGetInts() []int32 {
	return o.data.([]int32)
}

func (o *Object) ArrGetLongs() []int64 {
	return o.data.([]int64)
}

func (o *Object) ArrGetFloats() []float32 {
	return o.data.([]float32)
}

func (o *Object) ArrGetDoubles() []float64 {
	return o.data.([]float64)
}

func (o *Object) ArrGetRefs() []*Object {
	return o.data.([]*Object)
}

// for non array object
func NewObject(class *Class) *Object {
	if class.IsArray() {
	}
	return &Object{
		class: class,
		data:  NewVars(class.InsSlotNum()),
	}
}

// for non-array object, data is Vars
// all bool, byte, char, short, int can use SetInt and GetInt methods
func (o *Object) SetInt(int int32, index uint) {
	l := o.data.(Vars)
	l[index].Num = int // notice here we do not clear the ref field
}

func (o *Object) GetInt(i uint) int32 {
	l := o.data.(Vars)
	return l[i].Num
}

func (o *Object) SetFloat(f float32, index uint) {
	l := o.data.(Vars)
	l[index].Num = int32(math.Float32bits(f)) // float is saved as int32
}

func (o *Object) GetFloat(i uint) float32 {
	l := o.data.(Vars)
	return math.Float32frombits(uint32(l[i].Num))
}

// the index i must be the lower index, big-endian
func (o *Object) SetLong(long int64, i uint) {
	l := o.data.(Vars)
	l[i].Num = int32(long >> 32)
	l[i+1].Num = int32(long)
}

func (o *Object) GetLong(i uint) int64 {
	l := o.data.(Vars)
	return int64(l[i].Num)<<32 | int64(uint32(l[i+1].Num))
}

func (o *Object) SetDouble(d float64, i uint) {
	l := o.data.(Vars)
	long := math.Float64bits(d)
	l[i].Num = int32(long >> 32)
	l[i+1].Num = int32(long)
}

func (o *Object) GetDouble(i uint) float64 {
	l := o.data.(Vars)
	high := uint64(l[i].Num) << 32
	low := uint64(uint32(l[i+1].Num))
	return math.Float64frombits(high | low)
}

// set the ref addr
func (o *Object) SetRef(ref *Object, i uint) {
	l := o.data.(Vars)
	l[i].Ref = ref
}

func (o *Object) GetRef(i uint) *Object {
	l := o.data.(Vars)
	return l[i].Ref
}

// set Slot
func (o *Object) SetSlot(s *Slot, i uint) {
	l := o.data.(Vars)
	l[i] = *s
}

// quick way to get field
func (o *Object) GetInsFieldSlotIdx(name string) uint {
	idx := o.class.GetFieldDirect(name, "").vIdx
	return idx
}

func (o *Object) GetInsFieldRef(name string) *Object {
	i := o.GetInsFieldSlotIdx(name)
	return o.GetRef(i)
}

// debug, TODO
func (o *Object) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("Type: %q addr: 0x%x ", o.class.name, unsafe.Pointer(o)))
	if o.class.IsArray() {
		elen := cmn.ElementName(o.class.name)
		if cmn.IsPrimitiveType(elen) {
			switch elen {
			case "B", "Z":
				buf.WriteString(fmt.Sprintf("%v ", o.ArrGetBytes()))
			case "I":
				buf.WriteString(fmt.Sprintf("%v ", o.ArrGetInts()))
			case "C":
				buf.WriteString(fmt.Sprintf("%v ", o.ArrGetChars()))
			case "S":
				buf.WriteString(fmt.Sprintf("%v ", o.ArrGetShorts()))
			case "J":
				buf.WriteString(fmt.Sprintf("%v ", o.ArrGetLongs()))
			case "F":
				buf.WriteString(fmt.Sprintf("%v ", o.ArrGetFloats()))
			case "D":
				buf.WriteString(fmt.Sprintf("%v ", o.ArrGetDoubles()))
			}
		} else if cmn.IsArray(elen) {
			n := o.ArrayLength()
			elements := o.ArrGetRefs()
			for i := 0; i < int(n); i++ {
				if i > 0 {
					buf.WriteByte(' ')
				}
				buf.WriteString(elements[i].String())
			}
		} else {
			buf.WriteString(fmt.Sprintf("%v ", o.ArrGetRefs()))
		}
	} else {
		if o.class.name == "java/lang/String" {
			carr := o.GetInsFieldRef("value")
			str := cmn.UTF16ToUTF8(carr.ArrGetChars())
			buf.WriteString(fmt.Sprintf("%v ", str))
		}
	}
	return buf.String()
}
