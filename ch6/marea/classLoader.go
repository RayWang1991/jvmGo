package marea

type ClassLoader interface {
	ID() int
	Delegate() ClassLoader
	Initiate(n string) *Class
	Define(n string) *Class
	Verify(class *Class)
	Prepare(class *Class)
}

const (
	BootstrapClassLoaderId   = iota
	UserDefinedClassLoaderId
)
