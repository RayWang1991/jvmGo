package rtdt

import (
	"fmt"
	"jvmGo/ch6/marea"
	"math"
)

// TODO debug option
// TODO, size check

type OperandStack struct {
	size uint
	marea []marea.Slot // the max depth of operands is given by compiler
}

func NewOperandStack(maxDepth uint) *OperandStack {
	return &OperandStack{
		0,
		make([]marea.Slot, maxDepth), // all marea.Slot is initiated to {0,nil}
	}
}

// Get boolean, byte, char, short, int
func (o *OperandStack) PopInt() int32 {
	o.size--
	if debugFlag {
		fmt.Printf("%d %d\n", o.size, o.marea[o.size].Num)
	}
	return o.marea[o.size].Num // panic if size is 0u-1
}

func (o *OperandStack) PushInt(i int32) {
	s := &o.marea[o.size]
	s.Num = i
	s.Ref = nil
	o.size++
	if debugFlag {
		fmt.Printf("%d %d\n", o.size, s.Num)
	}
}

func (o *OperandStack) PopRef() *marea.Object {
	o.size--
	return o.marea[o.size].Ref
}

func (o *OperandStack) PushRef(r *marea.Object) {
	o.marea[o.size].Ref = r
	o.size++
}

func (o *OperandStack) PushLong(l int64) {
	high := int32(l >> 32)
	low := int32(l)
	o.marea[o.size].Num = low
	o.marea[o.size+1].Num = high // high is on top
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	s1, s2 := &o.marea[o.size-1], &o.marea[o.size-2]
	s1.Ref = nil
	s2.Ref = nil
	o.size -= 2
	return int64(s1.Num)<<32 | int64(uint32(s2.Num))
}

func (o *OperandStack) PushFloat(f float32) {
	o.marea[o.size].Num = int32(math.Float32bits(f))
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	s := &o.marea[o.size-1]
	s.Ref = nil
	o.size--
	return math.Float32frombits(uint32(s.Num))
}

func (o *OperandStack) PushDouble(f float64) {
	n := math.Float64bits(f)
	high := int32(n >> 32)
	low := int32(n)
	o.marea[o.size].Num = low
	o.marea[o.size+1].Num = high // high is on top
	o.size += 2
}

func (o *OperandStack) PopDouble() float64 {
	s1, s2 := &o.marea[o.size-1], &o.marea[o.size-2]
	s1.Ref = nil
	s2.Ref = nil
	o.size -= 2
	return math.Float64frombits(uint64(s1.Num)<<32 | uint64(uint32(s2.Num)))
}

// return the top marea.Slot, do not pop
func (o *OperandStack) Top() *marea.Slot {
	return o.GetSlot(0)
}

func (o *OperandStack) PushSlot(s *marea.Slot) {
	o.marea[o.size] = *s
	o.size++
}

func (o *OperandStack) PopSlot() *marea.Slot {
	r := &o.marea[o.size]
	r.Ref = nil
	o.size--
	return r
}

// i is the index to the top of stack, just get, not pop
func (o *OperandStack) GetSlot(i uint) *marea.Slot {
	if i >= uint(o.size) {
		panic("StackOutOfRange")
	}
	return &o.marea[o.size-i-1]
}
