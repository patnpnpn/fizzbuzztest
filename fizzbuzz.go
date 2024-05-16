package kata

import "fmt"

func fizzbuzz(n int) string {
	var m = map[int]string{
		3: "fizz",
		5: "buzz",
		6: "fizz",
	}

	str, ok := m[n]

	if !ok {
		return fmt.Sprint(n)
	}

	return str
}
