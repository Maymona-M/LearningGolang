package main

import "fmt"

func add(a int, b int) {
    fmt.Println("Sum of ", a, "and ", b)
	fmt.Println(" = ", a+b)
}

func subtract(a int, b int) {
    fmt.Println("Difference of ", a, "and ", b)
	fmt.Println(" = ", a-b)
}

func swap(a int, b int) {
	fmt.Println("Before Swap: a =", a, ", b =", b)
	a, b = b, a
	fmt.Println("After Swap: a =", a, ", b =", b)
}

func main() {
    add(3, 5)
	subtract(10, 4)
	swap(7,2)
}