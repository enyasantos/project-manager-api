package handler

import (
	"net/http"

	"github.com/enyasantos/project-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show project
// @Description Show a new project
// @Tags Project
// @Accept json
// @Produce json
// @Param id query string true "Project identification"
// @Success 200 {object} ShowProjectResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /project [get]
func ShowProjectHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "show-project", project)
}
