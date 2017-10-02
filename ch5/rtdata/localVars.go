package rtdata

import "math"

type LocalVars []slot // fixed-length array

func NewLocalVars(l uint) LocalVars {
	return make([]slot, l)
}

// all bool, byte, char, short, int can use SetInt and GetInt methods
func (l LocalVars) SetInt(int int32, index uint) {
	l[index].num = int // notice here we do not clear the ref field
}

func (l LocalVars) GetInt(i uint) int32 {
	return l[i].num
}

func (l LocalVars) SetFloat(f float32, index uint) {
	l[index].num = int32(math.Float32bits(f)) // float is saved as int32
}

func (l LocalVars) GetFloat(i uint) float32 {
	return math.Float32frombits(uint32(l[i].num))
}

// the index i must be the lower index, big-endian
func (l LocalVars) SetLong(long int64, i uint) {
	l[i].num = int32(long >> 32)
	l[i+1].num = int32(long)
}

func (l LocalVars) GetLong(i uint) int64 {
	return int64(l[i].num)<<32 | int64(uint32(l[i+1].num))
}

func (l LocalVars) SetDouble(d float64, i uint) {
	long := math.Float64bits(d)
	l[i].num = int32(long >> 32)
	l[i+1].num = int32(long)
}

func (l LocalVars) GetDouble(i uint) float64 {
	high := uint64(l[i].num) << 32
	low := uint64(uint32(l[i+1].num))
	return math.Float64frombits(high | low)
}

// set the ref addr
func (l LocalVars) SetRef(ref *Object, i uint) {
	l[i].ref = ref
}

func (l LocalVars) GetRef(i uint) *Object {
	return l[i].ref
}
