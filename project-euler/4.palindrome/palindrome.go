package main

import (
    "strconv"
    "fmt"
)

func main() {
    largest := 0
    a := 999
    for a > 99 {
        b := 999
        for b > 99 {
            if (isPalindrome(a*b) && a*b > largest) {
                largest = a*b
                fmt.Printf("%v x %v = %v\n", a, b, largest)
            }
            b --
        }
        a --
    }
    
}

func isPalindrome(someNum int) (bool) {
    someStr := strconv.Itoa(someNum)
    runes := []rune(someStr)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if (runes[i] != runes[j]) {
            return false
        }
    }

    return true
}