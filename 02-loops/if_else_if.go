package main

import "fmt"

func main() {

    marks := 75

    if marks >= 90 {
        fmt.Println("A grade")
    } else if marks >= 75 {
        fmt.Println("B grade")
    } else if marks >= 50 {
        fmt.Println("C grade")
    } else {
        fmt.Println("Fail")
    }
}