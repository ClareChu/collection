package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	list := NewArrayList()
	a := list.Add(1)
	println(a)
}

func TestNewArrayList(t *testing.T) {
	a := NewArrayList()
	assert.Equal(t, a.size, 0)
}
