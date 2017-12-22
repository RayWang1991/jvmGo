package cloader

import (
	"fmt"
	"jvmGo/jvm/classfile"
	"jvmGo/jvm/classpath"
	"jvmGo/jvm/cmn"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/utils"
	"os"
	"strings"
)

var __bstLoader *bstLoader

func BSTLoader() *bstLoader {
	return __bstLoader
}

func NewBstLoader(cp *classpath.ClassPath) *bstLoader {
	__bstLoader := &bstLoader{
		id: marea.BootstrapClassLoaderId,
		cp: cp,
	}
	return __bstLoader
}

func (b *bstLoader) SetUpBase() {
	b.Load(utils.CLASSNAME_Class)
	b.loadPrimitiveClasses()
	for _, c := range cache {
		setClzObj(c)
	}
	vm := b.Load(utils.CLASSNAME_VM)
	b.Initiate(vm)
}

func (b *bstLoader) loadPrimitiveClasses() {
	names := []string{
		utils.CLASSNAME_prim_boolean,
		utils.CLASSNAME_prim_byte,
		utils.CLASSNAME_prim_char,
		utils.CLASSNAME_prim_short,
		utils.CLASSNAME_prim_int,
		utils.CLASSNAME_prim_long,
		utils.CLASSNAME_prim_float,
		utils.CLASSNAME_prim_double,
		utils.CLASSNAME_prim_void,
	}
	for _, name0 := range names {
		c := &marea.Class{
		}
		c.SetClassName(name0)
		c.SetDefineLoader(b)
		c.SetInitLoader(b)
		c.SetInitiated(true)
		cache[name0] = c
		setClzObj(c)
	}
}

// private methods
// set class object for class
func setClzObj(c *marea.Class) {
	clsClass := cache[utils.CLASSNAME_Class]
	if clsClass != nil {
		clsObj := marea.NewObject(clsClass)
		clsObj.SetClzClass(c)
		c.SetClassObject(clsObj)
		utils.DLoaderPrintf("[CLASS] set clz obj for %s\n", c.ClassName())
	}
}

// adjust fields index basing on inheritance hierarchy
func adjustFields(c *marea.Class) {
	// find the highest class having been initiated or having no super class(Object)
	if c.HasAdjustedSlots() {
		return
	}
	defer c.SetAdjustedSlots(true)
	if c.Superclass() == nil {
		return
	}
	if !c.Superclass().HasAdjustedSlots() {
		adjustFields(c.Superclass())
	}
	istn := c.Superclass().InsSlotNum()
	c.SetInsSlotNum(c.InsSlotNum() + istn)
	// the static field is not inherited in data structure, it's compiler who redirect it from lower to higher
	for _, f := range c.FieldMap() {
		if !f.IsStatic() {
			f.SetVarIdx(f.VarIdx() + istn)
		}
	}
}

// not concurrently safe
var cache map[string]*marea.Class = make(map[string]*marea.Class) // class full name : class

type bstLoader struct {
	id int
	cp *classpath.ClassPath
}

func (b *bstLoader) ID() int {
	return b.id
}

func (b *bstLoader) Delegate() marea.ClassLoader {
	return nil
}

func (b *bstLoader) Load(n string) *marea.Class {
	//utils.DLoaderPrintf("Load %s\n", n)
	if cmn.IsArray(n) {
		return b.loadArrayClass(n)
	} else {
		return b.loadNormalClass(n)
	}
}

func (b *bstLoader) loadNormalClass(n string) *marea.Class {
	//utils.DLoaderPrintf("Initate %s\n", n)
	if c := cache[n]; c != nil {
		if c.InitLoader().ID() == b.id {
			return c
		} else {
			panic(utils.LinkageError)
		}
	} else {
		c := b.Define(n)
		b.Verify(c)
		b.Prepare(c)
		return c
	}
}

func (b *bstLoader) _loadClassDirect(n string) *marea.Class {
	cf, err := b.doLoadClassFile(n, b.cp)
	utils.DLoaderPrintf("load Class direct %s\n", n)

	if cf == nil {
		panic(utils.ClassNotFoundException)
	}
	if err != nil {
		panic(err)
	}
	c := b.doLoadClassFromFile(cf)
	if c == nil {
		panic(utils.ClassFormatError)
	}
	c.SetInitLoader(b)
	c.SetDefineLoader(b)
	cache[n] = c
	return c
}

func (b *bstLoader) Define(n string) *marea.Class {
	utils.DLoaderPrintf("define Class %s\n", n)

	c := b._loadClassDirect(n)
	t := c

	// load super class, not initiated yet
	for t.SuperclassName() != "" {
		s := cache[t.SuperclassName()]
		if s == nil {
			s = b._loadClassDirect(t.SuperclassName())
			setClzObj(s)
			b.setUpInterfaces(s)
		}
		t.SetSuperClass(s)
		t = s
	}

	setClzObj(c)
	adjustFields(c)
	b.setUpInterfaces(c)

	//b.doInitClass(c)
	return c
}

func (b *bstLoader) setUpInterfaces(c *marea.Class) {
	if len(c.InterfaceNames()) > 0 {
		intfs := make([]*marea.Class, 0, len(c.InterfaceNames()))
		for _, itfName := range c.InterfaceNames() {
			itf := b.Load(itfName)
			intfs = append(intfs, itf)
		}
		for _, intf := range intfs {
			utils.DLoaderPrintf(" %s", intf.ClassName())
		}
		c.SetInterfaces(intfs)
	}
}

func (b *bstLoader) Verify(c *marea.Class) {
	// verified in class file create progress
}

func (b *bstLoader) Prepare(c *marea.Class) {
	// prepared in NewClass() func
}

// bst loader is the top level class loader
func (b *bstLoader) doLoadClassFile(class string, cp *classpath.ClassPath) (*classfile.ClassFile, error) {
	className := strings.Replace(class, ".", "/", -1)
	className += ".class"
	classData, entry, err := cp.ReadClass(className)
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Errorf("open .class failed: %s", err))
		return nil, err
	}
	reader := classfile.NewClassReader(classData)
	cf, err := classfile.NewClassFile(reader)
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Errorf("parsing class file failed: %s", err))
		return nil, err
	}

	//cf.PrintDebugMessage()
	utils.DLoaderPrintf("Load Class File %s from %s done\n", cf.ClassName(), entry.String())
	return cf, nil
}

func (loader *bstLoader) doLoadClassFromFile(file *classfile.ClassFile) *marea.Class {
	c := marea.NewClass(file)
	//file.PrintDebugMessage()
	if utils.LoaderDebugFlag && utils.DebugFlag {
		c.PrintDebugMessage()
	}
	return c
}

func (loader *bstLoader) Initiate(c *marea.Class) {
	if c.HasInitiated() || scheduleInit[c.ClassName()] {
		return
	}
	loader.doInitClass(c)
}

var scheduleInit = map[string]bool{}

func (loader *bstLoader) doInitClass(c *marea.Class) {

	workList := make([]*marea.Class, 0, 16)
	workList = append(workList, c)
	scheduleInit[c.ClassName()] = true
	// super class is defined but not initiated
	for c = c.Superclass(); c != nil && !c.HasInitiated() && !scheduleInit[c.ClassName()]; c = c.Superclass() {
		scheduleInit[c.ClassName()] = true
		workList = append(workList, c)
	}

	// init super classes
	for len(workList) > 0 {
		todo := workList[len(workList)-1]
		// setClzObj(todo) // notice that classClass may call this, class class may be nil
		// set Clz obj before call clinit, in case of call ldc in <clinit> (java/lang/Math did)
		// set Clz obj move to class define
		//adjustFields(todo) // adjust field index for class, super first
		workList = workList[:len(workList)-1]

		// call <clinit>
		clinit := todo.GetClinit()
		if clinit != nil {
			call(clinit)
		} else {
			utils.DLoaderPrintf("NO <clinit> for %s\n", todo.ClassName())
		}
		// pos init, set class object
		todo.SetInitiated(true)
	}
}

// for load array class
func (b *bstLoader) loadArrayClass(n string) *marea.Class {
	if c := cache[n]; c != nil {
		if c.InitLoader().ID() == b.id {
			return c
		} else {
			panic(utils.LinkageError)
		}
	} else {
		c = b.doLoadArrayClass(n)
		return c
	}
}

func (b *bstLoader) doLoadArrayClass(n string) *marea.Class { // support load array recursively
	utils.DLoaderPrintf("LOAD %s\n", n)
	c := &marea.Class{}
	cache[n] = c

	c.SetClassName(n)
	setClzObj(c)
	c.SetSuperClassName(utils.CLASSNAME_Object)
	c.SetSuperClass(b.loadNormalClass(utils.CLASSNAME_Object))
	c.SetInterfaceNames([]string{
		utils.CLASSNAME_Cloneable, utils.CLASSNAME_Serializable,
	})
	c.SetInterfaces([]*marea.Class{
		b.loadNormalClass(utils.CLASSNAME_Cloneable), b.loadNormalClass(utils.CLASSNAME_Serializable),
	})

	elen := cmn.ElementName(n)
	utils.DLoaderPrintf("EleN is %s\n", elen)
	if cmn.IsPrimitiveType(elen) { // the element is primitive type
		// just create the array class
		c.SetFlags(cmn.ACC_PUBLIC)
	} else if cmn.IsArray(elen) { // the element is still array type
		elec := b.loadArrayClass(elen)
		c.SetFlags(elec.GetFlags() & (^uint16(cmn.ACC_INTERFACE)) & (^uint16(cmn.ACC_ABSTRACT))) //todo
	} else { // for non array Objects
		// should load the element type first
		var elec *marea.Class
		if elec = cache[elen]; elec == nil {
			elec = b.loadNormalClass(elen) // TODO, for b to init element ?
		}
		c.SetFlags(elec.GetFlags() & (^uint16(cmn.ACC_INTERFACE)) & (^uint16(cmn.ACC_ABSTRACT))) //todo
	}
	c.SetInitLoader(b)
	c.SetDefineLoader(b)
	return c
}

//TODO seperate init from loading, jvsm 5.5
