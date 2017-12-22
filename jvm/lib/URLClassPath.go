package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_URLClassPath, "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;", getLookupCacheURLs)
}

// private static native URL[] getLookupCacheURLs(ClassLoader var0);
// (Ljava/lang/ClassLoader;)[Ljava/net/URL;
func getLookupCacheURLs(frame *rtdt.Frame) {
	//TODO
	frame.OperandStack.PushRef(nil)
}
