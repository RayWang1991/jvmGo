package lib

import (
	"jvmGo/jvm/utils"
	"jvmGo/jvm/rtdt"
	"runtime"
	"os"
	"jvmGo/jvm/marea"
)

func init() {
	register(utils.CLASSNAME_System, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
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
