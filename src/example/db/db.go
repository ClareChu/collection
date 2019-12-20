package db

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataSource struct {
	Host      string `json:"host" yaml:"host"`
	Port      int    `json:"port" yaml:"port"`
	Username  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	DataBase  string `json:"dataBase" yaml:"data_base"`
	Charset   string `json:"charset" yaml:"charset"` //想要完全的支持 UTF-8 编码，你需要修改charset=utf8 为 charset=utf8mb4
	ParseTime bool   `json:"parseTime" yaml:"parse_time"`
	Loc       string `json:"loc" yaml:"loc"`
}

var MysqlConnectSet = wire.NewSet(NewMysqlConnect, DataSourceSet)

//NewMysqlConnect 新建mysql 的连接池
func NewMysqlConnect(dataSource *DataSource) (db *gorm.DB, err error) {
	// 连接数据库的url user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	// user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	mysql := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v", dataSource.Username,
		dataSource.Password,
		dataSource.Host,
		dataSource.Port,
		dataSource.DataBase,
		dataSource.Charset,
		dataSource.ParseTime,
		dataSource.Loc)
	db, err = gorm.Open("mysql", mysql)
	if err != nil {
		// todo logging
		return
	}
	return
	//defer db.Close()

}

var DataSourceSet = wire.NewSet(NewDataSource)

// 初始化db 数据

func NewDataSource() *DataSource {
	return &DataSource{}
}
