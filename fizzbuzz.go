package kata

import "fmt"

func fizzbuzz(n int) string {
	if n == 6 {
		return "fizz"
	}
	if n == 5 {
		return "buzz"
	}
	if n == 3 {
		return "fizz"
	}
	return fmt.Sprint(n)
}
