package classfile

type AttrInnerClasses struct {
	cp ConstantPool
	classes []AttrInnerClass
}

type AttrInnerClass struct {
	innerClassIndex       uint16
	outerClassIndex       uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func (innerClasses *AttrInnerClasses) ReadInfo(reader *ClassReader) uint32 {
	num := reader.ReadUint32()
	classNum := reader.ReadUint16()
	classes := make([]AttrInnerClass, 0, classNum)
	var i uint16
	for i = 0; i < classNum; i ++ {
		class := AttrInnerClass{}
		class.ReadInfo(reader)
		classes = append(classes, class)
	}
	innerClasses.classes = classes
	return num
}

func (innerClass *AttrInnerClass) ReadInfo(reader *ClassReader) {
	innerClass.innerClassIndex = reader.ReadUint16()
	innerClass.outerClassIndex = reader.ReadUint16()
	innerClass.innerNameIndex = reader.ReadUint16()
	innerClass.innerClassAccessFlags = reader.ReadUint16()
}
