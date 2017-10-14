package classpath

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// zip entry implements Entry interface
// representing a jar entry

type ZipEntry struct {
	absPath string // absolute path for jar
}

// ReadClass open the jar and search for that class
func (z *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	rc, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer rc.Close()
	for _, f := range rc.File {
		if f.Name == className {
			frc, err := f.Open()
			if err != nil {
				frc.Close()
				return nil, nil, err
			}
			data, err := ioutil.ReadAll(frc)
			frc.Close()
			if err != nil {
				return nil, nil, err
			}
			return data, z, nil
		}
	}
	return nil, nil, fmt.Errorf("not found class name %s", className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}

func NewZipEntry(path string) *ZipEntry {
	if !filepath.IsAbs(path) {
		panic("not abs path")
	}
	return &ZipEntry{path}
}
