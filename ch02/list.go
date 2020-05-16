package ch02

type Node struct {
	elem int
	next *Node
}

func NewList(elements ...int) *Node {
	if len(elements) == 0 {
		return nil
	}
	return &Node{elements[0], NewList(elements[1:]...)}
}

func ListRef(node *Node, n int) int {
	if n == 0 {
		return node.elem
	}
	return ListRef(node.next, n-1)
}

func Length(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + Length(node.next)
}

func Append(l1, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	return &Node{l1.elem, Append(l1.next, l2)}
}

func LengthByIter(node *Node) int {
	return lengthIterHelper(node, 0)
}

func lengthIterHelper(node *Node, count int) int {
	if node == nil {
		return count
	}
	return lengthIterHelper(node.next, count+1)
}
