package marea

import (
	"jvmGo/jvm/cmn"
)

var strPool = make(map[string]*Object)

func GetJavaString(key string, loader ClassLoader) *Object {
	if v := strPool[key]; v != nil {
		return v
	}
	sClass := loader.Load("java/lang/String")
	csClass := loader.Load("[C")

	craw := cmn.UTF8ToUTF16(key)
	c := NewArrayC(csClass, int32(len(craw)))

	ca := c.ArrGetChars()
	for i := range ca {
		ca[i] = craw[i]
	}

	s := NewObject(sClass)
	sf := sClass.GetFieldDirect("value", "[C")
	s.SetRef(c, sf.VarIdx())

	strPool[key] = s
	return s
}

func GetGoString(str *Object) string {
	cArray := str.GetInsFieldRef("value")
	return cmn.UTF16ToUTF8(cArray.ArrGetChars())
}

func GetInternString(javaString *Object) *Object {
	goStr := GetGoString(javaString)
	if str, ok := strPool[goStr]; ok {
		return str
	} else {
		strPool[goStr] = javaString
		return javaString
	}
}
