package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "_mysql_connection_string")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
