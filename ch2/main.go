package main

import (
	"fmt"
	"jvmGo/ch2/classpath"
	"strings"
)

func main() {
	cmd := ParseCmd()
	if cmd.helpFlag {
		printUsage()
	} else if cmd.versionFlag {
		// print version
		fmt.Println("version 0.0.1")
	} else {
		// startJVM
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.NewClassPath(cmd.xjreOption, cmd.cpOption)
	fmt.Printf("class path: %s class: %s args: %s\n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	className += ".class"
	classData, _, err := cp.ReadClass(className)
	if err == nil {
		fmt.Printf("data: %v", classData)
	} else {
		fmt.Printf("error: %s\n", err)
	}
}
