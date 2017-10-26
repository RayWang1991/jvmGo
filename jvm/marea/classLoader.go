package marea

type ClassLoader interface {
	ID() int
	Delegate() ClassLoader
	Load(n string) *Class     // wrapper for all type
	Initiate(n string) *Class // for non array type
	Define(n string) *Class
	Verify(class *Class)
	Prepare(class *Class)
	LoadArrayClass(n string) *Class // for array type
}

const (
	BootstrapClassLoaderId = iota
	UserDefinedClassLoaderId
)
