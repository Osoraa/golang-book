package main

// Squares an integer in place

import "fmt"

func square(x *float64) {   
    *x = *x * *x
}

func main() {
    x := 1.5
    square(&x)

    // Prints 2.25
    fmt.Println(x)
}
