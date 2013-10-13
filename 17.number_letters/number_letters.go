package main

import "fmt"

var numbers map[int]string

func main() {
	numbers = make(map[int]string)
	numbers[1] = "one"

	fmt.Println(numbers[1])
}