package ch02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountLeaves(t *testing.T) {
	n1 := &TreeNode{1, nil}
	n2 := &TreeNode{2, nil}
	n3 := &TreeNode{3, nil}
	n4 := &TreeNode{4, nil}
	n5 := &TreeNode{0, []*TreeNode{n1, n2}}
	root := &TreeNode{0, []*TreeNode{n5, n3, n4}}

	count := CountLeaves(root)
	if count != 4 {
		t.Errorf("count = %v, want 4", count)
	}
}

func TestMap(t *testing.T) {
	n1 := &TreeNode{1, nil}
	n2 := &TreeNode{2, nil}
	n3 := &TreeNode{3, nil}
	n4 := &TreeNode{4, nil}
	n5 := &TreeNode{0, []*TreeNode{n1, n2}}
	root := &TreeNode{0, []*TreeNode{n5, n3, n4}}

	Map(func(i interface{}) interface{} {
		intI, _ := i.(int)
		return intI * 2
	}, root)

	assert.Equal(t, 2, n1.data, "The two element should be same")
	assert.Equal(t, 4, n2.data, "The two element should be same")
	assert.Equal(t, 6, n3.data, "The two element should be same")
	assert.Equal(t, 8, n4.data, "The two element should be same")
}
