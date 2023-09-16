package handler

import (
	"net/http"

	"github.com/enyasantos/project-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List project
// @Description List all project
// @Tags Project
// @Accept json
// @Produce json
// @Success 200 {object} CreateOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func IndexProjectHandler(ctx *gin.Context) {
	projects := []schemas.Project{}

	if err := db.Find(&projects).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing projects")
		return
	}

	sendSuccess(ctx, "list-projects", projects)
}
