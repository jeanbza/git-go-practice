package main

import (
    "fmt"
    "math/big"
)

func main() {
    veryBigNum := big.NewInt(1)
    two := big.NewInt(2)

    // Manually doing 2^1000
    for i:=0; i<1000; i++ { 
        temp := new(big.Int)
        temp.Mul(veryBigNum, two)
        veryBigNum = temp
    }

    veryBigNumString := veryBigNum.String()
    sum := 0

    for _, j := range veryBigNumString {
        sum += int(j-48) // Because range returns ASCII value, and 48 = 0
    }

    fmt.Println(sum)
}