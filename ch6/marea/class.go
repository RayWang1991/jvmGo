package marea

import (
	"jvmGo/ch6/classfile"
	"jvmGo/ch6/cmn"
)

type Class struct {
	name           string
	flags          uint16
	cp             ConstantPool
	superClassName string // string as K
	superClass     *Class
	interfaceNames []string // string as K
	interfaces     []*Class
	loaderId       int
	fields         []*Field  // TODO map with K, name
	methods        []*Method // TODO map with K, name + descriptor?

	methodMap map[string]*Method
	fieldMap  map[string]*Field // filed name to Field information

	/*
	insFld2Inx map[string]int    // field name to slots, instance use
	stcFld2Inx map[string]int    // field name to slots, static use
	*/

	insSlotN    int  // instance slot num
	staticSlots Vars // should be initiated with constant attribute

	// TODO items
	//innerClasses    []*Class
	//enclosingMethod *Method
	//bootStrapMethods BootstrapMethods
	//sourceFile string
}

// TODO
func NewClass(file *classfile.ClassFile) *Class {
	c := &Class{}
	c.flags = file.AccessFlag()
	c.name = file.ClassName()
	c.cp = NewConstantPool(file.ConstantPool())
	c.superClassName = file.SuperClassName()
	//leave super class to loader
	c.interfaceNames = file.InterfaceNames()

	fs := file.FieldInfo()
	c.fields = make([]*Field, len(fs))
	//dup name should not be valid
	for i, f := range fs {
		c.fields[i] = NewField(f)
	}

	ms := file.MethodInfo()
	c.methods = make([]*Method, len(ms))
	//dup name+desc should not be valid
	for i,m := range ms{
		c.methods[i] = New
	}
	return nil
}

// TODO
// getters

// TODO
// setters

// access methods
func (c *Class) IsPublic() bool {
	return cmn.IsPublic(c.flags)
}

func (c *Class) IsFinal() bool {
	return cmn.IsFinal(c.flags)
}

func (c *Class) IsInterface() bool {
	return cmn.IsInterface(c.flags)
}

func (c *Class) IsAbstract() bool {
	return cmn.IsAbstract(c.flags)
}

func (c *Class) IsSuper() bool {
	return cmn.IsSuper(c.flags) // treat superclass methods specially when invoked by invokespecial
}
