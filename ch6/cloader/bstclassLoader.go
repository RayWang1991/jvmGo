package cloader

import (
	"fmt"
	"jvmGo/ch6/classfile"
	"jvmGo/ch6/classpath"
	"jvmGo/ch6/marea"
	"jvmGo/ch6/utils"
	"os"
	"strings"
)

// TODO

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
func (b *bstLoader) Initiate(n string) *marea.Class {
	if c := cache[n]; c != nil {
		if c.InitLoader().ID() == b.id {
			return c
		} else {
			panic(utils.LinkageError)
		}
	} else {
		c := b.Define(n)
		c.SetInitLoader(b)
		b.Verify(c)
		b.Prepare(c)
		return c
	}
}

func (b *bstLoader) Define(n string) *marea.Class {
	cf, err := doLoadClassFile(n, b.cp)
	if cf == nil {
		panic(utils.ClassNotFoundException)
	}
	if err != nil {
		panic(err)
	}
	c := doLoadClassFromFile(cf)
	if c == nil {
		panic(utils.ClassFormatError)
	}
	c.SetDefineLoader(b)
	cache[n] = c
	return c
}

func (b *bstLoader) Verify(c *marea.Class) {
	// verified in class file create progress
}

func (b *bstLoader) Prepare(c *marea.Class) {
	// prepared in NewClass() func
}

// bst loader is the top level class loader
func doLoadClassFile(class string, cp *classpath.ClassPath) (*classfile.ClassFile, error) {
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
	//cf.PrintDebugMessage() TODO cf debug
	fmt.Printf("Load Class File %s from %s\n", cf.ClassName(), entry.String())
	return cf, nil
}

func doLoadClassFromFile(file *classfile.ClassFile) *marea.Class {
	c := marea.NewClass(file)
	c.PrintDebugMessage()
	return c
}
