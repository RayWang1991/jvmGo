package classpath

type ClassPath struct{
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

// parse ClassPath entity from jre op and 
func Parse(jreOption, cpOption string)*ClassPath{
}
