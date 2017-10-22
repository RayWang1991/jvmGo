package classfile

import "math"

type IntegerInfo struct {
	val int32 // integer
}

func (i *IntegerInfo) ReadInfo(reader *ClassReader) {
	i.val = int32(reader.ReadUint32())
}

func (i *IntegerInfo) Value() int32 {
	return i.val
}

type FloatInfo struct {
	val float32 // float
}

func (f *FloatInfo) ReadInfo(reader *ClassReader) {
	f.val = math.Float32frombits(reader.ReadUint32())
}

func (f *FloatInfo) Value() float32 {
	return f.val
}

type LongInfo struct {
	val int64 // long
}

func (l *LongInfo) ReadInfo(reader *ClassReader) {
	l.val = int64(reader.ReadUint64())
}

func (l *LongInfo) Value() int64 {
	return l.val
}

type DoubleInfo struct {
	val float64 // double
}

func (d *DoubleInfo) ReadInfo(reader *ClassReader) {
	d.val = math.Float64frombits(reader.ReadUint64())
}

func (d *DoubleInfo) Value() float64 {
	return d.val
}
