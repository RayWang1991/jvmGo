package marea

import "jvmGo/jvm/utils"

func ndStr(name, desc string) string {
	return name + "&" + desc
}

func LookUpMethodVirtual(c, from *Class, name, desc string) *Method {
	m := c.LookUpMethod(name, desc)
	if m == nil {
		panic(utils.NoSuchMethodError)
	}

	if !isAccessableMethod(from, m) {
		panic(utils.IllegalAccessError)
	}
	return m
}
