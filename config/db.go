package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	var db = GetConfigEnv().Database
	connStr := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=%v&loc=%v&timeout=%v",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.DBName,
		db.ParseTime,
		db.Loc,
		db.Timeout,
	)
	fmt.Println(connStr)
	DB, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	fmt.Println("数据库连接成功")
}
