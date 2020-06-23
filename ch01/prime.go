package ch01

// 最小的除数是自己
func IsPrime(n int) bool {
	return n == smallestDivisor(n)
}

func smallestDivisor(n int) int {
	return findDivisor(n, 2)
}

func findDivisor(n int, testDivisor int) int {
	if testDivisor*testDivisor > n {
		return n
	}
	if n%testDivisor == 0 {
		return testDivisor
	}
	return findDivisor(n, testDivisor+1)
}
