package foobarbaz

// ...

type Bar struct {
	Foo *Foo
	X   int
}

func NewBar(foo *Foo) *Bar {
	return &Bar{Foo: foo}
}
