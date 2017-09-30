package classfile

import (
	"strings"
)

const (
	ACC_PUBLIC     uint16 = 0x0001
	ACC_FINAL             = 0x0010
	ACC_SUPER             = 0x0020
	ACC_INTERFACE         = 0x0200
	ACC_ABSTRACT          = 0x0400
	ACC_SYNTHETIC         = 0x1000
	ACC_ANNOTATION        = 0x2000
	ACC_ENUM              = 0x4000
)

var flagMap = map[uint16]string{
	ACC_PUBLIC:     "ACC_PUBLIC",
	ACC_FINAL:      "ACC_FINAL",
	ACC_SUPER:      "ACC_SUPER",
	ACC_INTERFACE:  "ACC_INTERFACE",
	ACC_ABSTRACT:   "ACC_ABSTRACT",
	ACC_SYNTHETIC:  "ACC_SYNTHETIC",
	ACC_ANNOTATION: "ACC_ANNOTATION",
	ACC_ENUM:       "ACC_ENUM",
}

func flagNumToString(num uint16) string {
	return flagsToString(decomposeFlags(num))
}

func decomposeFlags(num uint16) []uint16 {
	res := make([]uint16, 0, len(flagMap))
	for f := range flagMap {
		if f&num != 0 {
			res = append(res, f)
		}
	}
	return res
}

func flagsToString(flags []uint16) string {
	strs := make([]string, 0, len(flags))
	for _, f := range flags {
		strs = append(strs, flagMap[f])
	}
	return strings.Join(strs, ", ")
}
