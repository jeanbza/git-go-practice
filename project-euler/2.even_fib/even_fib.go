package main

import "fmt"

func main() {
    a := 1
    b := 1
    c := a+b
    max := 4000000
    sum := 0

    for a <= max && b < max {
        if (c % 2 == 0) {
        	sum += c
        	fmt.Printf("%v \n", c)
        }

        a = b
        b = c 
        c = a+b
    }
   
    fmt.Printf("Total: %v \n", sum)
}