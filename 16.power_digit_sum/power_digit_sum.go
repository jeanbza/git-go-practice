package main

import (
    "fmt"
    "math/big"
)

func main() {
    veryBigNum := big.NewInt(1)
    two := big.NewInt(2)

    for i:=0; i<1000; i++ {
        temp := new(big.Int)
        temp.Mul(veryBigNum, two)
        veryBigNum = temp
    }

    veryBigNumString := veryBigNum.String()
    sum := 0

    for _, j := range veryBigNumString {
        sum += int(j-48) // Because range returns ASCII table, and 48 = 0
    }

    fmt.Println(sum)
}