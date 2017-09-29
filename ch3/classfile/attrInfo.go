package classfile

type AttrInfo interface {
	// ReadInfo should read info from the byte stream and constant pool, and returns the should be byte numbers
	ReadInfo(reader ClassReader) uint64
}

type AttrInfoBase struct {
	nameIndex uint16
	length    uint64
}

const (
	ATTRNAME_CONSTVALUE                          = "ConstantValue"
	ATTRNAME_CODE                                = "Code"
	ATTRNAME_EXCEPTIONS                          = "Exceptions"
	ATTRNAME_INNERCLASSES                        = "InnerClasses"
	ATTRNAME_ENCLOSINGMETHOD                     = "EnclosingMethod"
	ATTRNAME_SYNTHETIC                           = "Synthetic"
	ATTRNAME_DEPRECATED                          = "Deprecated"
	ATTRNAME_SIGNATURE                           = "Signature"
	ATTRNAME_SOURCEFILE                          = "SourceFile"
	ATTRNAME_LINENUMBERTABLE                     = "LineNumberTable"
	ATTRNAME_LOCALVARIAVLETABLE                  = "LocalVariableTable"
	ATTRNAME_BOOTSTRAPMETHODS                    = "BootstrapMethods"
	ATTRNAME_METHODPARAMETERS                    = "MethodParameters"
	ATTRNAME_STACKMAPTABLE                       = "StackMapTable"                        // unsupported
	ATTRNAME_LOVALVARIAVLETYPETABLE              = "LocalVariableTypeTable"               // unsupported
	ATTRNAME_SOURCEDEBUGEXTENTION                = "SourceDebugExtension"                 // unsupported
	ATTRNAME_RUNTIMEVISIBLEANNOTATIONS           = "RuntimeVisibleAnnotations"            // unsupported
	ATTRNAME_RUNTIMEINVISIBLEANNOTATIONS         = "RuntimeInvisibleAnnotations"          // unsupported
	ATTRNAME_RUNTIMEVISIBLEPARAMTERANNOTATIONS   = "RuntimeVisibleParameterAnnotations"   // unsupported
	ATTRNAME_RUNTIMEINVISIBLEPARAMTERANNOTATIONS = "RuntimeInvisibleParameterAnnotations" // unsupported
	ATTRNAME_RUNTIMEVISIBLETYPEANNATATIONS       = "RuntimeVisibleTypeAnnotations"        // unsupported
	ATTRNAME_RUNTIMEINVISIBLETYPEANNATATIONS     = "RuntimeInvisibleTypeAnnotations"      // unsupported
	ATTRNAME_ANNOTATIONDEFAULT                   = "AnnotationDefault"                    // unsupported
)

// factory method for attribute info
func NewAttributeInfo(reader ClassReader, cp ConstantPool) AttrInfo {
	nameIndex := reader.ReadUint16()
	name := cp.getUtf8(nameIndex)
	switch name {
	case ATTRNAME_CONSTVALUE:
		return &AttrConstantValue{cp: cp}
	case ATTRNAME_CODE:
		return &AttrCode{cp: cp}
	case ATTRNAME_EXCEPTIONS:
		return &AttrExceptions{cp: cp}
	case ATTRNAME_INNERCLASSES:
		return &AttrInnerClasses{cp: cp}
	case ATTRNAME_ENCLOSINGMETHOD:
		return &AttrEnclosingMethod{cp: cp}
	case ATTRNAME_SYNTHETIC:
		return &AttrSynthetic{}
	case ATTRNAME_DEPRECATED:
		return &AttrDeprecated{}
	case ATTRNAME_SIGNATURE:
		return &AttrSignature{cp: cp}
	case ATTRNAME_SOURCEFILE:
		return &AttrSourceFile{cp: cp}
	case ATTRNAME_LINENUMBERTABLE:
		return &AttrLineNumberTable{}
	case ATTRNAME_LOCALVARIAVLETABLE:
		return &AttrLocalVarTable{}
	case ATTRNAME_BOOTSTRAPMETHODS:
		return &AttrBootstrapMethods{cp: cp}
	case ATTRNAME_METHODPARAMETERS:
		return &AttrMethodParameters{cp: cp}
	default:
		return &AttrUnsupported{}
	}
}
