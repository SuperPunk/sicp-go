package ch02

type Proc func(int) int

type Method func(*Node, *Node) *Node

type Accumulate func(Method, interface{}, *Node)
