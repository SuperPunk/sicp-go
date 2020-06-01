package ch02

type BinaryTreeNode struct {
	elem  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func ElementOfBinarySearchTree(x int, root *BinaryTreeNode) bool {
	if root == nil {
		return false
	}
	if root.elem == x {
		return true
	}
	if root.elem > x {
		return ElementOfBinarySearchTree(x, root.left)
	}
	return ElementOfBinarySearchTree(x, root.right)
}

func AdjoinSetForBinaryTree(x int, root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return MakeTree(x, nil, nil)
	}
	if x == root.elem {
		return root
	}
	if x < root.elem {
		return MakeTree(root.elem, AdjoinSetForBinaryTree(x, root.left), root.right)
	}
	return MakeTree(root.elem, root.left, AdjoinSetForBinaryTree(x, root.right))
}

func MakeTree(entry int, left *BinaryTreeNode, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{entry, left, right}
}
