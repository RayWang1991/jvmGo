package lib

import (
	"jvmGo/jvm/rtdt"
	"jvmGo/jvm/utils"
)

func init() {
	register(utils.CLASSNAME_AccessController, "doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", doPrivileged)
	register(utils.CLASSNAME_AccessController, "doPrivileged", "(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;", doPrivileged)
	register(utils.CLASSNAME_AccessController, "doPrivileged", "(Ljava/security/PrivilegedAction;Ljava/security/AccessControlContext;)Ljava/lang/Object;", doPrivileged)
	register(utils.CLASSNAME_AccessController, "getStackAccessControlContext", "()Ljava/security/AccessControlContext;", getStackAccessControlContext)
}

// @CallerSensitive
// public static native <T> T
//     doPrivileged(PrivilegedExceptionAction<T> action)
//     throws PrivilegedActionException;
// (Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;

// @CallerSensitive
// public static native <T> T doPrivileged(PrivilegedAction<T> action);
// (Ljava/security/PrivilegedAction;)Ljava/lang/Object;
func doPrivileged(frame *rtdt.Frame) {
	vars := frame.LocalVar
	action := vars.GetRef(0)

	stack := frame.OperandStack
	stack.PushRef(action)

	method := action.Class().GetMethodDirect("run", "()Ljava/lang/Object;")
	callMethod(method, frame)
}

// private static native AccessControlContext getStackAccessControlContext();
// ()Ljava/security/AccessControlContext;
func getStackAccessControlContext(frame *rtdt.Frame) {
	// todo
	frame.OperandStack.PushRef(nil)
}
