// package class path provides methods to find the Jre path, extension path and user path,
// in order to find the .class file in the disk
package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

// parse ClassPath entity from jre op and
func NewClassPath(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.getBootAndExtClassPath(jreOption)
	cp.getUserClassPath(cpOption)
	return cp
}

// get boot path and extend path from jre option
func (cp *ClassPath) getBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	jreExtLibPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.bootClassPath = NewWildcardEntry(jreLibPath)
	cp.extClassPath = NewWildcardEntry(jreExtLibPath)
}

const jreDefault = "/Library/Internet Plug-Ins/JavaAppletPlugin.plugin/Contents/Home"

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		fmt.Println("JAVA_HOME:" + jh)
		return filepath.Join(jh, "jre")
	}
	return jreDefault
	panic("not found jre folder")
}

func (cp *ClassPath) getUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	if exists(cpOption) {
		cp.userClassPath = NewEntry(cpOption)
	}
}

// find out whether the dir exists
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (cp *ClassPath) ReadClass(className string) (d []byte, e Entry, err error) {
	d, e, err = cp.bootClassPath.ReadClass(className)
	if err == nil {
		return
	}
	d, e, err = cp.extClassPath.ReadClass(className)
	if err == nil {
		return
	}
	return cp.userClassPath.ReadClass(className)
}
