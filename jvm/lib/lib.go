// lib package provides native methods for core class
package lib

import (
	"fmt"
	"jvmGo/jvm/marea"
	"jvmGo/jvm/rtdt"
	"reflect"
	"runtime"
)

var naMap map[string]NativeFunc // native function mapper

type NativeFunc func(frame *rtdt.Frame)

func init() {
	// init naMap
	objs := []NativeFunc{registerNatives}
	for _, f := range objs {
		t := reflect.ValueOf(f)
		name := runtime.FuncForPC(t.Pointer()).Name()
		fmt.Println(name)
	}
}

func CallNative(m *marea.Method) {
	// TODO
}
