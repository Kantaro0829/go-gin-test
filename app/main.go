package main

import (
	"github.com/Kantaro0829/go-gin-test/handler"
	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/gin-gonic/gin"
)

func main() {

	infra.DBInit()
	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/get", handler.Getting)
		user.PUT("/reg", handler.UserReg)
		user.POST("/login", handler.UserLogin)
		user.PUT("/update", handler.UpdateUser)
		user.DELETE("/delete", handler.DeleteUser)
		user.POST("/validate", handler.SampleJwtValidation)

	}
	todo := router.Group("todo")
	{
		todo.PUT("/create", handler.CreateTodo)
		todo.PUT("/update", handler.UpdateTodo)

	}

	router.Run(":3000")

}
