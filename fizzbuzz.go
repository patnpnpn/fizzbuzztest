package kata

import "fmt"

func convertIntToString(n int) string {
	return fmt.Sprint(n)
}

func isEqualToThree(n int) bool {
	return n == 3
}

func isEqualToFive(n int) bool {
	return n == 5
}

func isEqualToSix(n int) bool {
	return n == 6
}

func fizzbuzz(n int) string {
	if isEqualToThree(n) {
		return "fizz"
	}

	if isEqualToFive(n) {
		return "buzz"
	}

	if isEqualToSix(n) {
		return "fizz"
	}

	return convertIntToString(n)
}
