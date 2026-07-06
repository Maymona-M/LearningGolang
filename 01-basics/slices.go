package main

import "fmt"

func main() {

    slice := []int{1, 2, 3}

    slice = append(slice, 4)

    fmt.Println(slice)
}