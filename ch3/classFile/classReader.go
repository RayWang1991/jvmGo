package classFile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// read a byte
func (cr *ClassReader) ReadUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

// read 2 bytes
func (cr *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

// read 4 bytes
func (cr *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

// read 8 bytes
func (cr *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

func (cr *ClassReader) ReadBytes(n uint) []byte { // can do this? [n]byte
	val := []byte(cr.data[:n])
	cr.data = cr.data[n:]
	return val
}

func (cr *ClassReader) ReadUint16s() []uint16 {
	n := cr.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = cr.ReadUint16()
	}
	return s
}
