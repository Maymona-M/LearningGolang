package main

import "fmt"

type Student struct {
	Name string
	Age int
}


func (s Student) String() string {
	return fmt.Sprintf("%s (%d years old)", s.Name, s.Age)
}


func main(){

	s := Student{
		Name:"Ali",
		Age:20,
	}


	fmt.Println(s)

}