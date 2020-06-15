package ch03

import "sync"

// 运行一百万次，会复现出四种结果：101，121，11，100。110不会出现
func SimulateConcurrency(N int) map[int]bool {
	m := make(map[int]bool)
	var wg sync.WaitGroup
	for i := 1; i <= N; i++ {
		x := 10
		wg.Add(2)
		go func(x *int) {
			*x = *x * *x
			wg.Done()
		}(&x)

		go func(x *int) {
			*x = *x + 1
			wg.Done()
		}(&x)

		wg.Wait()
		m[x] = true
	}
	return m
}
