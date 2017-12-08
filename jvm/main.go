package main

import (
	"fmt"
	"jvmGo/jvm/classpath"
	"jvmGo/jvm/cloader"
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/marea"
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
	marea.DefaultLoader = bstLoader
	bstLoader.SetUpBase()
	// load class
	c := bstLoader.Initiate(cmd.class)
	m := c.GetMain()
	if m == nil {
		println("not found method 	'main'")
	} else {
		mainThread = rtdt.NewThread(1024)
		mainThread.PushFrame(rtdt.NewFrame(m,mainThread))
		interpretMain(m, cmd.args)
	}
}

