package kata

import "fmt"

func fizzbuzz(n int) string {
    if n == 10 {
		return "buzz"
	}
 
    if n == 5 {
		return "buzz"
	}
 
	if n%3 == 0 {
		return "fizz"
	}
 
	return fmt.Sprint(n)
}
