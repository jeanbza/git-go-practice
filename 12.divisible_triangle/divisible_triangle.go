package main

import (
    "fmt"
    "math"
)

func main() {
    targetNum := 1
    targetNumFactors := 500
    numFactorsAmnt := 0
    
    for i := 1; numFactorsAmnt < targetNumFactors; i++ {
        targetNum = generateTriangleNumber(i)
        numFactorsAmnt = numFactors(targetNum)
    }

    fmt.Printf("Smallest number with at least %v factors: %v\n\n", targetNumFactors, targetNum)
}

func generateTriangleNumber(someNum int) (int) {
    triangleNum := 0

    for i := 1; i <= someNum; i++ {
        triangleNum += i
    }

    return triangleNum
}

func numFactors(someNum int) (int) {
    numFactors := 0

    for i := 1; i <= int(math.Sqrt(float64(someNum))); i ++ {
        if (someNum % i == 0) {
            numFactors += 2
            if (someNum == 12) {
                fmt.Println(i)
            }
        }
    }

    if (isPerfectSquare(someNum)) {
        numFactors --
    }

    return numFactors
}

func isPerfectSquare(someInt int) (bool) {
    someFloat := float64(someInt)

    a := math.Sqrt(someFloat)
    b := float64(int(a))

    if (a == b) {
        return true
    } else {
        return false
    }
}