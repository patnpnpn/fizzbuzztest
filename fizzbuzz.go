package kata

import "fmt"

func convertIntToString(n int) string {
	return fmt.Sprint(n)
}

func isDivisibleByThree(n int) bool {
	return n%3 == 0
}

func isEqualToFive(n int) bool {
	return n == 5
}

func fizzbuzz(n int) string {
	if isDivisibleByThree(n) {
		return "fizz"
	}

	if isEqualToFive(n) {
		return "buzz"
	}

	return convertIntToString(n)
}
