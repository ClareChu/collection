//go:generate wire
//+build wireinject

package db

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)


func InitializeMysqlConnect() (db *gorm.DB, err error) {
	wire.Build(MysqlConnectSet)
	db = &gorm.DB{}
	return db, nil
}

func DataSourceInitialize() (dataSource *DataSource) {
	wire.Build(DataSourceSet)
	return &DataSource{}
}
