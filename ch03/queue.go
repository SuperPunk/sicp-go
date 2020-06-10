package ch03

type Queue struct {
	frontPtr *Pair
	rearPtr  *Pair
}

type Pair struct {
	element int
	next    *Pair
}

func MakeQueue() *Queue {
	return &Queue{nil, nil}
}

func InsertQueue(q *Queue, item int) *Queue {
	return func(q *Queue, newPair *Pair) *Queue {
		if emptyQueue(q) {
			q.frontPtr, q.rearPtr = newPair, newPair
		} else {
			q.rearPtr.next, q.rearPtr = newPair, newPair
		}
		return q
	}(q, &Pair{item, nil})
}

func DeleteQueue(q *Queue) *Queue {
	if emptyQueue(q) {
		panic("queue is empty")
	}
	q.frontPtr = q.frontPtr.next
	return q
}

func FrontQueue(q *Queue) int {
	return q.frontPtr.element
}

func RearQueue(q *Queue) int {
	return q.rearPtr.element
}

func emptyQueue(q *Queue) bool {
	return q.frontPtr == nil
}
