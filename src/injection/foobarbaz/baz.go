package foobarbaz

// ...

type Baz struct {
	Bar *Bar
}

func NewBaz(bar *Bar) *Baz {
	return &Baz{Bar: bar}
}
