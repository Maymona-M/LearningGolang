package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) changeName(newName string) {
	s.Name = newName
}

func (s *Student) changeAge(newAge int) {
	s.Age = newAge
}

func main() {

	s := Student{
		Name: "Ali",
		Age:  20,
	}

	s.changeName("Ahmed")
	s.changeAge(25)

	fmt.Println(s.Name)
	fmt.Println(s.Age)
}
