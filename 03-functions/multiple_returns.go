package main

import "fmt"

func divide(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func swapped(a int, b int) (int, int) {
	return b, a
}

func main() {
    q, r := divide(10, 3)
    fmt.Println("Quotient:", q)
    fmt.Println("Remainder:", r)

	swappedA, swappedB := swapped(10, 2)
	fmt.Println(swappedA, swappedB)
}