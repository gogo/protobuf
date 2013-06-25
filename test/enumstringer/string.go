package enumstringer

func (this TheTestEnum) String() string {
	switch this {
	case 0:
		return "a"
	case 1:
		return "blabla"
	case 2:
		return "z"
	}
	return "3"
}
