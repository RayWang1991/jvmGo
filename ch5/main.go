package main

import (
	"fmt"
	"jvmGo/ch5/classpath"
	"jvmGo/ch5/classfile"
	"strings"
	"log"
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
	fmt.Printf("class path: %s class: %s args: %s\n", cmd.cpOption, cmd.class, cmd.args)
	cp := classpath.NewClassPath(cmd.xjreOption, cmd.cpOption)
	cf := loadClass(cp, cmd.class, cmd.debugFlag)
	main := getMain(cf)
	if main == nil {
		println("not found method 	'main'")
	}
	interpret(main)
}

func loadClass(cp *classpath.ClassPath, class string, debug bool) *classfile.ClassFile {
	className := strings.Replace(class, ".", "/", -1)
	className += ".class"
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		log.Fatalf("open .class failed: %s", err)
	}
	if debug {
		fmt.Printf("data: %v", classData)
	}
	reader := classfile.NewClassReader(classData)
	cf, err := classfile.NewClassFile(reader)
	if err != nil {
		log.Fatalf("parsing class file failed: %s", err)
	}
	if debug {
		cf.PrintDebugMessage()
	}
	return cf
}

func getMain(cf *classfile.ClassFile) *classfile.MethodInfo {
	for _, m := range cf.MethodInfo() {
		if m.Name() == "main" && m.Description() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
