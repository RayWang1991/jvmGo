package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	versionFlag bool
	helpFlag    bool
	cpOption    string
	class       string
	args        []string
}

func ParseCmd() *Cmd {
	c := new(Cmd)
	flag.BoolVar(&c.helpFlag, "help", false, "print help message")
	flag.BoolVar(&c.helpFlag, "?", false, "print help message")
	flag.BoolVar(&c.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&c.cpOption, "classpath", "", "classpath")
	flag.StringVar(&c.cpOption, "cp", "", "classpath")
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
