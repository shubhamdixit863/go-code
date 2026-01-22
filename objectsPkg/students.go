package objectsPkg

type Student struct{
	Name string
	Marks int
}

func (s Student)IsPassed()bool{
	return s.Marks>=40
}
func (s Student) Grade() string {
	if s.Marks >= 75 {
		return "A"
	} else if s.Marks >= 60 {
		return "B"
	} else if s.Marks >= 40 {
		return "C"
	} else {
		return "F"
	}
}
