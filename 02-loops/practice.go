package main

import "fmt"

func main() {

	// print all even numbers from 1 to 20
	var even = []int{}
	for i:= 1; i<=20; i++ {
		if i%2 == 0 {
			even = append(even, i)
		}
	}
	fmt.Println("Even numbers:", even)

	fmt.Println("")

	// sum all numbers from 1 to 10
	sum := 0
	for i:=1; i<=10; i++ {
		sum += i
	}
	fmt.Println("Sum = ", sum)

	fmt.Println("")

	// find max in a list
	numbers := []int{3, 5, 2, 8, 1}
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("Max number:", max)

	fmt.Println("")

	// reverse a loop
	fmt.Println("Original Loop:", numbers)
	var rev_numbers = []int{}
	for i := len(numbers) - 1; i >= 0; i-- {
		rev_numbers = append(rev_numbers, numbers[i])
	}
	fmt.Println("Reversed Loop:", rev_numbers)
}