package rtdata

import (
	"fmt"
	"math"
)

// TODO debug option
// TODO, size check

type OperandStack struct {
	size uint
	data []slot // the max depth of operands is given by compiler
}

func NewOperandStack(maxDepth uint) *OperandStack {
	return &OperandStack{
		0,
		make([]slot, maxDepth), // all slot is initiated to {0,nil}
	}
}

// Get boolean, byte, char, short, int
func (o *OperandStack) PopInt() int32 {
	o.size--
	if debugFlag {
		fmt.Printf("%d %d\n", o.size, o.data[o.size].num)
	}
	return o.data[o.size].num // panic if size is 0u-1
}

func (o *OperandStack) PushInt(i int32) {
	s := &o.data[o.size]
	s.num = i
	s.ref = nil
	o.size++
	if debugFlag {
		fmt.Printf("%d %d\n", o.size, s.num)
	}
}

func (o *OperandStack) PopRef() *Object {
	o.size--
	return o.data[o.size].ref
}

func (o *OperandStack) PushRef(r *Object) {
	o.data[o.size].ref = r
	o.size++
}

func (o *OperandStack) PushLong(l int64) {
	high := int32(l >> 32)
	low := int32(l)
	o.data[o.size].num = low
	o.data[o.size+1].num = high // high is on top
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	s1, s2 := &o.data[o.size-1], &o.data[o.size-2]
	s1.ref = nil
	s2.ref = nil
	o.size -= 2
	return int64(s1.num)<<32 | int64(uint32(s2.num))
}

func (o *OperandStack) PushFloat(f float32) {
	o.data[o.size].num = int32(math.Float32bits(f))
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	s := &o.data[o.size-1]
	s.ref = nil
	o.size--
	return math.Float32frombits(uint32(s.num))
}

func (o *OperandStack) PushDouble(f float64) {
	n := math.Float64bits(f)
	high := int32(n >> 32)
	low := int32(n)
	o.data[o.size].num = low
	o.data[o.size+1].num = high // high is on top
	o.size += 2
}

func (o *OperandStack) PopDouble() float64 {
	s1, s2 := &o.data[o.size-1], &o.data[o.size-2]
	s1.ref = nil
	s2.ref = nil
	o.size -= 2
	return math.Float64frombits(uint64(s1.num)<<32 | uint64(uint32(s2.num)))
}

// return the top slot, do not pop
func (o *OperandStack) Top() *slot {
	return o.GetSlot(0)
}

func (o *OperandStack) PushSlot(s *slot) {
	o.data[o.size] = *s
	o.size++
}

func (o *OperandStack) PopSlot() *slot {
	r := &o.data[o.size]
	r.ref = nil
	o.size--
	return r
}

// i is the index to the top of stack, just get, not pop
func (o *OperandStack) GetSlot(i uint) *slot {
	if i >= uint(o.size) {
		panic("StackOutOfRange")
	}
	return &o.data[o.size-i-1]
}
