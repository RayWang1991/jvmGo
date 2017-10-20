package marea

import (
	"jvmGo/ch6/classfile"
	"jvmGo/ch6/cmn"
)

type Method struct {
	ClassMember
	maxStackDep uint16
	maxLocalVar uint16
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
	if m.IsNative() {
		// TODO native methods
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
