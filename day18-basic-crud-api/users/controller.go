package users

import (
	"go-basics/day18-basic-crud-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	result, err := getAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	utils.SendResponse(ctx, utils.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Users fetched successfully",
		Data:       result,
	})
}

func GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid id"})
		return
	}

	result, err := getUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	utils.SendResponse(ctx, utils.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "User fetched successfully",
		Data:       result,
	})
}

func CreateUser(ctx *gin.Context) {
	var body User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	result, err := createUser(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	utils.SendResponse(ctx, utils.Response{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "User created successfully",
		Data:       result,
	})
}

func UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid id"})
		return
	}

	var body User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	result, err := updateUser(id, body)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	utils.SendResponse(ctx, utils.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "User updated successfully",
		Data:       result,
	})
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid id"})
		return
	}

	if err := deleteUser(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	utils.SendResponse(ctx, utils.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "User deleted successfully",
		Data:       nil,
	})
}
