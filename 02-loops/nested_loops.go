package main

import "fmt"

func main() {

    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Println("i:", i, "j:", j)
        }
    }
}