package db

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Data struct {
	DataSource *DataSource `json:"datasource" yaml:"datasource"`
}

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

type Connect struct {
	DB *gorm.DB `json:"db"`
}

var ConnectSet = wire.NewSet(NewData)

//NewMysqlConnect 新建mysql 的连接池
func (data *Data) Connect() (connect *Connect, err error) {
	// 连接数据库的url user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	// user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	mysql := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		data.DataSource.Username,
		data.DataSource.Password,
		data.DataSource.Host,
		data.DataSource.Port,
		data.DataSource.DataBase,
		data.DataSource.Charset,
		data.DataSource.ParseTime,
		data.DataSource.Loc)
	db, err := gorm.Open("mysql", mysql)
	if err != nil {
		// todo logging
		return
	}
	return &Connect{DB: db}, nil
	//defer db.Close()

}

// 初始化db 数据

func NewData() *Data {
	return &Data{}
}
