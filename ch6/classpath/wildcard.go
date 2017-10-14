package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// wildcard entry implements Entry interface
// representing a '*' entry, it's a composite entry after all

func NewWildcardEntry(entryPath string) CompositeEntry {
	if !filepath.IsAbs(entryPath) {
		panic("not absolute path")
	}
	baseName := entryPath[:len(entryPath)-1] // should check for the len
	entries := []Entry{}
	filepath.Walk(baseName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != baseName {
			return filepath.SkipDir // do not support search recursively
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ",JAR") {
			entries = append(entries, NewZipEntry(path))
		}
		return nil
	})
	return entries
}
