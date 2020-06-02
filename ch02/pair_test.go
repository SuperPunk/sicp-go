package ch02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakePair(t *testing.T) {
	pair := MakePair(1, 2)
	assert.Equal(t, 1, pair("Left"))
	assert.Equal(t, 2, pair("right"))
}

func TestNewPair(t *testing.T) {
	pair := NewPair(1, 2)
	assert.Equal(t, 1, Left(pair))
	assert.Equal(t, 2, Right(pair))
}
