package cloader

import "jvmGo/ch6/marea"

// TODO

type ClassLoader interface {
	Delegate() ClassLoader
	Initiate(class *marea.Class)
	Define(class *marea.Class)
	Verify(class *marea.Class)
	Prepare(class *marea.Class)
}

const (
	BootstrapClassLoaderId   = iota
	UserDefinedClassLoaderId
)

var BstLoader ClassLoader = &bstLoader{}

type bstLoader struct {
}

func (b *bstLoader) Delegate() ClassLoader {
	return nil
}
func (b *bstLoader) Initiate(c *marea.Class) {
}
func (b *bstLoader) Define(c *marea.Class) {
}
func (b *bstLoader) Verify(c *marea.Class) {
}
func (b *bstLoader) Prepare(c *marea.Class) {
}
