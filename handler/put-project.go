package handler

import (
	"net/http"

	"github.com/enyasantos/project-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update project
// @Description Update a project
// @Tags Project
// @Accept json
// @Produce json
// @Param id query string true "Project Identification"
// @Param request body PutProjectRequest true "Opening data to Update"
// @Success 200 {object} UpdateProjectResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /project [put]
func PutProjectHandler(ctx *gin.Context) {
	request := PutProjectRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		sendError(ctx, http.StatusBadRequest, "Failed to parse JSON request body")
		return
	}

	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	project := schemas.Project{}

	if err := db.First(&project, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "project not found")
		return
	}

	project.Title = request.Title
	project.Description = request.Description
	project.StartDate = request.StartDate
	project.EndDate = request.EndDate
	project.Status = request.Status
	project.Value = request.Value

	if err := db.Save(&project).Error; err != nil {
		logger.Errorf("error updating project: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating project")
		return
	}

	sendSuccess(ctx, "update-project", request)
}
