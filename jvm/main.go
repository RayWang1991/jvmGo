package main

import (
	"fmt"
	"jvmGo/jvm/classpath"
	"jvmGo/jvm/cloader"
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
	bstLoader := cloader.NewBstLoader(cp)
	c := bstLoader.Initiate(cmd.class)
	m := c.GetMain()
	if m == nil {
		println("not found method 	'main'")
	} else {
		interpret(m)
	}
}
