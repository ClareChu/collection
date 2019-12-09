package src

import "errors"

type LinkedListInterface interface {
	LinkFirst(o interface{})
	LinkLast(o interface{})
	GetFirst() (item interface{}, err error)
	RemoveFirst() (o interface{}, err error)
	GetLast() (item interface{}, err error)
	linkBefore(element interface{}, node *Node)
	node(index int) (node *Node)
	isElementIndex(index int) bool
	checkElementIndex(index int) error
	Set(index int, element interface{}) (o interface{}, err error)
	unlink(x *Node) (element interface{})
	Remove(o interface{}) bool
	Size() (index int)
	indexOf(o interface{}) (index int)
	unlinkFirst() (o interface{})
	unlinkLast() (o interface{})
	Add(o interface{}) error
	Contains(o interface{}) bool
	RemoveLast() (o interface{})
	AddIndex(index int, element interface{}) (err error)
}

type Node struct {
	item interface{}
	next *Node
	prev *Node
}

type LinkedList struct {
	LinkedListInterface
	first *Node
	last  *Node
	size  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func NewNode(prev *Node, e interface{}, next *Node) *Node {
	return &Node{
		item: e,
		next: next,
		prev: prev,
	}
}

func (link *LinkedList) LinkFirst(o interface{}) {
	f := link.first
	node := NewNode(nil, o, f)
	if f == nil {
		link.last = node
	} else {
		f.prev = node
	}
	link.size++
}

func (link *LinkedList) LinkLast(o interface{}) {
	prev := link.last
	newNode := NewNode(prev, o, nil)
	link.last = newNode
	if prev == nil {
		link.first = newNode
	} else {
		link.last.next = newNode
	}
	link.size++
}

func (link *LinkedList) GetFirst() (item interface{}, err error) {
	node := link.first
	if node == nil {
		return nil, errors.New(" link list is nil ")
	}
	return node.item, nil
}

func (link *LinkedList) GetLast() (item interface{}, err error) {
	node := link.last
	if node == nil {
		return nil, errors.New(" link list is nil ")
	}
	return node.item, nil
}

func (link *LinkedList) RemoveFirst() (o interface{}, err error) {
	first := link.first
	if first == nil {
		return
	}
	return link.unlinkFirst(), nil
}

func (link *LinkedList) unlinkFirst() (o interface{}) {
	//第一个节点的data
	o = link.first.item
	// 第二个node
	n := link.first.next
	link.first = nil
	link.first = n
	if n == nil {
		link.last = nil
	} else {
		link.first.prev = nil
	}
	link.size--
	return o

}

func (link *LinkedList) RemoveLast() (o interface{}) {
	last := link.last
	if last == nil {
		return
	}
	return
}
func (link *LinkedList) unlinkLast() (o interface{}) {
	element := link.last.item
	prev := link.last.prev
	link.last = prev
	if prev == nil {
		link.first = nil
	} else {
		prev.next = nil
	}
	link.size--
	return element
}

func (link *LinkedList) Add(o interface{}) error {
	link.LinkLast(o)
	return nil
}

func (link *LinkedList) Contains(o interface{}) bool {
	if link.indexOf(o) != -1 {
		return true
	}
	return false
}

func (link *LinkedList) indexOf(o interface{}) (index int) {
	index = 0
	for x := link.first; x != nil; x = x.next {
		if x.item == o {
			return
		}
		index++
	}
	return -1
}

func (link *LinkedList) Size() (index int) {
	return link.size
}

func (link *LinkedList) Remove(o interface{}) bool {
	for x := link.first; x != nil; x = x.next {
		if x.item == o {
			link.unlink(x)
			return true
		}
	}
	return false
}

func (link *LinkedList) unlink(x *Node) (element interface{}) {
	prev := x.prev
	next := x.next
	element = x.item
	if prev == nil {
		link.first = next
	} else {
		prev.next = next
		x.prev = nil
	}
	if next == nil {
		link.last = prev
	} else {
		next.prev = prev
		x.next = nil
	}
	x.item = nil
	link.size--
	return element
}

/**
 * Replaces the element at the specified position in this list with the
 */
func (link *LinkedList) Set(index int, element interface{}) (o interface{}, err error) {
	err = link.checkElementIndex(index)
	if err != nil {
		return nil, err
	}
	node := link.node(index)
	oldVal := node.item
	node.item = element
	return oldVal, nil
}

func (link *LinkedList) checkElementIndex(index int) error {
	if !link.isElementIndex(index) {
		return errors.New(" Constructs an IndexOutOfBoundsException")
	}
	return nil
}

func (link *LinkedList) isElementIndex(index int) bool {
	if index >= 0 || index < link.size {
		return true
	}
	return false
}

func (link *LinkedList) node(index int) (node *Node) {
	if (link.size >> 1) > index {
		x := link.first
		for i := 0; i < index; i++ {
			x = x.next
		}
		return x
	} else {
		x := link.last
		for i := link.size - 1; i > index; i-- {
			x = x.prev
		}
		return x
	}
}

/**
 * Inserts the specified element at the specified position in this list.
 */
func (link *LinkedList) AddIndex(index int, element interface{}) (err error) {
	err = link.checkElementIndex(index)
	if err != nil {
		return err
	}
	if index == link.size {
		link.LinkLast(element)
	} else {
		link.linkBefore(element, link.node(index))
	}
	return nil
}
func (link *LinkedList)  linkBefore(element interface{}, node *Node) {
	prev := node.prev
	newNode := NewNode(prev, element, node)
	if prev == nil {
		link.first = newNode
	} else {
		prev.next = newNode
	}
	link.size++
}