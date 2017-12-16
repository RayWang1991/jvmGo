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
