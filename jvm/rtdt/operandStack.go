package rtdt

import (
	"jvmGo/jvm/marea"
	"jvmGo/jvm/utils"
	"math"
	"bytes"
	"fmt"
)

// TODO debug option
// TODO, size check

type OperandStack struct {
	size  uint
	slots []marea.Slot // the max depth of operands is given by compiler
}

func NewOperandStack(maxDepth uint) *OperandStack {
	return &OperandStack{
		0,
		make([]marea.Slot, maxDepth), // all slots.Slot is initiated to {0,nil}
	}
}

// Get boolean, byte, char, short, int
func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].Num // panic if size is 0u-1
}

func (o *OperandStack) PushInt(i int32) {
	s := &o.slots[o.size]
	s.Num = i
	s.Ref = nil
	o.size++
}

func (o *OperandStack) PopRef() *marea.Object {
	o.size--
	return o.slots[o.size].Ref
}

func (o *OperandStack) PopNonnilRef() *marea.Object {
	obj := o.PopRef() // instance object
	if obj == nil {
		panic(utils.NullPointerException)
	}
	return obj
}

func (o *OperandStack) PushRef(r *marea.Object) {
	o.slots[o.size].Ref = r
	o.size++
}

func (o *OperandStack) PushLong(l int64) {
	high := int32(l >> 32)
	low := int32(l)
	o.slots[o.size].Num = low
	o.slots[o.size+1].Num = high // high is on top
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	s1, s2 := &o.slots[o.size-1], &o.slots[o.size-2]
	s1.Ref = nil
	s2.Ref = nil
	o.size -= 2
	return int64(s1.Num)<<32 | int64(uint32(s2.Num))
}

func (o *OperandStack) PushFloat(f float32) {
	o.slots[o.size].Num = int32(math.Float32bits(f))
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	s := &o.slots[o.size-1]
	s.Ref = nil
	o.size--
	return math.Float32frombits(uint32(s.Num))
}

func (o *OperandStack) PushDouble(f float64) {
	n := math.Float64bits(f)
	high := int32(n >> 32)
	low := int32(n)
	o.slots[o.size].Num = low
	o.slots[o.size+1].Num = high // high is on top
	o.size += 2
}

func (o *OperandStack) PopDouble() float64 {
	s1, s2 := &o.slots[o.size-1], &o.slots[o.size-2]
	s1.Ref = nil
	s2.Ref = nil
	o.size -= 2
	return math.Float64frombits(uint64(s1.Num)<<32 | uint64(uint32(s2.Num)))
}

// return the top slots.Slot, do not pop
func (o *OperandStack) Top() *marea.Slot {
	return o.GetSlot(0)
}

func (o *OperandStack) PushSlot(s *marea.Slot) {
	o.slots[o.size] = *s
	o.size++
}

func (o *OperandStack) PopSlot() *marea.Slot {
	copy := o.slots[o.size-1] // copy slot
	o.slots[o.size-1].Ref = nil
	o.size--
	return &copy
}

// i is the index to the top of stack, just get, not pop
func (o *OperandStack) GetSlot(i uint) *marea.Slot {
	if i >= uint(o.size) {
		panic("OperandStackOutOfRange")
	}
	return &o.slots[o.size-i-1]
}

func (o *OperandStack) SetSlot(s *marea.Slot, i uint) {
	o.slots[i] = *s
}

func (o *OperandStack) GetSize() uint {
	return o.size
}


func (l *OperandStack) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("Size: %d",l.size))
	buf.WriteByte('[')
	for i, v := range l.slots {
		if i > 0 {
			buf.WriteByte(',')
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