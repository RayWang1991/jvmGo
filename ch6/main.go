package main

import (
	"fmt"
	"jvmGo/ch6/classpath"
	"jvmGo/ch6/cloader"
	"jvmGo/ch6/marea"
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
	m := getMain(c)
	if m == nil {
		println("not found method 	'main'")
	} else {
		interpret(m)
	}
}

func getMain(c *marea.Class) *marea.Method {
	return c.GetMethodDirect("main", "([Ljava/lang/String;)V")
}
