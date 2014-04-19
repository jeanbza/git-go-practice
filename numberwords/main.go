package main

import (
    "runtime"
    "strconv"
    "log"
)

func main() {
    inputFlt := 2519.04
    inputStr := strconv.FormatFloat(inputFlt, 'f', 2, 64)
    numberWords := [22]string{"", "", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
                              "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
                              "eighteen", "nineteen"}
    
    var sentence string
    var pos int
    var lastTwo int64
    var err error

    for i, c := range inputStr {
        if (i <= len(inputStr)-4) {
            pos = len(inputStr)-4-i

            switch (pos) {
                case 1:
                    sentence += " and "
                    
                    lastTwo, err = strconv.ParseInt(string(inputStr[i])+string(inputStr[i+1]), 10, 64)
                    checkError(err)

                    sentence += numberWords[lastTwo+2]
                case 2:
                    sentence += numberWords[c-46]+" hundred "
                case 3:
                    sentence += numberWords[c-46]+" thousand "
            }
        }
    }

    log.Println(inputFlt)
    log.Println(sentence)
}

// Stacktrace functionality for error handling
func checkError(err error) {
    if err != nil {
        var stack [4096]byte
        runtime.Stack(stack[:], false)
        log.Printf("%q\n%s\n", err, stack[:])
    }
}