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

	t.Run("should return buzz when input 5", func(t *testing.T) {
		input := 5
		want := "buzz"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return fizz when input 6", func(t *testing.T) {
		input := 6
		want := "fizz"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return 7 when input 7", func(t *testing.T) {
		input := 7
		want := "7"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return 8 when input 8", func(t *testing.T) {
		input := 8
		want := "8"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return fizz when input 9", func(t *testing.T) {
		input := 9
		want := "fizz"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return buzz when input 10", func(t *testing.T) {
		input := 10
		want := "buzz"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return 11 when input 11", func(t *testing.T) {
		input := 11
		want := "11"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return fizz when input 12", func(t *testing.T) {
		input := 12
		want := "fizz"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return 13 when input 13", func(t *testing.T) {
		input := 13
		want := "13"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return 14 when input 14", func(t *testing.T) {
		input := 14
		want := "14"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})

	t.Run("should return fizzbuzz when input 15", func(t *testing.T) {
		input := 15
		want := "fizzbuzz"

		get := fizzbuzz(input)

		if get != want {
			t.Errorf("should return %s when input %s", want, get)
		}
	})
}
