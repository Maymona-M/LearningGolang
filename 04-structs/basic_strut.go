package main

import "fmt"

type Student struct {
    Name string
    Age  int
}

type Car struct {
	Make  string
	Model string
	Year  int
}

func main() {

    s := Student{
        Name: "Ali",
        Age:  20,
    }
	fmt.Println("Student Details:")
    fmt.Println(s.Name)
    fmt.Println(s.Age)

	fmt.Println("")

	c := Car{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2020,
	}
	fmt.Println("Car Details:")
	fmt.Println(c.Make)
	fmt.Println(c.Model)
	fmt.Println(c.Year)
}