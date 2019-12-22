//go:generate wire

//+build wireinject

package db

import "github.com/google/wire"

func InitializeMysqlConnect() (*Data, error) {
	wire.Build(ConnectSet)
	return &Data{}, nil
}
