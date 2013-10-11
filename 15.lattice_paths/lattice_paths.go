package main

import (
    "fmt"
    "math"
)

func main() {
    sum := factorial((float64(2*20)))/math.Pow((factorial(20)),2)

    fmt.Println(sum)
}

func factorial(someNum float64) (float64) {
    product := float64(1)

    for i := float64(1); i <= someNum; i++ {
        product *= i
    }

    return product
}