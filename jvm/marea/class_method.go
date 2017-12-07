package marea

import (
	"fmt"
	"jvmGo/jvm/classfile"
	"jvmGo/jvm/cmn"
)

type Method struct {
	ClassMember
	maxStackDep uint16
	maxLocalVar uint16
	argDs       []string
	argSlotN    int
	retD        string
	code        []byte
	exceptions  []*AttrExceptionEntry //strings of exception from name
}

type AttrExceptionEntry struct {
	startPC   uint16 // startPC and endPC specified the range of the exception handler code
	endPC     uint16 // index into code array [startPC, endPC) historical bug for JVM designer
	handlerPC uint16 // index into code array, indicating where to start the handler
	catchType string
}

func NewMethod(from *Class, info *classfile.MethodInfo) *Method {
	if info.Name() == "write" {
		fmt.Println()
	}
	m := &Method{}
	m.class = from
	m.name = info.Name()
	m.flags = info.AccFlags()
	m.desc = info.Description()
	m.parseDesc()
	if m.IsNative() || m.IsAbstract() {
		// TODO native methods ,and abstract methods do not have code attr but may have exception table
		return m
	}

	codeAttr := info.GetCodeAttr()
	m.maxLocalVar = codeAttr.MaxLocals()
	m.maxStackDep = codeAttr.MaxStack()
	m.code = codeAttr.Code()
	excs := codeAttr.ExceptionTable()
	m.exceptions = make([]*AttrExceptionEntry, len(excs))
	for i, e := range excs {
		cName := e.GetCatchType(info.ConstantPool())
		entry := &AttrExceptionEntry{
			e.StartPC(),
			e.EndPC(),
			e.HandlerPC(),
			cName}
		m.exceptions[i] = entry
	}
	return m
}

func (m *Method) parseDesc() {
	args := []string{}
	type descState int
	const (
		START  = iota
		NEEDBR
	)
	inArg := true
	state := START
	res := 0
	lst := 1 // skip '('
	for ptr, c := range m.desc[1:] { // must be ascii bytes,skip '('
		if c == ')' {
			inArg = false
		}
		switch state {
		case START:
			if c == 'L' || c == '[' {
				lst = ptr
				state = NEEDBR
			} else {
				d := string(c)
				if inArg {
					args = append(args, d)
					if c == 'D' || c == 'L' {
						res += 2
					} else {
						res++
					}
				} else {
					m.retD = d
				}
			}
		case NEEDBR:
			if c == ';' {
				d := m.desc[lst:ptr]
				ptr++
				if inArg {
					res++
					args = append(args, d)
				} else {
					m.retD = d
				}
				state = START
			}
		}
	}
	m.argDs = args
	m.argSlotN = res
}

// getters
func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) MaxStackDep() uint16 {
	return m.maxStackDep
}

// for native methods
func (m *Method) SetMaxLocalVars(n uint16) {
	m.maxLocalVar = n
}

func (m *Method) MaxLocalVars() uint16 {
	return m.maxLocalVar
}

func (m *Method) ArgSlotNum() int {
	return m.argSlotN
}

func (m *Method) RetD() string {
	return m.retD
}

func (m *Method) ArgDs() []string {
	return m.argDs
}

// access methods
func (m *Method) IsPublic() bool {
	return cmn.IsPublic(m.flags)
}

func (m *Method) IsPrivate() bool {
	return cmn.IsPrivate(m.flags)
}

func (m *Method) IsProtected() bool {
	return cmn.IsProtected(m.flags)
}

func (m *Method) IsFinal() bool {
	return cmn.IsFinal(m.flags)
}

func (m *Method) IsStatic() bool {
	return cmn.IsStatic(m.flags)
}

func (m *Method) IsSynchronized() bool {
	return cmn.IsSynchronized(m.flags)
}

func (m *Method) IsVarargs() bool {
	return cmn.IsVarargs(m.flags)
}

func (m *Method) IsNative() bool {
	return cmn.IsNative(m.flags)
}

func (m *Method) IsAbstract() bool {
	return cmn.IsAbstract(m.flags)
}
