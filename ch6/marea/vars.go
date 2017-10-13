package marea

import (
	"math"
)

type Vars []Slot // fixed-length array

func NewLocalVars(l uint) Vars {
	return make([]Slot, l)
}

// all bool, byte, char, short, int can use SetInt and GetInt methods
func (l Vars) SetInt(int int32, index uint) {
	l[index].Num = int // notice here we do not clear the ref field
}

func (l Vars) GetInt(i uint) int32 {
	return l[i].Num
}

func (l Vars) SetFloat(f float32, index uint) {
	l[index].Num = int32(math.Float32bits(f)) // float is saved as int32
}

func (l Vars) GetFloat(i uint) float32 {
	return math.Float32frombits(uint32(l[i].Num))
}

// the index i must be the lower index, big-endian
func (l Vars) SetLong(long int64, i uint) {
	l[i].Num = int32(long >> 32)
	l[i+1].Num = int32(long)
}

func (l Vars) GetLong(i uint) int64 {
	return int64(l[i].Num)<<32 | int64(uint32(l[i+1].Num))
}

func (l Vars) SetDouble(d float64, i uint) {
	long := math.Float64bits(d)
	l[i].Num = int32(long >> 32)
	l[i+1].Num = int32(long)
}

func (l Vars) GetDouble(i uint) float64 {
	high := uint64(l[i].Num) << 32
	low := uint64(uint32(l[i+1].Num))
	return math.Float64frombits(high | low)
}

// set the ref addr
func (l Vars) SetRef(ref *Object, i uint) {
	l[i].Ref = ref
}

func (l Vars) GetRef(i uint) *Object {
	return l[i].Ref
}
