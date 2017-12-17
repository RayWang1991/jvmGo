package marea

type ClassLoader interface {
	ID() int
	Delegate() ClassLoader
	Load(n string) *Class // wrapper for all type
	Define(n string) *Class
	Initiate(class *Class) // init class
	Verify(class *Class)
	Prepare(class *Class)
	//LoadArrayClass(n string) *Class // for array type
}

const (
	BootstrapClassLoaderId   = iota
	UserDefinedClassLoaderId
)

var DefaultLoader ClassLoader
