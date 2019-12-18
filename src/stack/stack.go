package stack

import "github.com/ClareChu/collection/src/collection"

type StackInterface interface {
	Push(es ...interface{}) //向栈顶添加元素
	Pop() interface{}   //移除栈顶元素
	Top() interface{}   //获取栈顶元素（不删除）
	Clear() bool        //清空栈
	Size() int          //获取栈的元素个数
	IsEmpty() bool      //判断栈是否是空栈
}

type Stack struct {
	StackInterface
	list *collection.LinkedList
}

func NewStack() *Stack {
	list := &collection.LinkedList{}
	return &Stack{
		list:           list,
	}
}

func (s *Stack) Push(es ...interface{}) {
	for _, e := range es  {
		s.list.Add(e)
	}
}

func (s *Stack) Pop() interface{} {
	o, err := s.list.GetLast()
	if err != nil {
		return nil
	}
	o = s.list.RemoveLast()
	return o
}

func (s *Stack) Top() interface{} {
	o, err := s.list.GetLast()
	if err != nil {
		return nil
	}
	return o
}


func (s *Stack) Clear() bool {
	s.list.Clear()
	return true
}

func (s *Stack) Size() (index int) {
	return s.list.Size()
}

func (s *Stack) IsEmpty() bool {
	return s.list.IsElementIndex(0)
}