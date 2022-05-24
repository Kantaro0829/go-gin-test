package handler

import (
	"fmt"
	"net/http"

	//"github.com/Kantaro0829/go-gin-test/auth"
	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/json"
	"github.com/Kantaro0829/go-gin-test/model"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var taskJson json.CreateTaskJson

	if err := c.ShouldBindJSON(&taskJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	title := taskJson.Title
	isDone := taskJson.IsDone

	db := infra.DBInit()
	task := model.Task{Title: title, IsDone: isDone}

	if result := db.Create(&task); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data was inserted"})
}

func UpdateTodo(c *gin.Context) {
	var updateTaskJson json.UpdateTaskJson

	if err := c.ShouldBindJSON(&updateTaskJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := updateTaskJson.Id
	title := updateTaskJson.Title
	isDone := updateTaskJson.IsDone

	db := infra.DBInit()
	task := model.Task{}
	//task.ID = id

	if result := db.Model(&task).Where("id = ?", id).Updates(model.Task{Title: title, IsDone: isDone}); result.Error != nil {
		fmt.Println("DBの更新ができていません")
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": 503})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 200})

}
