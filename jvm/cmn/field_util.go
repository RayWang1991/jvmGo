package cmn

func SlotNum(desc string) uint8 {
	switch desc {
	case "D", "J":
		return 2
	default:
		//  "B", "C", "F", "I", "S", "Z", "L;","["
		return 1
	}
}
