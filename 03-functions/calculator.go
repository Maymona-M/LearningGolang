package main

import "fmt"

func adding(a, b int) int {
	return a + b
}

func subtracting(a, b int) int {
	return a - b
}

func multiplying(a, b int) int {
	return a * b
}

func dividing(a, b int) int {
	return a / b
}

func main() {

	var firstNum int
	var secondNum int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&firstNum)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&secondNum)

	fmt.Println("")

	fmt.Println("Press 1 for Addition")
	fmt.Println("Press 2 for Subtraction")
	fmt.Println("Press 3 for Multiplication")
	fmt.Println("Press 4 for Division")

	fmt.Println("")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	fmt.Println("")

	if choice == 1 {
		fmt.Println("Sum:", adding(firstNum, secondNum))
	} else if choice == 2 {
		fmt.Println("Difference:", subtracting(firstNum, secondNum))
	} else if choice == 3 {
		fmt.Println("Product:", multiplying(firstNum, secondNum))
	} else if choice == 4 {
		fmt.Println("Quotient:", dividing(firstNum, secondNum))
	}

}
