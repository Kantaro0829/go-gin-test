package handler

import (
	"fmt"
	"net/http"

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
	var age uint8
	age = 12
	mail := "test1@gmail.com"
	db := infra.DBInit()
	user := model.User{Mail: mail, Password: hash, Age: age}
	result := db.Create(&user)

	if result.Error != nil {
		panic("failed to insert the data to the users table")
	}
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func UserReg(c *gin.Context) {
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
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}
	fmt.Println("登録されたパスワード")
	fmt.Println(user.Password)

	c.JSON(http.StatusOK, gin.H{"message": "data was inserted"})
}

func UserLogin(c *gin.Context) {
	var userLoginJson json.JsonRequestUser

	if err := c.ShouldBindJSON(&userLoginJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mail := userLoginJson.UserMail
	password := userLoginJson.UserPassword
	fmt.Println(mail)
	fmt.Println(password)
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12) //ハッシュ

	// if err != nil {
	// 	panic("failed to hash password")
	// }

	db := infra.DBInit()
	user := model.User{}

	// Get first matched record
	result := db.Select("password").Where("mail = ?", mail).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	if isAuthorized := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); isAuthorized != nil {

		fmt.Println("不一致")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "failed to login"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "succeed login"})

}
