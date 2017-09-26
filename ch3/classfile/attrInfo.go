package attribute

import "jvmGo/ch3/classfile"

type AttrInfo interface {
	ReadInfo(reader classfile.ClassReader)
}

type AttrInfoBase struct{
  nameIndex uint16
  length uint64
}


