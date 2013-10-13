package main

import "fmt"

var numbers map[int]string

func main() {
    numbers = make(map[int]string)
    numbers[1] = "one"
    numbers[2] = "two"
    numbers[3] = "three"
    numbers[4] = "four"
    numbers[5] = "five"
    numbers[6] = "six"
    numbers[7] = "seven"
    numbers[8] = "eight"
    numbers[9] = "nine"
    numbers[10] = "ten"
    numbers[20] = "twenty"
    numbers[30] = "thirty"
    numbers[40] = "forty"
    numbers[50] = "fifty"
    numbers[60] = "sixty"
    numbers[70] = "seventy"
    numbers[80] = "eighty"
    numbers[90] = "ninety"
    numbers[100] = "onehundred"
    numbers[200] = "twohundred"
    numbers[300] = "threehundred"
    numbers[400] = "fourhundred"
    numbers[500] = "fivehundred"
    numbers[600] = "sixhundred"
    numbers[700] = "sevenhundred"
    numbers[800] = "eighthundred"
    numbers[900] = "ninehundred"
    numbers[1000] = "onethousand"

    fmt.Println(numbers[1])
}