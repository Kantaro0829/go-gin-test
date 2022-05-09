package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
 	//"gorm.io/gorm"
//     "github.com/Kantaro0829/go-gin-test/infra"
//     "github.com/Kantaro0829/go-gin-test/model"
    "github.com/Kantaro0829/go-gin-test/handler"
)

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }
//
// type User struct {
// 	gorm.Model
// 	Mail     string
// 	Password []byte
// 	Age      uint8
// }

//
func main() {
// 	dsn := "root:ecc@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
	// Migrate the schema
// 	d := infra.DBInit()
// 	d.AutoMigrate(&model.Product{})
// 	d.AutoMigrate(&model.User{})

	//Create
	//d.Create(&model.Product{Code: "MODEL TEST", Price: 100})
// 	user := model.User{Mail: "test@gmail.com", Password: hash, Age: 12}
//  result := d.Create(&user)
//  if result.Error != nil {
//      panic("failed to insert the data to the users table")
//      fmt.Println(result.Error)
//   }
//   fmt.Println(user.ID)


	router := gin.Default()
	router.GET("/get", handler.Getting)
	router.POST("post", handler.Posting)
	router.Run(":3000")
}
