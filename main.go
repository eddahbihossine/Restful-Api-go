package main

// typecasting in go
// typecasting is the process of converting one data type to another data type
// you go int(a)

import "fmt"

func calculate(a float64 ,b int ) int{
    return int(a) + b
}


func main() {
    fmt.Println("Hello, world!")
    fmt.Println("The sum of 10 and 20 is: ", calculate(1.2,3))
}
