package src

const (
	/**
	 * Default initial capacity.
	 */

	DEFAULT_CAPACITY int = 10
)

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
		return true
	} else if a.size < index {
		return false
	} else {
		return false
	}
}

func (a *ArrayList) RemoveObject(index int) bool {

	return false
}
