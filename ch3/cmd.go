package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	debugFlag bool
	versionFlag bool
	helpFlag    bool
	cpOption    string
	xjreOption  string
	class       string
	args        []string
}

func ParseCmd() *Cmd {
	c := new(Cmd)
	flag.BoolVar(&c.debugFlag, "debug", false, "debug flag")
	flag.BoolVar(&c.helpFlag, "help", false, "print help message")
	flag.BoolVar(&c.helpFlag, "?", false, "print help message")
	flag.BoolVar(&c.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&c.cpOption, "classpath", "", "classpath")
	flag.StringVar(&c.cpOption, "cp", "", "classpath")
	flag.StringVar(&c.cpOption, "Xjre", "", "path to jre(nonstandard)")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		c.class = args[0]
		c.args = args[1:]
	}
	return c
}

func printUsage() {
	fmt.Printf("Usage: %s [-optons] class [args...]\n", os.Args[0])
}
