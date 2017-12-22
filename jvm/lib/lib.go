// lib package provides native methods for core class
package lib

import (
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
)

var naMap map[string]NativeFunc = make(map[string]NativeFunc) // native function mapper

type NativeFunc func(frame *rtdt.Frame)

func init() {
	// init naMap
	//objs := []NativeFunc{registerNatives}
	//for _, f := range objs {
	//	t := reflect.ValueOf(f)
	//	name := runtime.FuncForPC(t.Pointer()).Name()
	//	fmt.Println(name)
	//}
	// register method
}

func register(cname, mname, desc string, nf NativeFunc) {
	naMap[key(cname, mname, desc)] = nf
}

func key(cname, mname, desc string) string {
	return cname + " " + mname + " " + desc
}

func CallNative(m *marea.Method, t *rtdt.Thread) {
	utils.DNativePrintf("[Call Native] %s %s ", m.Name(), m.Class().ClassName())
	if method := naMap[key(m.Class().ClassName(), m.Name(), m.Desc())]; method != nil {
		utils.DNativePrintf("RES: Found\n")
		method(t.CurrentFrame())
	} else {
		utils.DNativePrintf("[NATIVE] %s %s RES: Unsupported\n", m.Name(), m.Desc())
		t.PopFrame()
	}
}
