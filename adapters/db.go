package adapters

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Store *gorm.DB
}

func InitDB(dsn string) DB {

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("can not connect to db")
	}

	var entites = []any{}

	if migrationErr := conn.AutoMigrate(entites...); migrationErr != nil {
		panic("auto migarion fail")
	}

	return DB{Store: conn}
}
