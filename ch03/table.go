package ch03

// 一维表(key-value)实现方式：
//    1) 维护一个"骨架"list，itemPtr指向数据项，backbonePtr指向后继骨架节点
type BackboneNode struct {
	itemPtr     *ItemNode
	backbonePtr *BackboneNode
}

type ItemNode struct {
	key   string
	value int
}
