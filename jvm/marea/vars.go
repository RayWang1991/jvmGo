package marea

import (
	"math"
	"bytes"
	"fmt"
)

type Vars []Slot // fixed-length array

func NewVars(l uint) Vars {
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

func (l Vars) SetLong(long int64, i uint) {
	l[i+1].Num = int32(long >> 32) // high, i+1
	l[i].Num = int32(long)         // low, i
}

func (l Vars) GetLong(i uint) int64 {
	low := uint32(l[i].Num)
	high := uint32(l[i+1].Num)
	return int64(high)<<32 | int64(low)
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
	if int(i) >= len(l) {
		fmt.Printf("len %d i %d", len(l), i)
	}
	return l[i].Ref
}

// set Slot
func (l Vars) SetSlot(s *Slot, i uint) {
	l[i] = *s
}

func (l Vars) String() string {
	buf := bytes.Buffer{}
	buf.WriteByte('[')
	for i, v := range l {
		if i > 0 {
			buf.WriteByte(' ')
		}
		if v.Ref != nil {
			buf.WriteString(v.Ref.String())
		} else {
			buf.WriteString(fmt.Sprintf("%d", v.Num))
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
