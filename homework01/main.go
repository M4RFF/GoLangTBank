package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	nums := make([]int, 100)

	// we are ranging numbers
	for i := range nums {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
