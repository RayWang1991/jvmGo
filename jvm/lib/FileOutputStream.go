package io

import (
	"jvmGo/jvm/rtdt"
	"io"
	"os"
	"jvmGo/jvm/marea"
	"fmt"
	"jvmgo-book/v1/code/go/src/jvmgo/ch11/native"
)

func init(){
	native.Register()
}

/**
 * Writes a sub array as a sequence of bytes.
 * @param b the data to be written
 * @param off the start offset in the data
 * @param len the number of bytes that are written
 * @param append {@code true} to first advance the position to the
 *     end of file
 * @exception IOException If an I/O error has occurred.
 */
// private native void writeBytes(byte b[], int off, int len, boolean append);
// ([BIIZ)V
func writeBytes(f *rtdt.Frame) {
	vars := f.LocalVar
	this := vars.GetRef(0)   // this file output stream
	bs := vars.GetRef(1)     // bytes
	off := vars.GetInt(2)    // offset
	length := vars.GetInt(3) // len
	//append := vars.GetInt(4) != 0 // append, todo

	var w io.WriteCloser
	var isFile = false
	var err error
	// try to get Path
	pathField := this.Class().InstField("path")
	pathRef := this.GetRef(pathField.VarIdx())
	if pathRef != nil {
		path := marea.GetGoString(pathRef)
		fmt.Printf("WRITE BYTES to %s\n", path)
		w, err = os.Open(path)
		isFile = true
		if err != nil {
			panic(fmt.Errorf("open %s %s", path, err))
		}
	} else {
		fdField := this.Class().InstField("fd") // private FileDescriptor fd
		fdRef := this.GetRef(fdField.VarIdx())
		if fdRef == nil {
			panic("NOT FOUND PATH NAME AND FD!")
		}
		fddField := fdRef.Class().InstField("fd")
		fd := fdRef.GetInt(fddField.VarIdx())
		// todo
		switch fd {
		case 1: // std out
			w = os.Stdout
		case 2: // std err
			w = os.Stderr
		default:
			panic(fmt.Errorf("unsupported fd %d", fd))
		}
	}

	// todo, do write
	bytes := bs.ArrGetGoBytes()
	bytes = bytes[off:off+length]
	w.Write(bytes)
	if isFile {
		w.Close()
	}
}
