package main

import (
	"fmt"
	"restapi/business/domain"
	"restapi/business/entity"
	"restapi/business/usecase"
	"restapi/handler/rest"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := Init()

	d := domain.Init(db)

	uc := usecase.Init(d)

	r := rest.Init(uc)

	r.Run()
}

func Init() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "password", "localhost", "3306", "benchmark")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		panic(err)
	}

	return db
}
