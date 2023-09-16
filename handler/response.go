package handler

import (
	"fmt"
	"net/http"

	"github.com/enyasantos/project-manager/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateProjectResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProjectResponse `json:"data"`
}

type DeleteProjectResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProjectResponse `json:"data"`
}

type ShowProjectResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProjectResponse `json:"data"`
}

type IndexProjectResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProjectResponse `json:"data"`
}

type UpdateProjectResponse struct {
	Message string                  `json:"message"`
	Data    schemas.ProjectResponse `json:"data"`
}
