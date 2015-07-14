package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/thedarksavant/gorm/models"
)

func main() {
	db, err := gorm.Open("mysql", "test:goober@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	db.DB()
	db.AutoMigrate(&User{})
}
