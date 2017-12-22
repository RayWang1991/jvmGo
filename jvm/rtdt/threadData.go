package rtdt

import (
	"jvmGo/jvm/utils"
	"fmt"
)

// Thread represents a thread
type Thread struct {
	pc    int32 // pc register for current executing code
	stack Stack
}

// fixed stack
func NewThread(maxDep uint32) *Thread {
	return &Thread{
		pc: 0,
		stack: Stack{
			maxSize:     maxDep,
			currentSize: 0,
			top:         nil}}
}

// PC

func (t *Thread) PC() int32 {
	return t.pc
}

func (t *Thread) SetPC(p int32) {
	t.pc = p
}

// Frame

func (t *Thread) PushFrame(f *Frame) {
	t.stack.push(f)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top
}

func (t *Thread) GetFrameSize() uint32 {
	return t.stack.currentSize
}

func (t *Thread) RollBackPCIfNeeded() {
	f := t.CurrentFrame()
	if nil == f {
		return
	}
	f.SetPC(t.PC())
}

// Stack represents a program stack for a thread
type Stack struct {
	maxSize     uint32 // max frame num
	currentSize uint32 // current frame num
	top         *Frame // top frame
}

func (s *Stack) push(f *Frame) {
	s.currentSize++
	if s.currentSize > s.maxSize {
		panic(utils.StackOverFlow)
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
	s.top = s.top.next
	return r
}

// run loop
func (t *Thread) PrintStack() {
	f := t.CurrentFrame()
	fmt.Printf("CALL STACK:\n")
	for c := f; c != nil; c = c.GetNext() {
		fmt.Printf("%s.%s()\n", c.Method().Class().ClassName(), c.Method().Name())
	}
}
