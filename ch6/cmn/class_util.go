package cmn

import "strings"

func IsArray(name string) bool {
	return strings.HasSuffix(name, "[")
}
