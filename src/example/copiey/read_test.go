package example

import (
	"testing"
)

type Foo struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Bar  Bar
}

type Bar struct {
	Ha string `json:"ha"`
}

func TestIsZero(t *testing.T) {
	foo := &Foo{
		Name: "sasasasas",
		Id:   2,
		Bar: Bar{
			Ha: "sasasa",
		},
	}
	fo := &Foo{}
	Copy(foo, fo)
}
