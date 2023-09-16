package handler

import (
	"net/http"

	"github.com/enyasantos/project-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create project
// @Description Create a new project
// @Tags Project
// @Accept json
// @Produce json
// @Param request body CreateProjectRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateProjectHandler(ctx *gin.Context) {
	request := CreateProjectRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		sendError(ctx, http.StatusBadRequest, "Failed to parse JSON request body")
		return
	}

	project := schemas.Project{
		Title:       request.Title,
		Description: request.Description,
		StartDate:   request.StartDate,
		EndDate:     *request.EndDate,
		Status:      request.Status,
		Value:       *request.Value,
	}

	if err := db.Create(&project).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-project", project)
}
