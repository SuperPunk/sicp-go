package ch02

type TreeNode struct {
	data     int
	children []*TreeNode
}

func CountLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.children == nil {
		return 1
	}
	count := 0
	for _, child := range root.children {
		count += CountLeaves(child)
	}
	return count
}

func Map(proc Proc, root *TreeNode) {
	if root == nil {
		return
	}
	if root.children == nil {
		root.data = proc(root.data)
	}
	for _, child := range root.children {
		Map(proc, child)
	}
}
