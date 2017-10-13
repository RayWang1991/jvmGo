package rtdata

import (
	"math"
	"fmt"
	"jvmGo/ch6/cmn"
)

// TODO debug option
// TODO, size check

type OperandStack struct {
	size uint
	data []cmn.Slot // the max depth of operands is given by compiler
}

func NewOperandStack(maxDepth uint) *OperandStack {
	return &OperandStack{
		0,
		make([]cmn.Slot, maxDepth), // all cmn.Slot is initiated to {0,nil}
	}
}

// Get boolean, byte, char, short, int
func (o *OperandStack) PopInt() int32 {
	o.size--
	if debugFlag {
		fmt.Printf("%d %d\n", o.size, o.data[o.size].Num)
	}
	return o.data[o.size].Num // panic if size is 0u-1
}

func (o *OperandStack) PushInt(i int32) {
	s := &o.data[o.size]
	s.Num = i
	s.Ref = nil
	o.size++
	if debugFlag {
		fmt.Printf("%d %d\n", o.size, s.Num)
	}
}

func (o *OperandStack) PopRef() *cmn.Object {
	o.size--
	return o.data[o.size].Ref
}

func (o *OperandStack) PushRef(r *cmn.Object) {
	o.data[o.size].Ref = r
	o.size++
}

func (o *OperandStack) PushLong(l int64) {
	high := int32(l >> 32)
	low := int32(l)
	o.data[o.size].Num = low
	o.data[o.size+1].Num = high // high is on top
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	s1, s2 := &o.data[o.size-1], &o.data[o.size-2]
	s1.Ref = nil
	s2.Ref = nil
	o.size -= 2
	return int64(s1.Num)<<32 | int64(uint32(s2.Num))
}

func (o *OperandStack) PushFloat(f float32) {
	o.data[o.size].Num = int32(math.Float32bits(f))
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	s := &o.data[o.size-1]
	s.Ref = nil
	o.size--
	return math.Float32frombits(uint32(s.Num))
}

func (o *OperandStack) PushDouble(f float64) {
	n := math.Float64bits(f)
	high := int32(n >> 32)
	low := int32(n)
	o.data[o.size].Num = low
	o.data[o.size+1].Num = high // high is on top
	o.size += 2
}

func (o *OperandStack) PopDouble() float64 {
	s1, s2 := &o.data[o.size-1], &o.data[o.size-2]
	s1.Ref = nil
	s2.Ref = nil
	o.size -= 2
	return math.Float64frombits(uint64(s1.Num)<<32 | uint64(uint32(s2.Num)))
}

// return the top cmn.Slot, do not pop
func (o *OperandStack) Top() *cmn.Slot {
	return o.GetSlot(0)
}

func (o *OperandStack) PushSlot(s *cmn.Slot) {
	o.data[o.size] = *s
	o.size++
}

func (o *OperandStack) PopSlot() *cmn.Slot {
	r := &o.data[o.size]
	r.Ref = nil
	o.size--
	return r
}

// i is the index to the top of stack, just get, not pop
func (o *OperandStack) GetSlot(i uint) *cmn.Slot {
	if i >= uint(o.size) {
		panic("StackOutOfRange")
	}
	return &o.data[o.size-i-1]
}
