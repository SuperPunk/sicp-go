package ch01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrime(t *testing.T) {
	assert.Equal(t, true, IsPrime(1))
	assert.Equal(t, true, IsPrime(2))
	assert.Equal(t, true, IsPrime(3))
	assert.Equal(t, false, IsPrime(4))
	assert.Equal(t, true, IsPrime(5))
	assert.Equal(t, false, IsPrime(6))
	assert.Equal(t, true, IsPrime(7))
	assert.Equal(t, false, IsPrime(8))
	assert.Equal(t, false, IsPrime(9))
	assert.Equal(t, false, IsPrime(10))
	assert.Equal(t, true, IsPrime(11))
	assert.Equal(t, false, IsPrime(12))
	assert.Equal(t, true, IsPrime(13))
	assert.Equal(t, false, IsPrime(14))
	assert.Equal(t, false, IsPrime(15))
	assert.Equal(t, false, IsPrime(16))
	assert.Equal(t, true, IsPrime(17))
	assert.Equal(t, false, IsPrime(18))
	assert.Equal(t, true, IsPrime(19))
	assert.Equal(t, false, IsPrime(20))
}
