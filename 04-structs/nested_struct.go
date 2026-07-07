package main

import "fmt"

// inner struct
type Address struct {
    City    string
    Country string
}

// outer struct
type Student2 struct {
    Name    string
    Age     int
    Address Address // nested struct
}

func main() {

    s := Student2{
        Name: "Ali",
        Age:  20,
        Address: Address{
            City:    "Doha",
            Country: "Qatar",
        },
    }

    fmt.Println("Name:", s.Name)
    fmt.Println("City:", s.Address.City)
    fmt.Println("Country:", s.Address.Country)
}