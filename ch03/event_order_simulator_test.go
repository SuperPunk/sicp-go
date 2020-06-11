package ch03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombine(t *testing.T) {
	ret := Interleave([]string{"a", "b", "c"}, []string{"x", "y", "z"})

	assert.Equal(t, 20, len(ret))
	assert.ElementsMatch(t, ret, [][]string{
		{"a", "b", "c", "x", "y", "z"},
		{"a", "x", "b", "y", "c", "z"},
		{"x", "a", "b", "c", "y", "z"},
		{"x", "a", "y", "z", "b", "c"},
		{"a", "b", "x", "c", "y", "z"},
		{"a", "x", "b", "y", "z", "c"},
		{"x", "a", "b", "y", "c", "z"},
		{"x", "y", "a", "b", "c", "z"},
		{"a", "b", "x", "y", "c", "z"},
		{"a", "x", "y", "b", "c", "z"},
		{"x", "a", "b", "y", "z", "c"},
		{"x", "y", "a", "b", "z", "c"},
		{"a", "b", "x", "y", "z", "c"},
		{"a", "x", "y", "b", "z", "c"},
		{"x", "a", "y", "b", "c", "z"},
		{"x", "y", "a", "z", "b", "c"},
		{"a", "x", "b", "c", "y", "z"},
		{"a", "x", "y", "z", "b", "c"},
		{"x", "a", "y", "b", "z", "c"},
		{"x", "y", "z", "a", "b", "c"},
	})
}
