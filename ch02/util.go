package ch02

type Method func(*Node, *Node) *Node

type Op func(interface{}, interface{}) interface{}

type Proc func(interface{}) interface{}

func Accumulate(op Op, initial interface{}, root *Node) interface{} {
	if root == nil {
		return initial
	}
	return op(root.elem, Accumulate(op, initial, root.next))
}
