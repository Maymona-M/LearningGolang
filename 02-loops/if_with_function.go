package main

import "fmt"

func getNumber() int {
    return 10
}

func main() {

    if num := getNumber(); num > 5 {
        fmt.Println("Big number:", num)
    }
}