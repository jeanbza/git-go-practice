package main

import "fmt"

func main() {
    index := 1
    targetNum := 600851475143

    for index <= targetNum {
        if (isPrime(index) && targetNum % index == 0) {
            fmt.Println(index)
            targetNum /= index
        }

        index ++
    }
}

/**
 * As it turns out, this is not a correct isPrime function. Not sure how this program returned the correct answer.
 * A correct implementation can be found in 7.1000st_prime
 */
func isPrime(someNum int) (bool) {
    index := someNum-1

    for index > 0 {
        if (someNum % index == 0) {
            return true
        }
        index --;
    }

    return false
}