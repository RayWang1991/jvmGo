package rtdata

type slot struct {
	num int32   // for return addr, boolean, byte, char, int
	ref *Object // for reference type
	// for double, and long, it takes two slot
}

func (s *slot) Copy() *slot {
	return &slot{s.num, s.ref}
}
