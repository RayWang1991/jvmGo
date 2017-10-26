package marea

import "jvmGo/jvm/cmn"

func (c *Class) GetMain() *Method {
	return c.GetMethodDirect(cmn.METHOD_MAIN_NAME, cmn.METHOD_MAIN_DESC)
}

func (c *Class) GetClinit() *Method {
	return c.GetMethodDirect(cmn.METHOD_CLINIT_NAME, cmn.METHOD_CLINIT_DESC)
}
