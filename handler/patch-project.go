package handler

import (
	"net/http"

	"github.com/enyasantos/project-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Patch project
// @Description Partially update a project
// @Tags Project
// @Accept json
// @Produce json
// @Param id query string true "Project Identification"
// @Param request body PatchProjectRequest true "Fields to Update"
// @Success 200 {object} UpdateProjectResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /project [patch]
func PatchProjectHandler(ctx *gin.Context) {
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

	request := PatchProjectRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		sendError(ctx, http.StatusBadRequest, "Failed to parse JSON request body")
		return
	}

	// Atualize apenas os campos que foram fornecidos na solicitação
	if request.Title != nil {
		project.Title = *request.Title
	}
	if request.Description != nil {
		project.Description = *request.Description
	}
	if request.StartDate != nil {
		project.StartDate = *request.StartDate
	}
	if request.EndDate != nil {
		project.EndDate = *request.EndDate
	}
	if request.Status != nil {
		project.Status = *request.Status
	}
	if request.Value != nil {
		project.Value = *request.Value
	}

	if err := db.Save(&project).Error; err != nil {
		logger.Errorf("error updating project: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating project")
		return
	}

	sendSuccess(ctx, "update-project", project)
}
