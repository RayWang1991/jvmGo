package classfile

import (
	"strings"
)

type ACC_TYPE int

const (
	ACC_TYPE_CLASS  ACC_TYPE = iota
	ACC_TYPE_FIELD
	ACC_TYPE_METHOD
)

// TODO distinguish class ,field and method acc

// ACCESS FLAG For basic

// All ACCESS FLAG
const (
	ACC_PUBLIC       uint16 = 0x0001
	ACC_FINAL               = 0x0010
	ACC_SUPER               = 0x0020
	ACC_INTERFACE           = 0x0200
	ACC_ABSTRACT            = 0x0400
	ACC_SYNTHETIC           = 0x1000
	ACC_ANNOTATION          = 0x2000
	ACC_ENUM                = 0x4000
	ACC_PRIVATE             = 0x0002
	ACC_PROTECTED           = 0x0004
	ACC_STATIC              = 0x0008
	ACC_SYNCHRONIZED        = 0x0020
	ACC_BRIDGE              = 0x0040
	ACC_VARARGS             = 0x0080
	ACC_NATIVE              = 0x0100
	ACC_STRICT              = 0x0800
	ACC_VOLATILE            = 0x0040
	ACC_TRANSIENT           = 0x0080
)

var flagMapClass = map[uint16]string{
	ACC_PUBLIC:     "ACC_PUBLIC",
	ACC_FINAL:      "ACC_FINAL",
	ACC_SUPER:      "ACC_SUPER",
	ACC_INTERFACE:  "ACC_INTERFACE",
	ACC_ABSTRACT:   "ACC_ABSTRACT",
	ACC_SYNTHETIC:  "ACC_SYNTHETIC",
	ACC_ANNOTATION: "ACC_ANNOTATION",
	ACC_ENUM:       "ACC_ENUM",
}

var flagMapField = map[uint16]string{
	ACC_PUBLIC:    "ACC_PUBLIC",
	ACC_PRIVATE:   "ACC_PRIVATE",
	ACC_PROTECTED: "ACC_PROTECTED",
	ACC_STATIC:    "ACC_STATIC",
	ACC_FINAL:     "ACC_FINAL",
	ACC_VOLATILE:  "ACC_VOLATILE",
	ACC_TRANSIENT: "ACC_TRANSIENT",
	ACC_SYNTHETIC: "ACC_SYNTHETIC",
	ACC_ENUM:      "ACC_ENUM",
}

var flagMapMethod = map[uint16]string{
	ACC_PUBLIC:       "ACC_PUBLIC",
	ACC_PRIVATE:      "ACC_PRIVATE",
	ACC_PROTECTED:    "ACC_PROTECTED",
	ACC_STATIC:       "ACC_STATIC",
	ACC_FINAL:        "ACC_FINAL",
	ACC_SYNCHRONIZED: "ACC_SYNCHRONIZED",
	ACC_BRIDGE:       "ACC_BRIDGE",
	ACC_VARARGS:      "ACC_VARARGS",
	ACC_NATIVE:       "ACC_NATIVE",
	ACC_ABSTRACT:     "ACC_ABSTRACT",
	ACC_STRICT:       "ACC_STRICT",
	ACC_SYNTHETIC:    "ACC_SYNTHETIC",
}

func flagNumToString(num uint16, acc_type ACC_TYPE) string {
	return flagsToString(decomposeFlags(num, acc_type), acc_type)
}

func getFlagMap(acc_type ACC_TYPE) map[uint16]string {
	switch acc_type {
	case ACC_TYPE_CLASS:
		return flagMapClass
	case ACC_TYPE_FIELD:
		return flagMapField
	case ACC_TYPE_METHOD:
		return flagMapMethod
	default:
		return nil
	}
}

func decomposeFlags(num uint16, acc_type ACC_TYPE) []uint16 {
	flagMap := getFlagMap(acc_type)
	res := make([]uint16, 0, len(flagMap))
	for f := range flagMap {
		if f&num != 0 {
			res = append(res, f)
		}
	}
	return res
}

func flagsToString(flags []uint16, acc_type ACC_TYPE) string {
	flagMap := getFlagMap(acc_type)
	strs := make([]string, 0, len(flags))
	for _, f := range flags {
		strs = append(strs, flagMap[f])
	}
	return strings.Join(strs, ", ")
}
