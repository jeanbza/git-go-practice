package main

import "fmt"

func main() {
    primeNums := 0
    index := 1

    for primeNums <= 10001 {
        if (isPrime(index)) {
            fmt.Printf("%v more to go\n", 10001-primeNums)
            fmt.Printf("%v is the prime\n\n", index)
            primeNums ++
        }

        index ++
    }
}

func isPrime(someNum int) (bool) {
    index := someNum-1

    for index > 1 {
        if (someNum % index == 0) {
            return false
        }
        index --;
    }

    return true
}