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
}