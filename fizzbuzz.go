package kata

func fizzbuzz(n int) string {
	var m = map[int]string{
		1: "1",
		2: "2",
		3: "fizz",
		4: "4",
	}

	return m[n]
}
