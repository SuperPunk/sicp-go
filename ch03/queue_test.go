package ch03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := MakeQueue()
	InsertQueue(queue, 1)
	InsertQueue(queue, 2)
	InsertQueue(queue, 3)

	assert.Equal(t, FrontQueue(queue), 1)
	assert.Equal(t, RearQueue(queue), 3)

	DeleteQueue(queue)

	assert.Equal(t, FrontQueue(queue), 2)
	assert.Equal(t, RearQueue(queue), 3)
}
