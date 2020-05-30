package ch02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElementOfSet(t *testing.T) {
	set := EnumerateInterval(1, 5)
	assert.Equal(t, true, ElementOfSet(1, set))
	assert.Equal(t, false, ElementOfSet(0, set))
}

func TestIntersectionSet(t *testing.T) {
	set1 := EnumerateInterval(1, 5)
	set2 := EnumerateInterval(3, 10)
	set := IntersectionSet(set1, set2)
	assert.Equal(t, 3, set.elem)
	assert.Equal(t, 4, set.next.elem)
	assert.Equal(t, 5, set.next.next.elem)
}

func TestElementOfOrderedSet(t *testing.T) {
	set := EnumerateInterval(1, 10)
	assert.Equal(t, true, ElementOfOrderedSet(1, set))
	assert.Equal(t, true, ElementOfOrderedSet(10, set))
	assert.Equal(t, false, ElementOfOrderedSet(0, set))
	assert.Equal(t, false, ElementOfOrderedSet(11, set))
}
