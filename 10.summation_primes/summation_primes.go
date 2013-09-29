package main

import (
	"fmt"
)

func main() {
	sum := int64(2)

	for i := 3; i < 2000000; i+=2 {
		if (isPrime(i)) {
			sum += int64(i)
		}
	}

	fmt.Printf("\n\nSum: %v\n", sum)
}

func isPrime(someNum int) (bool) {
    for j := 3; j < someNum-1; j+=2 {
        if (someNum % j == 0) {
            return false
        }
    }

    return true
}