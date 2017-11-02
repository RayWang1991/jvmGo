package classfile

import "jvmGo/jvm/cmn"

type Utf8Info struct {
	val string // []byte data
}

func (u *Utf8Info) ReadInfo(reader *ClassReader) {
	length := reader.ReadUint16() // length for the utf8 info in bytes
	bs := reader.ReadBytes(uint(length))
	u.val = cmn.MUTF8ToUTF8(bs)
}

func (u *Utf8Info) String() string {
	return u.val
}
