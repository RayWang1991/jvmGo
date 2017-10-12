package marea

type ClassRef struct {
	//from      *Class // point to Class where it is from
	className string // Key for class
	class     *Class
}

func NewClassRef(cn string) *ClassRef {
	cr := &ClassRef{}
	cr.className = cn
	return cr
}

func (c *ClassRef) ClassName() string {
	return c.className
}

func (c *ClassRef) Class() *Class {
	return c.Class()
}

func (c *ClassRef) SetClass(cls *Class) {
	c.class = cls
}
