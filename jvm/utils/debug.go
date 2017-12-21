package utils

import (
	"fmt"
)

var DebugFlag = true

func Dprintf(format string, a ... interface{}) {
	if DebugFlag {
		fmt.Printf(format, a ...)
	}
}

var LoaderDebugFlag = false

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

var MainDebugFlag = false

func DMainPrintf(fmt string, a ... interface{}) {
	if MainDebugFlag {
		Dprintf(fmt, a ...)
	}
}
