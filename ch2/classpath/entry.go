package classpath

import "os"

type Entry interface {
	// ReadClass read the class name and return the data for that class, the final entry where the class locate, and
	// error if there is
	ReadClass(className string) ([]byte, Entry, error)
	// string returns the representation of the entry
	String() string
}

const seperator = os.PathListSeparator

func NewEntry(path string)Entry{
	// TODO
}