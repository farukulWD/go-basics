package utils

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int    `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func SendResponse(ctx *gin.Context, response Response) {
	ctx.JSON(response.StatusCode, response)
}
