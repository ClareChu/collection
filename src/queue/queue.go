package queue

import "github.com/ClareChu/collection/src/collection"

/*
队列的特性较为单一，基本操作即初始化、获取大小、添加元素、移除元素等。
最重要的特性就是满足先进先出
*/
type Queue struct {
	list *collection.LinkedList
}

func NewQueue() *Queue {
	list := collection.NewLinkedList()
	return &Queue{list: list}
}

//加入队列
func (queue *Queue) Put(data interface{}) (err error) {
	err = queue.list.Add(data)
	return err
}

//pop出队列
func (queue *Queue) Pop() (o interface{}, err error) {
	o, err = queue.list.GetFirst()
	if err != nil {
		return
	}
	o, err = queue.list.RemoveFirst()
	return
}


//获得队列的长度
func (queue *Queue) GetSize() (index int) {
	return queue.list.Size()
}