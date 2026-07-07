package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Program started...")

	time.Sleep(3 * time.Second)

	fmt.Println("3 seconds later")

	current := time.Now()

	fmt.Println("Current time:", current)

	day := current.Day()
	month := current.Month()
	year := current.Year()

	fmt.Printf("Current date: %02d-%02d-%d\n", day, month, year)

}