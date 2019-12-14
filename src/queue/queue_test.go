package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	index := queue.GetSize()
	assert.Equal(t, index, 0)
}

func TestQueue_Put(t *testing.T) {
	queue := NewQueue()
	err := queue.Put(1)
	assert.Equal(t, nil, err)

	//查看当前队列的size
	index := queue.GetSize()
	assert.Equal(t, index, 1)

	o, err := queue.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, o, 1)

	//查看当前队列的size
	index = queue.GetSize()
	assert.Equal(t, index, 0)
}

func TestQueue_Pop(t *testing.T) {
	queue := NewQueue()
	o, err := queue.Pop()
	assert.Equal(t, o, nil)
	assert.Equal(t, " link list is nil ", err.Error())

	err = queue.Put(1)
	assert.Equal(t, nil, err)

	err = queue.Put(2)
	assert.Equal(t, nil, err)

	o, err = queue.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, o, 1)
}