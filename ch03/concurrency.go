package ch03

func SimulateConcurrency(N int) map[int]bool {
	m := make(map[int]bool)
	x := 100
	go func(x *int) {
		*x = *x * *x
	}(&x)
	go func(x *int) {
		*x = *x + 1
	}(&x)
	m[x] = true
	return m
}
