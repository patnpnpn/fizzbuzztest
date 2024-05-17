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

func fizzbuzz(n int) string {
	if isEqualToThree(n) {
		return "fizz"
	}

	if isEqualToFive(n) {
		return "buzz"
	}

	return convertIntToString(n)
}
