package foobarbaz

import "github.com/google/wire"

func initializeBaz() *Baz {
	wire.Build(NewBar, NewBaz, NewFoo)
	return &Baz{}
}