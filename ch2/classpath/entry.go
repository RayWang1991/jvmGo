package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type Entry interface {
	// ReadClass read the class name and return the data for that class, the final entry where the class locate, and
	// error if there is
	ReadClass(className string) ([]byte, Entry, error)
	// string returns the representation of the entry
	String() string
}

const seperator = os.PathListSeparator

// wrapper for initiation dir, wildcard, composite, zip
func NewEntry(path string) Entry {
	if strings.HasSuffix(path, "*") {
		return NewWildcardEntry(path)
	}
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") {
		return NewZipEntry(path)
	}
	if strings.Contains(path, strings(seperator)) {
		return NewComposite(path)
	}
	return NewDirEntry(path)
}
