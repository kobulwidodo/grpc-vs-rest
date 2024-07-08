package main

import (
	"fmt"
	"grpc/business/domain"
	"grpc/business/entity"
	"grpc/business/usecase"
	"grpc/handler/grpc"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := Init()

	d := domain.Init(db)

	uc := usecase.Init(d)

	g := grpc.Init(uc)

	g.Run()
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
