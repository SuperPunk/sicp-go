package ch01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, 6, GCD(123456, 7890))
	assert.Equal(t, 1, GCD(123456, 1))
	assert.Equal(t, 123456, GCD(123456, 123456))
}
