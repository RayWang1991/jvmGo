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

	/*
	if !isAccessableMethod(from, m) {
		//debug
		fmt.Printf("IllegalAccess from %s method %s mclz %s %t\n", from.name, m.name, m.class.name, m.IsPublic())
		panic(utils.IllegalAccessError)
	}
	*/
	return m
}
