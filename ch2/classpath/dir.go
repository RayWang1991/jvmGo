package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry is an entry representing a basic absolute directory
type DirEntry struct {
	//absolute directory path for that entry
	absDir string
}

func (entry *DirEntry) String() string {
	return entry.absDir
}

func (entry *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	p := filepath.Join(entry.absDir, className)
	d, err := ioutil.ReadFile(p)
	return d, entry, err
}

// NewDirEntry returns an directory entry containing an absolute path
// if the path can not be converted to an absolute path, it panics
func NewDirEntry(path string) *DirEntry {
	// check the path
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{path}
}
