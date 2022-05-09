package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/Kantaro0829/go-gin-test/model"
)

func DBInit () *gorm.DB{
    dsn := "root:ecc@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
    	panic("failed to connect database")
    }

    // Migrate the schema
   	//d := infra.DBInit()
   	db.AutoMigrate(&model.Product{})
    db.AutoMigrate(&model.User{})

    return db
}