package main

import (
	"fmt"
	"learninggolang/06-packages/calculator"
	"learninggolang/06-packages/student"
)

func main() {

	resultAdd := calculator.Add(5, 3)
	resultSubtract := calculator.Subtract(5, 3)
	resultMultiply := calculator.Multiply(5, 3)
	resultDivideQuotient, resultDivideRemainder := calculator.Divide(10, 3)

	fmt.Println("Added result:", resultAdd)
	fmt.Println("Subtracted result:", resultSubtract)
	fmt.Println("Multiplied result:", resultMultiply)
	fmt.Println("Divided result: Quotient =", resultDivideQuotient, ", Remainder =", resultDivideRemainder)


	s := student.Student{
		Name:"Ali",
		Age:20,
		School:"XYZ University",
	}


	fmt.Println(s.Introduce())

}
