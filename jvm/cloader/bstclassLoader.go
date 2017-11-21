package cloader

import (
	"fmt"
	"jvmGo/jvm/classfile"
	"jvmGo/jvm/classpath"
	"jvmGo/jvm/cmn"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
	"os"
	"strings"
)

// TODO
// TODO the loader thread is just memory keeper now, add concurrent logic

func NewBstLoader(cp *classpath.ClassPath) marea.ClassLoader {
	return &bstLoader{
		id: marea.BootstrapClassLoaderId,
		cp: cp,
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
	fmt.Printf("Load %s\n", n)
	if cmn.IsArray(n) {
		return b.LoadArrayClass(n)
	} else {
		return b.Initiate(n)
	}
}

func (b *bstLoader) Initiate(n string) *marea.Class {
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

func (b *bstLoader) Define(n string) *marea.Class {
	cf, err := b.doLoadClassFile(n, b.cp)
	fmt.Printf("define Class %s\n", n)

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
	cache[n] = c
	c.SetInitLoader(b)
	c.SetDefineLoader(b)
	b.doInitClass(c)
	return c
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

	cf.PrintDebugMessage()
	utils.DLoaderPrintf("Load Class File %s from %s\n", cf.ClassName(), entry.String())
	return cf, nil
}

func (loader *bstLoader) doLoadClassFromFile(file *classfile.ClassFile) *marea.Class {
	c := marea.NewClass(file)
	c.PrintDebugMessage()
	return c
}

var scheduleInit = map[string]bool{}

// call <clint> for class
func (loader *bstLoader) doInitClass(c *marea.Class) {
	if c.HasInitiated() || scheduleInit[c.ClassName()] {
		return
	}
	// set up loader thread
		utils.DInitPrintf("\n##INIT CLASS## %s\n", c.ClassName())
	t := GetLoaderThread()
	oldF := t.GetFrameSize()
	for c != nil {
		if !c.HasInitiated() && !scheduleInit[c.ClassName()] {
			scheduleInit[c.ClassName()] = true
			m := c.GetClinit()
			if m != nil {
				f := rtdt.NewFrame(m, t)
				// TODO, hack
				/*
					if c.ClassName() == "utils.CLASSNAME_System" {
						// call initializeSystemClass
						m = c.GetMethodDirect(cmn.METHOD_initializeSystemClass_NAME, cmn.METHOD_initializeSystemClass_DESC)
						t.PushFrame(rtdt.NewFrame(m, t))
					}
				*/
				t.PushFrame(f)
				defer func(c *marea.Class) { // TODO, actually init is done earlier
					c.SetInitiated(true)
					delete(scheduleInit, c.ClassName())
				}(c)
			}
		} else if c.HasInitiated() {
			break
		}
		s := c.SuperclassName()
		if s != "" {
			c = loader.Initiate(c.SuperclassName())
		} else {
			c = nil
		}
	}

	// loop if needed
	if t.GetFrameSize() > oldF {
		loop(t)
	}
}

// for load array class
func (b *bstLoader) LoadArrayClass(n string) *marea.Class {
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
	c := &marea.Class{}
	cache[n] = c

	c.SetClassName(n)
	c.SetSuperClassName(utils.CLASSNAME_Object)
	c.SetSuperClass(b.Initiate(utils.CLASSNAME_Object))
	c.SetInterfaceNames([]string{
		utils.CLASSNAME_Cloneable, utils.CLASSNAME_Serializable,
	})
	c.SetInterfaces([]*marea.Class{
		b.Initiate(utils.CLASSNAME_Cloneable), b.Initiate(utils.CLASSNAME_Serializable),
	})

	elen := cmn.ElementName(n)
	fmt.Printf("EleN is %s\n", elen)
	if cmn.IsPrimitiveType(elen) { // the element is primitive type
		// just create the array class
		c.SetFlags(cmn.ACC_PUBLIC)
	} else if cmn.IsArray(elen) { // the element is still array type
		elec := b.LoadArrayClass(elen)
		c.SetFlags(elec.GetFlags())
	} else { // for non array Objects
		// should load the element type first
		var elec *marea.Class
		if elec = cache[elen]; elec == nil {
			elec = b.Initiate(elen) // TODO, for b to init element ?
		}
		c.SetFlags(elec.GetFlags())
	}
	c.SetInitLoader(b)
	c.SetDefineLoader(b)
	return c
}
