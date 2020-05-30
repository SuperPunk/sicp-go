package ch02

func ElementOfSet(x interface{}, set *Node) bool {
	if set == nil {
		return false
	}
	if x == set.elem {
		return true
	}
	return ElementOfSet(x, set.next)
}

func ElementOfOrderedSet(x int, set *Node) bool {
	if set == nil {
		return false
	}
	if x == set.elem {
		return true
	}
	if x < set.elem.(int) {
		return false
	}
	return ElementOfOrderedSet(x, set.next)
}

func AdjoinSet(x interface{}, set *Node) *Node {
	if ElementOfSet(x, set) {
		return set
	}
	return &Node{x, set}
}

func IntersectionSet(set1, set2 *Node) *Node {
	if set1 == nil || set2 == nil {
		return nil
	}
	if ElementOfSet(set1.elem, set2) {
		return &Node{set1.elem, IntersectionSet(set1.next, set2)}
	}
	return IntersectionSet(set1.next, set2)
}
