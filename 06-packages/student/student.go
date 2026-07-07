package student

type Student struct {
	Name string
	Age int
	School string
}


func (s Student) Introduce() string {

	return "Hi " + s.Name
}