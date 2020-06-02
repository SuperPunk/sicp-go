package ch02

type Selector func(int, int) int

// 将pair数据表示成实体(过程)
type Pair func(Selector) int

type PairDispatcher func(op string) int

// 将pair数据表示成实体(过程)
func NewPair(x, y int) Pair {
	return func(s Selector) int { return s(x, y) }
}

// 以消息的方式传递操作的名字,此处是Left
func Left(p Pair) int {
	return p(func(x, y int) int { return x })
}

// 以消息的方式传递操作的名字,此处是Left
func Right(p Pair) int {
	return p(func(x, y int) int { return y })
}

func MakePair(x, y int) PairDispatcher {
	return func(op string) int {
		if op == "Left" {
			return x
		} else {
			return y
		}
	}
}
