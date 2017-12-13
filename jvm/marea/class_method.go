package marea

import (
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

	n := len(m.desc)
loop:
	for ptr := 1; ptr < n; ptr++ {
		c := m.desc[ptr]
		if c == ')' {
			inArg = false
		}
		switch state {
		case START:
			if inArg {
				if c == 'L' { // ref
					lst = ptr
					state = NEEDBR
				} else if c == '[' { // array
					lst = ptr - 1
					ptr++
					for c = m.desc[ptr]; c == '['; ptr++ {
					}
					if c != 'L' { // base
						d := m.desc[lst:ptr+1]
						args = append(args, d)
						res++
					} else {
						state = NEEDBR
					}
				} else { // base
					d := string(c)
					args = append(args, d)
					if c == 'D' || c == 'J' {
						res += 2
					} else {
						res++
					}
				}
			} else {
				m.retD = m.desc[ptr+1:]
				if m.retD != "" && m.retD[len(m.retD)-1] == ';' {
					m.retD = m.retD[:len(m.retD)-1]
				}
				break loop
			}
		case NEEDBR:
			if c == ';' {
				d := m.desc[lst:ptr] // delete;
				//ptr++
				if inArg {
					res++
					args = append(args, d)
				} else {
					m.retD = d
					break loop
				}
				state = START
			}
		}
	}
	m.argDs = args
	m.argSlotN = res
}

// for native methods
func (m *Method) SetMaxLocalVars(n uint16) {
	m.maxLocalVar = n
}

func (m *Method) SetMaxStackDep(n uint16) {
	m.maxStackDep = n
}

func (m *Method) SetCode(byteCodes []byte) {
	m.code = byteCodes
}

// getters
func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) MaxStackDep() uint16 {
	return m.maxStackDep
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

func HackMethod(class *Class, flags uint16, name, desc string, code []byte) *Method {
	m := &Method{
		ClassMember: ClassMember{
			class: class,
			name:  name,
			desc:  desc,
			flags: flags,
		},
		code: code,
	}
	m.parseDesc()
	return m
}
