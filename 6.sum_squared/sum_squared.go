package main

import (
	"fmt"
	"math"
)

func main() {
    squaredSum := float64(0)
    sumSquared := float64(0)

    for i := float64(0); i <= 100; i++ {
        squaredSum += math.Pow(i,2)
        sumSquared += i
    }

    sumSquared = math.Pow(sumSquared,2)
   
    fmt.Printf("Total: %v \n", int32(sumSquared-squaredSum))
}