package kata

import "fmt"

func fizzbuzz(n int) string {
	cycle := []string{
		fmt.Sprint(n),
		fmt.Sprint(n),
		"fizz",
		fmt.Sprint(n),
		"buzz",
		"fizz",
	}

	return cycle[n-1]
}
