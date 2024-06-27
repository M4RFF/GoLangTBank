package fizzbuzz

func Fizzbuzz(i int) (string, int) {

	switch {
	case i%15 == 0:
		return "FizzBuzz", i

	case i%5 == 0:
		return "Buzz", i

	case i%3 == 0:
		return "Fizz", i

	}
	return "", i
}
