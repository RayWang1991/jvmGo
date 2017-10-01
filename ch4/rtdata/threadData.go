package rtdata

import "jvmGo/ch4/errcode"

// Thread represents a thread
type Thread struct {
	pc    int32 // pc register, TODO
	stack Stack
}

func (t *Thread) PC() int32 {
	return t.pc
}

func (t *Thread) SetPC(p int32) {
	t.pc = p
}

func (t *Thread) PushFrame(f *Frame) {
	t.stack.push(f)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top
}

// Stack represents a program stack for a thread
type Stack struct {
	maxSize     uint32 // max frame num
	currentSize uint32 // current frame num
	top         *Frame // top frame
}

func (s *Stack) push(f *Frame) {
	s.currentSize ++
	if s.currentSize > s.maxSize {
		panic(errcode.StackOverFlow)
	}
	f.next = s.top
	s.top = f
}

func (s *Stack) pop() *Frame {
	if s.top == nil {
		panic("empty stack, can not pop")
	}
	s.currentSize--
	r := s.top
	r.next = nil // for gc
	s.top = s.top.next
	return r
}
