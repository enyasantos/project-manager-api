package handler

import (
	"fmt"
	"net/http"

	"github.com/enyasantos/project-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete project
// @Description Delete a new project
// @Tags Project
// @Accept json
// @Produce json
// @Param id query string true "Project identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func DeleteProjectHandler(ctx *gin.Context) {
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

	if err := db.Delete(&project).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting project with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-project", project)
}
