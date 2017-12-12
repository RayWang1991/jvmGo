package cmn

import (
	"strings"
)

func IsArray(name string) bool {
	return strings.HasPrefix(name, "[")
}

func ElementName(name string) string { // ignore '[
	return name[1:]
}

func IsPrimitiveType(name string) bool {
	if len(name) != 1 {
		return false
	}
	c := name[0]
	switch c {
	case 'Z', 'B', 'C', 'S', 'F', 'D', 'I', 'J':
		return true
	default:
		return false
	}

	if name[0] == '[' || name[0] == 'L' {
		return false
	} else {
		return true
	}
}

// for array class name to element the L and ;
func SimClassName(name string) string {
	if len(name) == 0 || name[0] != '[' {
		return name
	}
	// element first L
	i := strings.IndexByte(name, 'L')
	if i < 0 {
		// not find L
		return name
	}
	return name[:i] + name[i+1:len(name)-1]
}

func ToDoted(str string) string {
	return strings.Replace(str, "/", ".", -1)
}

func ToSlash(str string) string {
	return strings.Replace(str, ".", "/", -1)
}
