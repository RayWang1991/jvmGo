package marea

func (c *Class) GetMain() *Method {
	return c.GetMethodDirect("main", "([Ljava/lang/String;)V")
}

func (c *Class) GetCInit() *Method {
	return c.GetMethodDirect("<clinit>", "()V")
}
