package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) introduce() {
	fmt.Println("My name is", s.Name)
	fmt.Println("My age is", s.Age)
}

func main() {

	student := Student{
		Name: "Ali",
		Age:  20,
	}

	student.introduce()
}