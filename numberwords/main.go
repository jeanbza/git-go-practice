package main

import (
    "strconv"
    "log"
)

func main() {
    inputFlt := 2523.04
    inputStr := strconv.FormatFloat(inputFlt, 'f', 2, 64)
    numberWords := [12]string{"", "", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    for _, c := range inputStr {
        log.Println(numberWords[c-46])
    }
}