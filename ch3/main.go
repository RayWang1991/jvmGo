package main

import (
	"fmt"
	"jvmGo/ch3/classfile"
	"jvmGo/ch3/classpath"
	"log"
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
	if err != nil {
		log.Fatalf("open .class failed: %s", err)
	}
	if cmd.debugFlag {
		fmt.Printf("data: %v", classData)
	}
	reader := classfile.NewClassReader(classData)
	cf, err := classfile.NewClassFile(reader)
	if err != nil {
		log.Fatalf("parsing class file failed: %s", err)
	}
	cf.PrintDebugMessage()
}
