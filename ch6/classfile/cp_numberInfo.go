package classfile

import "math"

type IntegerInfo struct {
	val int32 // integer
}

func (i *IntegerInfo) ReadInfo(reader *ClassReader) {
	i.val = int32(reader.ReadUint32())
}

type FloatInfo struct {
	val float32 // float
}

func (f *FloatInfo) ReadInfo(reader *ClassReader) {
	f.val = math.Float32frombits(reader.ReadUint32())
}

type LongInfo struct {
	val int64 // long
}

func (l *LongInfo) ReadInfo(reader *ClassReader) {
	l.val = int64(reader.ReadUint64())
}

type DoubleInfo struct {
	val float64 // double
}

func (d *DoubleInfo) ReadInfo(reader *ClassReader) {
	d.val = math.Float64frombits(reader.ReadUint64())
}
