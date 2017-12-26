package lib

import (
	"jvmGo/jvm/utils"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/marea"
	"fmt"
	"os"
	"io"
)

func init() {
	register(utils.CLASSNAME_FileInputStream, "open0", "(Ljava/lang/String;)V", open0)
	register(utils.CLASSNAME_FileInputStream, "readBytes", "([BII)I", readBytes)
}

/**
 * Reads a subarray as a sequence of bytes.
 * @param b the data to be written
 * @param off the start offset in the data
 * @param len the number of bytes that are written
 * @exception IOException If an I/O error has occurred.
 */
//private native int readBytes(byte b[], int off, int len) throws IOException;
//([BII)I
func readBytes(f *rtdt.Frame) {
	vars := f.LocalVar
	this := vars.GetRef(0)   // this file output stream
	bs := vars.GetRef(1)     // bytes
	off := vars.GetInt(2)    // offset
	length := vars.GetInt(3) // len
	//append := vars.GetInt(4) != 0 // append, todo

	//debug
	fmt.Printf("this:\n %s\n%s", this, this.ListAllFields())

	var r io.ReadCloser
	var isFile = false
	var err error
	// try to get Path
	pathField := this.Class().InstField("path")
	pathRef := this.GetRef(pathField.VarIdx())
	if pathRef != nil {
		path := marea.GetGoString(pathRef)
		fmt.Printf("READ BYTES to %s\n", path)
		r, err = os.Open(path)
		if err != nil {
			//panic(fmt.Errorf("open %s %s", path, err))
		} else {
			isFile = true
		}
	}
	panic("todo")
	if !isFile {
		fdField := this.Class().InstField("fd") // private FileDescriptor fd
		fdRef := this.GetRef(fdField.VarIdx())
		if fdRef == nil {
			panic("NOT FOUND PATH NAME AND FD!")
		}
		fddField := fdRef.Class().InstField("fd")
		fd := fdRef.GetInt(fddField.VarIdx())
		// todo
		switch fd {
		case 0: // std in
			r = os.Stdin
		default:
			r = os.Stdin
			//panic(fmt.Errorf("unsupported fd %d", fd))
		}
	}

	// todo, do write
	bytes := bs.ArrGetGoBytes()
	bytes = bytes[off:off+length]
	n, err := r.Read(bytes)
	if isFile {
		r.Close()
	}
	//todo
	if err != nil {
		panic(fmt.Errorf("read %s %s", r, err))
	}
	f.OperandStack.PushInt(int32(n))
}

// private native void open0(String name) throws FileNotFoundException;
// (Ljava/lang/String;)V
func open0(f *rtdt.Frame) {
	//todo
	this := f.LocalVar.GetRef(0)
	jname := f.LocalVar.GetRef(1)
	name := marea.GetGoString(jname)
	fmt.Printf("this %s name %s\n", this, name)
}
