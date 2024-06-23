<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
package Homework01

func main() {
	// TODO тут напишите цикл с вызовом FizzBuzz
	// fmt.Println(fizzbuzz.FizzBuzz(10))
=======
package main

// type Smth func(s string) string

// func main() {
// 	s := suf("Нет!")
// 	println(s("да"))
// 	println(s("Да"))
// }

// func suf(suf string) func(s string) string {
// 	return func(s string) string {
// 		return s + suf
// 	}
// }

import (
	"fmt"
	"strings"
)

type Mod func(s string) string

func main() {
	variable := "X"
	fmt.Println(applyer(variable, IncreasingValue))
	twoDiagonals(variable)
	fmt.Println(applyer(variable, IncreasingValue))
}

// // from lowercase to uppercase "x" -> "X"
// func upper(s string) string {
// 	return strings.ToUpper(s)
// }

// repeats a string (s) 15 times
func IncreasingValue(s string) string {
	repeated := strings.Repeat(s, 15) // I used stackoverflow for this line
	return repeated
}

func applyer(s string, mods ...Mod) string {
	for _, mod := range mods {
		s = mod(s)
	}
	return s
}

func twoDiagonals(s string) {
	diagonal := 13                  // this is the length of diagonals
	for i := 0; i < diagonal; i++ { // for a row
		for j := 0; j < diagonal; j++ { // for a column
			// if the 1st condition is okay then add "X" to the top-left to button-right
			// if the 2nd condition is okay then "X" to the top-right to button-left
			if j == i || j == diagonal-i-1 {
				fmt.Print(s) // if the condition is met, print "X"
			} else {
				fmt.Print(" ") // if the condition isn't met, print empty " "
			}
		}
		fmt.Println() // moves to the next line at the end of the rows
	}
>>>>>>> 6f8c6c7 (Prep for the 1st homework from the second course)
=======
package Homework01

func main() {
	// TODO тут напишите цикл с вызовом FizzBuzz
	// fmt.Println(fizzbuzz.FizzBuzz(10))
>>>>>>> 12c2a32 (Prep for the 1st homework from the second course)
=======
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
>>>>>>> 588f809 (practical part form the 2nd course lecture 1)
=======
package Homework01

func main() {
	// TODO тут напишите цикл с вызовом FizzBuzz
	// fmt.Println(fizzbuzz.FizzBuzz(10))
>>>>>>> 5312f43 (I solved the 1st hw, but I've not realized how to add color there)
}
