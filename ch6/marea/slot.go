package marea

type Slot struct {
	Num int32   // for return addr, boolean, byte, char, int
	Ref *Object // for reference type
	// for double, and long, it takes two slot
}

func (s *Slot) Copy() *Slot {
	return &Slot{s.Num, s.Ref}
}

