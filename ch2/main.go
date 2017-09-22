package main

import "fmt"

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
}
