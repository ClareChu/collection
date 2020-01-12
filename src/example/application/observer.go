package application

type Subject struct {
	Observers map[string]func() interface{}
}

var sub *Subject

func NewSubject() *Subject {
	if sub == nil {
		return &Subject{Observers: map[string]func() interface{}{}}
	}
	return sub
}

func goFunc1(f func() interface{}) {
	f()
}

func Notify() (o interface{}) {
	for _, v := range sub.Observers {
		goFunc1(v)
	}
	return
}

func Register(model string, o func() interface{}) {
	sub := NewSubject()
	sub.Observers[model] = o
}
