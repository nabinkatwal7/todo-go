package controller

import (
	"net/http"
	"strconv"
	"todo-api/db"
	"todo-api/helper"
	"todo-api/model"

	"github.com/gin-gonic/gin"
)

func AddTodo(context *gin.Context){
	var input model.ToDo

	if err := context.ShouldBindJSON(&input); err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedTodo, err := input.Save()

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, savedTodo)
}

func GetAllEntries(context *gin.Context){
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"todos": user.Todos})
}

func UpdateTodo(context *gin.Context){
	id := context.Param("id")

	todoId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingTodo := model.ToDo{}
	existingTodo.UserID = user.ID
	existingTodo.ID = uint(todoId)

	if result := db.Database.First(&existingTodo); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	if err := context.ShouldBindJSON(&existingTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _,err := existingTodo.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, existingTodo)
}

func DeleteTodo(context *gin.Context){
	id := context.Param("id")

	todoId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingTodo := model.ToDo{}
	existingTodo.UserID = user.ID
	existingTodo.ID = uint(todoId)


	if result := db.Database.First(&existingTodo); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	if err := existingTodo.Delete(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
}