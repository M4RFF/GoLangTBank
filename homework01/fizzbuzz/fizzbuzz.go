package fizzbuzz
<<<<<<< HEAD
<<<<<<< HEAD

func FizzBuzz(i int) string {
	// TODO
	return ""
<<<<<<< HEAD
}
=======
>>>>>>> c57f921 (Preparation for the 1st homework from the second course)
=======

import "fmt"

func FizzBuzz(i int) string {

	switch {
	case i%15 == 0:

		return "FizzBuzz"

	case i%3 == 0:

		return "Fizz"

	case i%5 == 0:

		return "Buzz"

	default:

		fmt.Println(i)
	}
	return ""
=======
>>>>>>> ebed57b (Preparation for the 1st homework from the second course)
}
>>>>>>> 6f8c6c7 (Prep for the 1st homework from the second course)
