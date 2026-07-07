package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("example.txt")
	// returns two values, the file and an error. If the file does not exist, err will be non-nil.

	if err != nil {
		fmt.Println("Something went wrong:")
		fmt.Println(err)
		return
	}

	fmt.Println("File opened successfully")

	file.Close()
}