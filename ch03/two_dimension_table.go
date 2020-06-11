package ch03

type TableBackbone struct {
	subTableHeaderPtr *SubTableHeader
	next              *TableBackbone
}

type SubTableHeader struct {
	subTableTitle    string
	subTableBackbone *SubTableBackboneNode
}

// 子表骨架
type SubTableBackboneNode struct {
	dataPtr *SubTableItemNode
	next    *SubTableBackboneNode
}

// 数据项key-value pair
type SubTableItemNode struct {
	key   string
	value string
}
