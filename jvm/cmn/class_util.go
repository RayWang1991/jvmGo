package cmn

import (
	"strings"
	"jvmGo/jvm/utils"
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

// for array class name to delete the L and ;
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

var PrimitiveNames = map[string]string{
	"Z": utils.CLASSNAME_prim_boolean,
	"B": utils.CLASSNAME_prim_byte,
	"C": utils.CLASSNAME_prim_char,
	"S": utils.CLASSNAME_prim_short,
	"F": utils.CLASSNAME_prim_float,
	"D": utils.CLASSNAME_prim_double,
	"I": utils.CLASSNAME_prim_int,
	"J": utils.CLASSNAME_prim_long,
}

func ToClassName(desc string) string {
	if len(desc) == 1 {
		switch desc[0] {
		case 'Z':
			return utils.CLASSNAME_prim_boolean
		case 'B':
			return utils.CLASSNAME_prim_byte
		case 'C':
			return utils.CLASSNAME_prim_char
		case 'S':
			return utils.CLASSNAME_prim_short
		case 'F':
			return utils.CLASSNAME_prim_float
		case 'D':
			return utils.CLASSNAME_prim_double
		case 'I':
			return utils.CLASSNAME_prim_int
		case 'J':
			return utils.CLASSNAME_prim_long
		default:
			panic("unsupported desc" + desc)
		}
	}
	var i = 0
	for desc[i] == '[' {
		i++
	}
	if desc[i] == 'L' {
		return desc[:i] + desc[i+1:len(desc)-1]
	}
	return desc
	// must be
}
