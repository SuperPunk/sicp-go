package ch02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkAdjoinSetForBinaryTree(b *testing.B) {
	// construct BST
	var root *BinaryTreeNode
	for i := 0; i < b.N; i++ {
		root = AdjoinSetForBinaryTree(i, root)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ElementOfBinarySearchTree(b.N, root)
	}
}

func TestElementOfBinarySearchTree(t *testing.T) {
	var root *BinaryTreeNode
	for i := 0; i < 10000; i++ {
		root = AdjoinSetForBinaryTree(i, root)
	}

	assert.Equal(t, false, ElementOfBinarySearchTree(10000, root))
}
