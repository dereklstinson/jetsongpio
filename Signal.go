package jetsongpio

//Signal is used to set signals and for returned signals.  Flags can be set using its methods
type Signal int

//CreateSignal creates an unidentified signal.  It is just a convenence function.
//If wanting to set a wanting to create a signal flag then use flg:=CreateSignal()
//These flags are not thread safe.  So don't pass pointers of them.
func CreateSignal() Signal {
	return Signal(0)
}
func sig() Signal {
	return Signal(0)
}

//LOW sets signal to LOW and returns LOW.
//
//LOW is used for general GPIO uses like checking if a pin is LOW or setting a pin LOW
func (s *Signal) LOW() Signal {
	*s = 1
	return *s
}

//HIGH sets signal to HIGH and returns HIGH.
//
//HIGH is used for general GPIO uses like checking if a pin is HIGH or setting a pin HIGH
func (s *Signal) HIGH() Signal {
	*s = 2
	return *s
}

//RISING sets signal to RISING and returns RISING.
//
//RISING is used for interupts or polling
func (s *Signal) RISING() Signal {
	*s = 3
	return *s
}

//FALLING sets signal to FALLING and returns FALLING.
//
//FALLING is used for interupts or polling
func (s *Signal) FALLING() Signal {
	*s = 4
	return *s
}

//NONE sets signal to NONE and returns NONE.
//
//NONE is used for interupts or polling
func (s *Signal) NONE() Signal {
	*s = 5
	return *s
}

//BOTH sets signal to BOTH and returns BOTH.
//
//BOTH  is used for interupts or polling
func (s *Signal) BOTH() Signal {
	*s = 6
	return *s
}
func (s Signal) String() string {
	flg := s
	switch s {
	case flg.HIGH():
		return "HIGH"
	case flg.LOW():
		return "LOW"
	case flg.RISING():
		return "RISING"
	case flg.FALLING():
		return "FALLING"
	case flg.NONE():
		return "NONE"
	case flg.BOTH():
		return "BOTH"
	default:
		return "Undefined"
	}
}
