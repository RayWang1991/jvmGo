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

var LoaderDebugFlag = true

var InitDebugFlag = true

func DLoaderPrintf(fmt string, a ... interface{}) {
	if LoaderDebugFlag {
		Dprintf(fmt, a ...)
	}
}

func DInitPrintf(fmt string, a ... interface{}) {
	if InitDebugFlag {
		Dprintf(fmt, a ...)
	}
}

var IstrDebugFlag = true

func DIstrPrintf(fmt string, a ... interface{}) {
	if IstrDebugFlag {
		Dprintf(fmt, a ...)
	}
}

var ThreadDebugFlag = true

func DThreadPrintf(fmt string, a ... interface{}) {
	if ThreadDebugFlag {
		Dprintf(fmt, a ...)
	}
}
