package lib

import (
	"jvmGo/jvm/utils"
	"jvmGo/jvm/rtdt"
	"runtime"
	"os"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/cmn"
)

func init() {
	register(utils.CLASSNAME_System, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	register(utils.CLASSNAME_System, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	register(utils.CLASSNAME_System, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	register(utils.CLASSNAME_System, "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	register(utils.CLASSNAME_System, "setErr0", "(Ljava/io/PrintStream;)V", setErr0)
}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtdt.Frame) {
	vars := frame.LocalVar
	props := vars.GetRef(0)

	stack := frame.OperandStack
	stack.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class().GetMethodDirect("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread()
	loader := frame.Method().Class().DefineLoader()
	for key, val := range _sysProperties() {
		jKey := marea.GetJavaString(key, loader)
		jVal := marea.GetJavaString(val, loader)
		ops := rtdt.NewOperandStack(3)
		ops.PushRef(props)
		ops.PushRef(jKey)
		ops.PushRef(jVal)
		df := dummyFrame(ops, thread)
		thread.PushFrame(df)
		callMethod(setPropMethod, df)
	}
}

func _sysProperties() map[string]string {
	return map[string]string{
		"java.version":        "0.0.1",
		"java.vendor":         "govm",
		"java.class.version":  "52.0",
		"os.name":             runtime.GOOS,
		"os.arch":             runtime.GOARCH,
		"file.separator":      string(os.PathListSeparator),
		"path.separator":      string(os.PathSeparator),
		"line.separator":      "\n",
		"user.dir":            ".",
		"file.encoding":       "UTF-8",
		"sun.stdout.encoding": "UTF-8",
		"sun.stderr.encoding": "UTF-8",
	}
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtdt.Frame) {
	vars := frame.LocalVar
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dst := vars.GetRef(2)
	dstPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if dst == nil || src == nil {
		panic(utils.NullPointerException)
	}
	srcCls := src.Class()
	destCls := dst.Class()
	if !srcCls.IsArray() || !destCls.IsArray() {
		panic(utils.ArrayStoreException)
	}
	srcEleName := cmn.ElementName(srcCls.ClassName())
	destEleName := cmn.ElementName(destCls.ClassName())
	isSrcEleP := cmn.IsPrimitiveType(srcEleName)
	isDestEleP := cmn.IsPrimitiveType(destEleName)
	if isSrcEleP != isDestEleP || (isSrcEleP && isDestEleP && srcEleName != destEleName) {
		panic(utils.ArrayStoreException)
	}
	srcN := src.ArrayLength()
	destN := dst.ArrayLength()
	if srcPos < 0 || dstPos < 0 || length < 0 || srcPos+length > srcN || dstPos+length > destN {
		panic(utils.ArrayIndexOutOfBoundsException)
	}
	switch src.Data().(type) {
	case []int8:
		_src := src.Data().([]int8)[srcPos: srcPos+length]
		_dst := dst.Data().([]int8)[dstPos: dstPos+length]
		copy(_dst, _src)
	case []int16:
		_src := src.Data().([]int16)[srcPos: srcPos+length]
		_dst := dst.Data().([]int16)[dstPos: dstPos+length]
		copy(_dst, _src)
	case []int32:
		_src := src.Data().([]int32)[srcPos: srcPos+length]
		_dst := dst.Data().([]int32)[dstPos: dstPos+length]
		copy(_dst, _src)
	case []int64:
		_src := src.Data().([]int64)[srcPos: srcPos+length]
		_dst := dst.Data().([]int64)[dstPos: dstPos+length]
		copy(_dst, _src)
	case []uint16:
		_src := src.Data().([]uint16)[srcPos: srcPos+length]
		_dst := dst.Data().([]uint16)[dstPos: dstPos+length]
		copy(_dst, _src)
	case []float32:
		_src := src.Data().([]float32)[srcPos: srcPos+length]
		_dst := dst.Data().([]float32)[dstPos: dstPos+length]
		copy(_dst, _src)
	case []float64:
		_src := src.Data().([]float64)[srcPos: srcPos+length]
		_dst := dst.Data().([]float64)[dstPos: dstPos+length]
		copy(_dst, _src)
	case []*marea.Object:
		_src := src.Data().([]*marea.Object)[srcPos: srcPos+length]
		_dst := dst.Data().([]*marea.Object)[dstPos: dstPos+length]
		copy(_dst, _src)
	default:
		panic("can not be!")
	}
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(f *rtdt.Frame) {
	in := f.LocalVar.GetRef(0)
	loader := f.Method().Class().DefineLoader()
	sysC := loader.Load(utils.CLASSNAME_System)
	inField := sysC.StatField("in")
	sysC.SetRef(in, inField.VarIdx())
}

// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(f *rtdt.Frame) {
	out := f.LocalVar.GetRef(0)
	loader := f.Method().Class().DefineLoader()
	sysC := loader.Load(utils.CLASSNAME_System)
	outField := sysC.StatField("out")
	sysC.SetRef(out, outField.VarIdx())
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(f *rtdt.Frame) {
	err := f.LocalVar.GetRef(0)
	loader := f.Method().Class().DefineLoader()
	sysC := loader.Load(utils.CLASSNAME_System)
	errField := sysC.StatField("err")
	sysC.SetRef(err, errField.VarIdx())
}
