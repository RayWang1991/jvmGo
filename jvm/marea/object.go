package marea

// An Object represents a general reference type, for from type, interface type, and array type
type Object struct {
	class *Class
	Vars
}

func NewObject(class *Class) *Object {
	vs := NewVars(class.InsSlotNum())
	return &Object{
		class: class,
		Vars:  vs,
	}
}

func (o *Object) Class() *Class {
	return o.class
}
