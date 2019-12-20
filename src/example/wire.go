//+build wireinject

package main

import "github.com/google/wire"

func Initialize() *controller.Iris {
	wire.Build(controller.NewIris)
	return &controller.Iris{}
}
