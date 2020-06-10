package ch03

func MakeSimplifiedWithdraw(balance int) func(amount int) int {
	return func(amount int) int {
		balance -= amount
		return balance
	}
}
