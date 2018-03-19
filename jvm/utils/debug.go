package utils

import (
	"fmt"
)

var DebugFlag = false

func Dprintf(format string, a ... interface{}) {
	if DebugFlag {
		fmt.Printf(format, a ...)
	}
}

var LoaderDebugFlag = true

func DLoaderPrintf(fmt string, a ... interface{}) {
	if LoaderDebugFlag {
		Dprintf(fmt, a ...)
	}
}

var IstrDebugFlag = false

func DIstrPrintf(fmt string, a ... interface{}) {
	if IstrDebugFlag {
		Dprintf(fmt, a ...)
	}
}

var NativeDebugFlag = true

func DNativePrintf(fmt string, a ... interface{}) {
	if NativeDebugFlag {
		Dprintf(fmt, a ...)
	}
}

var ThreadDebugFlag = false

func DThreadPrintf(fmt string, a ... interface{}) {
	if ThreadDebugFlag {
		Dprintf(fmt, a ...)
	}
}

var MainDebugFlag = true

func DMainPrintf(fmt string, a ... interface{}) {
	if MainDebugFlag {
		Dprintf(fmt, a ...)
	}
}

var CallTraceFlag = false

func DCallTracePrintf(format string, a ... interface{}) {
	if CallTraceFlag {
		fmt.Printf(format, a...)
	}
}
