package kata

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
	return "1"
}

func fizzbuzz(n int) string {
	fb := NewFizzBuzz(n)

	return fb.Fizzbuzzing()
}
