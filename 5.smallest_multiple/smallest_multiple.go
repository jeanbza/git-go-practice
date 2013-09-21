package main

import (
    "os"
    "fmt"
)

func main() {
    x := 1

    for true {
        i := 2*3*5*7*11*13*17*19*x

        if (i % 20 == 0 && i % 19 == 0 && i % 18 == 0 && i % 17 == 0 && i % 16 == 0 && i % 15 == 0 && i % 14 == 0 && i % 13 == 0 && i % 12 == 0 && i % 11 == 0 && i % 10 == 0) {
            fmt.Printf("Number IS: %v\n", i)
            os.Exit(0)
        } else {
            fmt.Printf("Number is NOT: %v\n", i)
        }

        x ++
    }
}