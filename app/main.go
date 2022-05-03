package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Mail     string
	Password []byte
	Age      uint8
}

//
func main() {
	dsn := "root:ecc@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	//ハッシュ
	hash, err := bcrypt.GenerateFromPassword([]byte("plain text"), 12)
	if err != nil {
		panic("failed to hash password")
	}
	fmt.Println(hash)

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
	    user := User{Mail: "test@gmail.com", Password: hash, Age: 12}
	    result := db.Create(&user)
        if result.Error != nil {
            panic("failed to insert the data to the users table")
            fmt.Println(result.Error)
        }
        fmt.Println(user.ID)
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":3000")
}
