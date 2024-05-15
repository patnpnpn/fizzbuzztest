package kata

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1 when input 1", func(t *testing.T) {
		input := 1
		want := "1"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return 2 when input 2", func(t *testing.T) {
		input := 2
		want := "2"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return fizz when input 3", func(t *testing.T) {
		input := 3
		want := "fizz"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return 4 when input 4", func(t *testing.T) {
		input := 4
		want := "4"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})
}
