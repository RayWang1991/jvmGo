package classpath

import (
	"bytes"
	"fmt"
	"strings"
)

// CompositeEntry represent a list of entry
type CompositeEntry []Entry

func (c CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		if b, e, err := entry.ReadClass(className); err == nil {
			return b, e, nil
		}
	}
	return nil, nil, fmt.Errorf("not found %s in %s", className, c.String())
}

func (c CompositeEntry) String() string {
	buf := bytes.Buffer{}
	for i, e := range c {
		if i > 0 {
			buf.WriteRune(seperator)
		}
		buf.WriteString(e.String())
	}
	return buf.String()
}

func NewComposite(pathList string) CompositeEntry {
	paths := strings.Split(pathList, string(seperator))
	entrs := make([]Entry, 0, len(paths))
	for _, p := range paths {
		e := NewEntry(p)
		entrs = append(entrs, e)
	}
	return CompositeEntry(entrs)
}
