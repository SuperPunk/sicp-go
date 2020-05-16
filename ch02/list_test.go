package ch02

import "testing"

func TestListRef(t *testing.T) {
	list := NewList(1, 2, 3, 4, 5)
	elem := ListRef(list, 3)

	if elem != 4 {
		t.Errorf("elem = %v, want 1", elem)
	}
}
