package main

import "fmt"

var ones map[int]string
var tens map[int]string
var hundreds map[int]string
var thousands map[int]string

func main() {
    ones = make(map[int]string)
    tens = make(map[int]string)
    hundreds = make(map[int]string)
    thousands = make(map[int]string)

    ones[1] = "one"
    ones[2] = "two"
    ones[3] = "three"
    ones[4] = "four"
    ones[5] = "five"
    ones[6] = "six"
    ones[7] = "seven"
    ones[8] = "eight"
    ones[9] = "nine"
    ones[10] = "ten"
    ones[11] = "eleven"
    ones[12] = "twelve"
    ones[13] = "thirteen"
    ones[14] = "fourteen"
    ones[15] = "fifteen"
    ones[16] = "sixteen"
    ones[17] = "seventeen"
    ones[18] = "eighteen"
    ones[19] = "nineteen"
    tens[2] = "twenty"
    tens[3] = "thirty"
    tens[4] = "forty"
    tens[5] = "fifty"
    tens[6] = "sixty"
    tens[7] = "seventy"
    tens[8] = "eighty"
    tens[9] = "ninety"
    hundreds[1] = "onehundred"
    hundreds[2] = "twohundred"
    hundreds[3] = "threehundred"
    hundreds[4] = "fourhundred"
    hundreds[5] = "fivehundred"
    hundreds[6] = "sixhundred"
    hundreds[7] = "sevenhundred"
    hundreds[8] = "eighthundred"
    hundreds[9] = "ninehundred"
    thousands[1] = "onethousand"

    first := 0
    second := 0
    third := 0

    for i := 1; i <= 1000; i++ {
        fmt.Println("")
        fmt.Println(i)
        first = ((i/100)%10)

        if ((i/10)%10 < 2) {
            second = 0
            third = i%20
        } else {
            second = (i/10)%10
            third = i%10
        }

        if (i == 1000) {
            fmt.Print("onethousands")
        }

        fmt.Print(hundreds[first])

        if (first > 0 && second > 0) {
            fmt.Print("and")
        }

        fmt.Print(tens[second])
        fmt.Print(ones[third])
        fmt.Println()

        // if (first > 0) {
        //     fmt.Printf("%v", hundreds[first])

        //     if (second != 0 && third != 0) {
        //         fmt.Printf("%v", "and")
        //     }
        // }

        // j /= 10

        // if (j >= 20) {
        //     fmt.Printf("%v", tens[j%10])
        // }

        // j /= 10

        // if (j > 0) {
        //     fmt.Printf("%v", ones[j])
        // }
    }
}