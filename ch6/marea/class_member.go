package marea

type ClassMember struct {
	class *Class // pointer to class
	flags uint16
	name  string
	desc  string
}

// getters
func (c *ClassMember) Class() *Class {
	return c.class
}

func (c *ClassMember) Flags() uint16 {
	return c.flags
}

func (c *ClassMember) Name() string {
	return c.name
}

func (c *ClassMember) Desc() string {
	return c.desc
}
