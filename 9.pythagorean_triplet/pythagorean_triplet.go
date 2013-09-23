package main

import (
    "math"
	"fmt"
)

func main() {
    for a := 3; a < 1000; a++ {
        for b := a+1; b < 1000; b++ {
            cSquared := int(math.Pow(float64(a),2)+math.Pow(float64(b),2))
            c := int(math.Sqrt(float64(cSquared)))
            if (isPerfectSquare(cSquared) && a+b+c == 1000) {
                fmt.Printf("%v^2 + %v^2 = %v^2\n", a, b, c)
                fmt.Printf("%v*%v*%v=%v\n\n", a, b, c, a*b*c)
            }
        }
    }
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