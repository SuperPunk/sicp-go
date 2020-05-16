package ch02

type Selector func(int, int) int

type Pair func(Selector) int

func NewPair(x, y int) Pair {
	return func(s Selector) int { return s(x, y) }
}

func Left(p Pair) int {
	return p(func(x, y int) int { return x })
}

func Right(p Pair) int {
	return p(func(x, y int) int { return y })
}
