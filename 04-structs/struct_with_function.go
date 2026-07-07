package main

import "fmt"

type Student1 struct {
	Name string
	Age  int
}

func printStudent(s Student1) {
	fmt.Println(s.Name, s.Age)
}

type School struct {
	Name     string
	Location string
	Capacity int
}

func printSchool(sch School) {
	fmt.Println(sch.Name, sch.Location, sch.Capacity)
}

func main() {

	s := Student1{"Sara", 22}
	printStudent(s)

	sch := School{"ABC International", "Qatar", 1500}
	printSchool(sch)
}
