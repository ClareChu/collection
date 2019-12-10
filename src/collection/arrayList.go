package collection

import "errors"

type ArrayListInterface interface {
	Collection
}

type ArrayList struct {
	ArrayListInterface
	ElementData []interface{}
	size        int
}

func NewArrayList() (arrayList *ArrayList) {
	return &ArrayList{ElementData: nil,
		size: 0,
	}
}

func (a *ArrayList) Add(o interface{}) bool {
	a.ElementData = append(a.ElementData, o)
	a.size++
	return true
}

func (a *ArrayList) Remove(index int) bool {
	if a.size >= index {
		a.ElementData = append(a.ElementData[:index], a.ElementData[index:]...)
		a.size--
		return true
	} else if a.size < index {
		return false
	} else {
		return false
	}
}

func (a *ArrayList) RemoveObject(o interface{}) (err error) {
	if o == nil {
		return errors.New("remove object is nil")
	} else {
		for i := 0; i < a.size; i++ {
			if o == a.ElementData[i] {
				a.ElementData = append(a.ElementData[:i], a.ElementData[i:]...)
				a.size--
				i--
			}
		}
	}

	return errors.New("run time is error")
}

func (a *ArrayList) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) Contains(value interface{}) bool {
	for _, curValue := range a.ElementData {
		if curValue == value {
			return true
		}
	}

	return false
}

func (a *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		return nil
	}
	return a.ElementData[index]
}
