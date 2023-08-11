package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DB struct {
	username string
	password string
	host     string
	port     int
	dbName   string
}

type BaseDBModel struct {
	dbConnect *gorm.DB
}

func (baseDBModel BaseDBModel) Connect() *gorm.DB {
	mydb := DB{
		username: "tiktok",
		password: "2aebEzEnWFEi44k2",
		host:     "47.113.178.29",
		port:     3300,
		dbName:   "demo",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mydb.username, mydb.password, mydb.host, mydb.port, mydb.dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭自动识别结构体复数
		},
	})
	if err != nil {
		panic(err)
	}
	return db
}
