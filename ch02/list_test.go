package ch02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListRef(t *testing.T) {
	list := NewList(1, 2, 3, 4, 5)
	elem := ListRef(list, 3)

	if elem != 4 {
		t.Errorf("elem = %v, want 1", elem)
	}
}

func TestEnumerateIntervalN(t *testing.T) {
	list := EnumerateIntervalN(5)
	assert.Equal(t, 1, ListRef(list, 0))
	assert.Equal(t, 2, ListRef(list, 1))
	assert.Equal(t, 3, ListRef(list, 2))
	assert.Equal(t, 4, ListRef(list, 3))
	assert.Equal(t, 5, ListRef(list, 4))
}
