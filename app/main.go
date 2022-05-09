package main

import (
	"github.com/Kantaro0829/go-gin-test/handler"
	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/gin-gonic/gin"
)

func main() {

	infra.DBInit()
	router := gin.Default()
	router.GET("/get", handler.Getting)
	router.POST("/post", handler.Posting)
	router.Run(":3000")

}
