package classpath

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

// zip entry implements Entry interface
// representing a jar entry

// TODO 1 test concurrently
// TODO 2 add cache for files

var zips = make(map[string]*ZipEntry)

var mux = sync.Mutex{}

const ExpireTime = time.Minute * 5

type ZipEntry struct {
	absPath string          // absolute path for jar
	using   chan struct{}   // closed or nil when entry is not using
	ready   chan struct{}   // closed when entry is ready
	rc      *zip.ReadCloser // used to read and close the entry
}

// Concurrently safe
// ReadClass open the jar and search for that class
func (_z *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	z := getZipEntry(_z.absPath)

	// in using
	z.using = make(chan struct{})
	// release using state
	defer close(z.using)

	for _, f := range z.rc.File {
		if f.Name == className {
			data, err := readFile(f)
			if err != nil {
				return nil, nil, err
			}
			return data, z, nil
		}
	}
	return nil, nil, fmt.Errorf("not found class name %s", className)
}

func getZipEntry(absPath string) *ZipEntry {
	mux.Lock()
	e := zips[absPath]
	if e == nil {
		// This is the first read for abs path
		// do the work and cache the result
		// then broadcast it is ready
		e = &ZipEntry{
			absPath: absPath,
			using:   nil,
			ready:   make(chan struct{}),
			rc:      nil,
		}
		zips[absPath] = e
		mux.Unlock()
		// close the entry and delete it from the map when it's expired
		go delayClose(e)
		rc, err := zip.OpenReader(absPath)
		if err != nil {
			e.rc = nil
		} else {
			e.rc = rc
		}
		close(e.ready)
	} else {
		// e is already asked
		mux.Unlock()
		<-e.ready
	}
	return e
}

func delayClose(z *ZipEntry) {
	t := time.NewTimer(ExpireTime)
	<-t.C
	t.Stop()
	mux.Lock()
	delete(zips, z.absPath) // broadcast z is expired
	mux.Unlock()
	if z.using != nil {
		<-z.using
	}
	z.rc.Close()
}

func readFile(f *zip.File) ([]byte, error) {
	frc, err := f.Open()
	if err != nil {
		frc.Close()
		return nil, err
	}
	data, err := ioutil.ReadAll(frc)
	frc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (z *ZipEntry) String() string {
	return z.absPath
}

func NewZipEntry(path string) *ZipEntry {
	if !filepath.IsAbs(path) {
		panic("not abs path")
	}
	if z := zips[path]; z != nil {
		return z
	}
	return &ZipEntry{path, nil, nil, nil}
}
