package classfile

type AttrInfo interface {
	// ReadInfo should read info from the byte stream and constant pool, and returns the should be byte numbers
	ReadInfo(reader ClassReader, cp ConstantPool) uint64
}

type AttrInfoBase struct {
	nameIndex uint16
	length    uint64
}

const (
	ATTRNAME_CONSTVALUE                          = "ConstantValue"
	ATTRNAME_CODE                                = "Code"
	ATTRNAME_STACKMAPTABLE                       = "StackMapTable" // unsupported
	ATTRNAME_EXCEPTIONS                          = "Exceptions"
	ATTRNAME_INNERCLASSES                        = "InnerClasses"
	ATTRNAME_ECLOSINGMETHOD                      = "EnclosingMethod"
	ATTRNAME_SYNTHETIC                           = "Synthetic"
	ATTRNAME_SIGNATURE                           = "Signature"
	ATTRNAME_SOURCEFILE                          = "SourceFile"
	ATTRNAME_SOURCEDEBUGEXTENTION                = "SourceDebugExtension"
	ATTRNAME_LINENUMBERTABLE                     = "LineNumberTable"
	ATTRNAME_LOCALVARIAVLETABLE                  = "LocalVariableTable"
	ATTRNAME_LOVALVARIAVLETYPETABLE              = "LocalVariableTypeTable"
	ATTRNAME_DEPRECATED                          = "Deprecated"                           // unsupported
	ATTRNAME_RUNTIMEVISIBLEANNOTATIONS           = "RuntimeVisibleAnnotations"            // unsupported
	ATTRNAME_RUNTIMEINVISIBLEANNOTATIONS         = "RuntimeInvisibleAnnotations"          // unsupported
	ATTRNAME_RUNTIMEVISIBLEPARAMTERANNOTATIONS   = "RuntimeVisibleParameterAnnotations"   // unsupported
	ATTRNAME_RUNTIMEINVISIBLEPARAMTERANNOTATIONS = "RuntimeInvisibleParameterAnnotations" // unsupported
	ATTRNAME_RUNTIMEVISIBLETYPEANNATATIONS       = "RuntimeVisibleTypeAnnotations"        // unsupported
	ATTRNAME_RUNTIMEINVISIBLETYPEANNATATIONS     = "RuntimeInvisibleTypeAnnotations"      // unsupported
	ATTRNAME_ANNOTATIONDEFAULT                   = "AnnotationDefault"                    // unsupported
	ATTRNAME_BOOTSTRAPMETHODS                    = "BootstrapMethods"
	ATTRNAME_METHODPARAMETERS                    = "MethodParameters"
)

// factory method for attribute info
func NewAttributeInfo(reader ClassReader, cp ConstantPool) AttrInfo {
	nameIndex := reader.ReadUint16()
	name := cp.getUtf8(nameIndex)
	switch name {
	case ATTRNAME_CONSTVALUE:
		return &AttrConstantValue{}
	case ATTRNAME_CODE:
		return &AttrCode{}
	}
}
