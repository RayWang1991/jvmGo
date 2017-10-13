package marea

import (
	cf "jvmGo/ch6/classfile"
	"fmt"
)

type BootstrapMethodEntry struct {
	ref  *MethodRef
	args []interface{} // here the constant pool has been initiated
}

func NewBootstrapMethodEntry(cp ConstantPool, entry *cf.AttrBootstrapMethodEntry) *BootstrapMethodEntry {

	e := &BootstrapMethodEntry{}
	ref := cp[entry.MethodRefIndex()].(*MethodRef) // panics if the type doesn't fit
	e.ref = ref

	argIns := entry.ArgIndexes()
	args := make([]interface{}, len(argIns))
	e.args = args
	for i, aI := range argIns {
		a := cp[aI]
		switch a := a.(type) {
		case int32, float32, int64, float64, string, ClassInfo, *MethodRef, *MethodHandle:
			args[i] = a
		default:
			panic(fmt.Errorf("unsupported type %T", a))
		}
	}
	return e
}

type BootstrapMethods []*BootstrapMethodEntry

func NewBootstrapMethods(cp ConstantPool, methods cf.AttrBootstrapMethods) BootstrapMethods {
	table := methods.Table()
	bsm := make([]*BootstrapMethodEntry, len(table))
	for i, entry := range table {
		bsm[i] = NewBootstrapMethodEntry(cp, &entry)
	}
	return bsm
}
