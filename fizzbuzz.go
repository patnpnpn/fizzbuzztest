package kata

import "fmt"

func fizzbuzz(n int) string {
	var values []string

	for i := 1; i <= n; i++ {
		values = append(values, fmt.Sprint(i))
	}

	for i := 3 - 1; i < n; i += 3 {
		values[i] = "fizz"
	}

	for i := 5 - 1; i < n; i += 5 {
		values[i] = "buzz"
	}

	return values[n-1]
}
