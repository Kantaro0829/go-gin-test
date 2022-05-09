package handler

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/json"
	"github.com/Kantaro0829/go-gin-test/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Getting(c *gin.Context) {
	//ハッシュ
	hash, err := bcrypt.GenerateFromPassword([]byte("plain text"), 12)
	if err != nil {
		panic("failed to hash password")
	}

	db := infra.DBInit()
	user := model.User{Mail: "test1@gmail.com", Password: hash, Age: 12}
	result := db.Create(&user)

	if result.Error != nil {
		panic("failed to insert the data to the users table")
	}
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func Posting(c *gin.Context) {
	var userJson json.JsonRequestUser

	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mail := userJson.UserMail
	age := userJson.UserAge
	password := userJson.UserPassword

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12) //ハッシュ

	if err != nil {
		panic("failed to hash password")
	}

	db := infra.DBInit()
	user := model.User{Mail: mail, Password: hashedPassword, Age: age}
	if result := db.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
		fmt.Println(result.RowsAffected)
		fmt.Println(reflect.TypeOf(result))
		fmt.Println(reflect.TypeOf(result.Error))

		panic("failed to insert the data to the users table")
	}

	// if result.Error != nil {

	// 	fmt.Println(result.Error)
	// 	fmt.Println(result.RowsAffected)
	// 	panic("failed to insert the data to the users table")

	// }
	c.JSON(http.StatusOK, gin.H{"message": "data was inserted"})
}
