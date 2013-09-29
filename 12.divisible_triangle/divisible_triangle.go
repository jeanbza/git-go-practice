package main

import (
    "fmt"
    "math"
)

func main() {
    currentTriangleNum := 1
    targetNumDivisors := 500
    numDivisors := 0
    
    for i := 1; numDivisors < targetNumDivisors; i++ {
        currentTriangleNum = generateTriangleNumber(i)
        numDivisors = getNumDivisors(currentTriangleNum)
    }

    fmt.Printf("Smallest number with at least %v factors: %v\n\n", targetNumDivisors, currentTriangleNum)
}

func generateTriangleNumber(someNum int) (int) {
    triangleNum := 0

    for i := 1; i <= someNum; i++ {
        triangleNum += i
    }

    return triangleNum
}

func getNumDivisors(someNum int) (int) {
    getNumDivisors := 0

    for i := 1; i <= int(math.Sqrt(float64(someNum))); i ++ {
        if (someNum % i == 0) {
            getNumDivisors += 2
            if (someNum == 12) {
                fmt.Println(i)
            }
        }
    }

    if (isPerfectSquare(someNum)) {
        getNumDivisors --
    }

    return getNumDivisors
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