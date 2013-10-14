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

    for i := 1; i <= 120; i++ {
        fmt.Println("")
        fmt.Println(i)
        fmt.Println((i/100)%10)
        fmt.Println((i/10)%10)
        fmt.Println(i%10)

        if (i/100 > 0 && (i/100)%10 != 0) {
            fmt.Printf("%v", hundreds[(i/100)%10])
        }        

        if (i >= 20 && i/10 > 0 && (i/10)%10 != 0) {
            if (i >= 100) {
                fmt.Printf("%v", "and")
            }

            fmt.Printf("%v", tens[(i/10)%10])
        }

        if (i < 20) {
            fmt.Println(ones[i%20])
        } else {
            if (i%10 != 0) {
                fmt.Println(ones[i%10])
            }
        }
    }
}