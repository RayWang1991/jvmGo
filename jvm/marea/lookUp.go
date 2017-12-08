package marea

import (
	"jvmGo/jvm/utils"
	"fmt"
)

func ndStr(name, desc string) string {
	return name + "&" + desc
}

func LookUpMethodVirtual(c, from *Class, name, desc string) *Method {
	utils.Dprintf("[LOOK UP VIRTUAL] clz %s, from %s, name %s\n", c.name, from.name, name)
	m := c.LookUpMethod(name, desc)
	if m == nil {
		//debug
		fmt.Printf("Not Find Virtual %s %s from %s in %s\n", name, desc, from.name, c.name)
		panic(utils.NoSuchMethodError)
	}

	if !isAccessableMethod(from, m) {
		panic(utils.IllegalAccessError)
	}
	return m
}
