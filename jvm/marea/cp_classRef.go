package marea

import "jvmGo/jvm/cmn"

type ClassRef struct {
	className string // Key for from
	from      *Class
	ref       *Class // cache
}

func NewClassRef(cn string, cls *Class) *ClassRef {
	cr := &ClassRef{
		className: cn,
		from:      cls,
	}
	return cr
}

func (c *ClassRef) ClassName() string {
	return c.className
}

func (c *ClassRef) FromClass() *Class {
	return c.from
}

// not concurrently safe
func (c *ClassRef) Ref() *Class {
	if c.ref == nil {
		c.resolveRef()
	}
	return c.ref
}

func (c *ClassRef) SetFromClass(cls *Class) {
	c.from = cls
}

func (c *ClassRef) resolveRef() {
	l := c.from.defLoader
	name := cmn.SimClassName(c.className)
	cls := l.Load(name)
	c.ref = cls
}
