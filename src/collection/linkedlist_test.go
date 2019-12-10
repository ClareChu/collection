package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	link := NewLinkedList()
	err := link.Add(1)
	assert.Equal(t, nil, err)
	item, err := link.GetFirst()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, item)
}

func TestLinkedList_GetFirst(t *testing.T) {
	link := NewLinkedList()
	item, err := link.GetFirst()
	assert.Equal(t, " link list is nil ", err.Error())
	assert.Equal(t, nil, item)
}

func TestLinkedList_AddIndex(t *testing.T) {
	link := NewLinkedList()
	err := link.AddIndex(2, 1)
	assert.Equal(t, " Constructs an IndexOutOfBoundsException", err.Error())


	err = link.Add(1)
	assert.Equal(t, nil, err)

	err = link.AddIndex(0, 2)
	assert.Equal(t, nil, err)
}

func TestLinkedList_Contains(t *testing.T) {
	link := NewLinkedList()
	b := link.Contains(1)
	assert.Equal(t, false, b)

	err := link.Add(1)
	assert.Equal(t, nil, err)

	b = link.Contains(1)
	assert.Equal(t, true, b)
}
