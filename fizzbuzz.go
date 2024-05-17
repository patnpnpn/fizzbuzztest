package kata

import "fmt"

type Fizzbuzzer interface {
	Fizzbuzzing() string
}

type Fizzbuzz struct {
	n int
}

func NewFizzBuzz(n int) Fizzbuzzer {
	return Fizzbuzz{n: n}
}

func (fb Fizzbuzz) Fizzbuzzing() string {
	if fb.n == 3 {
		return "fizz"
	}

	return fmt.Sprint(fb.n)
}

func fizzbuzz(n int) string {
	fb := NewFizzBuzz(n)

	return fb.Fizzbuzzing()
}
