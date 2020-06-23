package ch03

import (
	"github.com/stretchr/testify/assert"
	"sicp/ch01"
	"testing"
)

func TestStreamEnumerateInterval(t *testing.T) {
	s := StreamEnumerateInterval(1, 10)
	assert.Equal(t, 3, s.StreamRef(2))
}

func TestStream_StreamFilter(t *testing.T) {
	s := StreamEnumerateInterval(10000, 1000000).StreamFilter(ch01.IsPrime)
	assert.Equal(t, 10009, s.next().value)
}
