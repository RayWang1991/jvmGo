package utils

import "log"

var DebugFlag = true

func Dprintf(fmt string, a ... interface{}) {
	if DebugFlag {
		log.Printf(fmt, a ...)
	}
}

var LoaderDebugFlag = false

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
