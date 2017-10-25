package marea

import (
	"fmt"
	"jvmGo/ch6/classfile"
	"jvmGo/ch6/cmn"
	"strings"
)

type Class struct {
	name           string
	flags          uint16
	cp             ConstantPool
	superClassName string // string as K
	superClass     *Class
	interfaceNames []string // string as K
	interfaces     []*Class
	hasInited      bool

	initLoader ClassLoader
	defLoader  ClassLoader

	methodMap map[string]*Method // map with K, name & desc
	fieldMap  map[string]*Field  // map with K, name

	/*
		insFld2Inx map[string]int    // field name to slots, instance use
		stcFld2Inx map[string]int    // field name to slots, static use
	*/

	insSlotN uint // instance slot num
	Vars          // static vars, should be initiated with constant attribute

	// TODO items
	//innerClasses    []*FromClass
	//enclosingMethod *Method
	//bootStrapMethods BootstrapMethods
	//sourceFile string
}

// TODO
func NewClass(file *classfile.ClassFile) *Class {
	c := &Class{}
	c.flags = file.AccessFlag()
	c.name = file.ClassName()
	c.cp = NewConstantPool(file.ConstantPool(), c)
	c.superClassName = file.SuperClassName()
	//leave super from to loader
	c.interfaceNames = file.InterfaceNames()

	ifs := file.InstFields() // instance fields
	sfs := file.StatFields() // static fields

	fmap := make(map[string]*Field, len(sfs)+len(ifs))
	c.fieldMap = fmap

	// static fields
	vs := NewVars(file.StatSlotN())
	c.Vars = vs
	vi := uint(0)
	for _, f := range sfs {
		fd := NewField(c, f)
		fd.vIdx = vi
		vi += uint(fd.sn)

		// set default value if needed
		if ind := f.GetConstantValueIndex(); ind >= 0 {
			c.SetStatField(fd, uint16(ind))
		}
		fmap[fd.name] = fd // TODO, for java language specification only
	}

	c.insSlotN = uint(len(ifs))
	vi = 0
	for _, f := range ifs {
		fd := NewField(c, f)
		vi += uint(fd.sn)
		fmap[fd.name] = fd // TODO, for jls only, may use name and desc to unify a field
	}

	ms := file.MethodInfo()
	mmap := make(map[string]*Method, len(ms))
	c.methodMap = mmap
	//dup name+desc should not be valid
	for _, m := range ms {
		md := NewMethod(c, m)
		mmap[ndStr(md.name, md.desc)] = md
	}

	return c
}

// functional
func (c *Class) SetStatField(f *Field, i uint16) {
	switch f.Desc() {
	case "B", "C", "I", "S", "Z":
		v := c.cp.GetInteger(i)
		c.SetInt(v, f.vIdx)
	case "D":
		v := c.cp.GetDouble(i)
		c.SetDouble(v, f.vIdx)
	case "J":
		v := c.cp.GetLong(i)
		c.SetLong(v, f.vIdx)
	case "F":
		v := c.cp.GetFloat(i)
		c.SetFloat(v, f.vIdx)
	default:
		//  unsupported now
	}
}

// wrappers for look up field
func (c *Class) InstField(name string) *Field {
	f := c.fieldMap[name]
	if f != nil && f.IsStatic() {
		panic("not instance field")
	}
	return f
}

func (c *Class) StatField(name string) *Field {
	f := c.fieldMap[name]
	if f != nil && !f.IsStatic() {
		panic("not static field")
	}
	return f
}

func (c *Class) Method(name, desc string) *Method {
	s := ndStr(name, desc)
	return c.methodMap[s]
}

// getters
func (c *Class) ClassName() string {
	return c.name
}

func (c *Class) SuperclassName() string {
	return c.superClassName
}

func (c *Class) Superclass() *Class {
	return c.superClass
}

func (c *Class) FlagString() string {
	return cmn.FlagNumToString(c.flags, cmn.ACC_TYPE_CLASS)
}

func (c *Class) InterfaceNames() []string {
	return c.interfaceNames
}

func (c *Class) Interfaces() []*Class {
	return c.interfaces
}

func (c *Class) ConstantPool() ConstantPool {
	return c.cp
}

// class loader

func (c *Class) InitLoader() ClassLoader {
	return c.initLoader
}

func (c *Class) SetInitLoader(l ClassLoader) {
	c.initLoader = l
}

func (c *Class) DefineLoader() ClassLoader {
	return c.defLoader
}

func (c *Class) SetDefineLoader(l ClassLoader) {
	c.defLoader = l
}

// field & method
func (c *Class) FieldMap() map[string]*Field {
	return c.fieldMap
}

func (c *Class) MethodMap() map[string]*Method {
	return c.methodMap
}

func (c *Class) InsSlotNum() uint {
	return c.insSlotN
}

// getter for method
func (c *Class) GetMethodDirect(name, desc string) *Method {
	return c.methodMap[ndStr(name, desc)]
}

func (c *Class) GetFieldDirect(name, desc string) *Field {
	return c.fieldMap[ndStr(name, desc)]
}

// look up field recursively
func (c *Class) LookUpField(name string) *Field {
	if f := c.fieldMap[name]; f != nil {
		return f
	}
	if len(c.interfaces) > 0 {
		for _, infc := range c.interfaces {
			if f := infc.LookUpField(name); f != nil {
				return f
			}
		}
	}
	if c.superClass != nil {
		return c.superClass.LookUpField(name)
	}
	return nil
}

// look up method
func (c *Class) LookUpMethod(name, desc string) *Method {
	k := ndStr(name, desc)
	if m := c.LookUpMethodInClass(k); m != nil {
		return m
	}
	return c.LookUpMethodInInterface(k)
}

// look up method in class hierarchy
func (c *Class) LookUpMethodInClass(key string) *Method {
	for ; c != nil; c = c.superClass {
		if m := c.methodMap[key]; m != nil {
			return m
		}
	}
	return nil
}

// look up method in interfaces
func (c *Class) LookUpMethodInInterface(key string) *Method {
	for _, in := range c.interfaces {
		if m := in.methodMap[key]; m != nil && !(m.IsAbstract() && m.IsPrivate() && m.IsStatic()) {
			return m
		}
	}
	for _, in := range c.interfaces {
		if m := in.LookUpMethodInInterface(key); m != nil {
			return m
		}
	}
	return nil
}

func (c *Class) LookUpMethodDirectly(name, desc string) *Method {
	return c.methodMap[ndStr(name, desc)]
}

// setters for static fields

// return the package name
func (c *Class) PackageName() string {
	n := c.name
	i := strings.LastIndex(n, `/`)
	if i > 0 {
		return n[0:i]
	} else {
		return "" // default package
	}
}

// array
func (c *Class) IsArray() bool {
	return cmn.IsArray(c.name)
}

// init
func (c *Class) hasInitiated() bool {
	return c.hasInited
}

func (c *Class) setInitiated(b bool) {
	c.hasInited = b
}

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
	// treat superclass methods specially when invoked by invokespecial,
	// changes from jdk 1.2 ~ 1.1
	return cmn.IsSuper(c.flags)
}

// debug
func (c *Class) PrintDebugMessage() {
	fmt.Printf("class: %s\n", c.ClassName())      // magic
	fmt.Printf("super: %s\n", c.SuperclassName()) // magic
	fmt.Printf("flags: %s\n", cmn.FlagNumToString(c.flags, cmn.ACC_TYPE_CLASS))
	//fmt.Print(c.cp.String())
	fmt.Printf("interfaces(%d items): %s \n", len(c.interfaceNames), strings.Join(c.interfaceNames, ","))
	// Fields and Methods
	fmt.Printf("Fields (%d items):\n", len(c.fieldMap))
	i := 0
	for _, f := range c.fieldMap {
		fmt.Printf("#%d: %s %s\n, %s\n", i, f.name, f.desc, cmn.FlagNumToString(f.flags, cmn.ACC_TYPE_FIELD))
		i++
	}
	fmt.Printf("Methods(%d items):\n", len(c.methodMap))
	i = 0
	for _, m := range c.methodMap {
		fmt.Printf("#%d: %s %s\n, %s\n", i, m.name, m.desc, cmn.FlagNumToString(m.flags, cmn.ACC_TYPE_METHOD))
		if m.IsNative() {
			fmt.Println("Native Method")
		} else {
			// TODO
			//fmt.Printf("%s",classfile.CodeInst(m.Code()).String())
		}
	}
}
