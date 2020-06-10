package ch03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSimplifiedWithdraw(t *testing.T) {
	W := MakeSimplifiedWithdraw(100)
	assert.Equal(t, 80, W(20))
	assert.Equal(t, 60, W(20))
}
