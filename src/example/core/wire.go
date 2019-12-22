//+build wireinject
//go:generate wire


package core

import "github.com/google/wire"

func InitializeConfig(fileName string, i interface{}) ([]byte, error) {
	wire.Build(ConfigSet)
	return []byte{}, nil
}
