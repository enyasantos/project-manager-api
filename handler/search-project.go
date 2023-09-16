package handler

import (
	"net/http"
	"time"

	"github.com/enyasantos/project-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Search projects
// @Description Search projects based on specified criteria.
// @Tags Project
// @Accept json
// @Produce json
// @Param title query string false "Title of the project"
// @Param status query string false "Status of the project"
// @Param start_date query string false "Start date (YYYY-MM-DD) of the project"
// @Param end_date query string false "End date (YYYY-MM-DD) of the project"
// @Param order_by query string false "Order by 'value' (asc or desc)"
// @Success 200 {array} Project
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /projects [get]

func SearchProjectHandler(ctx *gin.Context) {
	projects := []schemas.Project{}

	query := db.Model(&schemas.Project{})

	if title := ctx.Query("title"); title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if status := ctx.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err == nil {
			query = query.Where("start_date >= ?", startDate)
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err == nil {
			query = query.Where("end_date <= ?", endDate)
		}
	}

	if err := query.Find(&projects).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if orderBy := ctx.Query("order_by"); orderBy != "" {
		switch orderBy {
		case "asc":
			query = query.Order("value ASC")
		case "desc":
			query = query.Order("value DESC")
		default:
			//invalid values
		}
	}

	if err := query.Find(&projects).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "project not found")
		return
	}

	sendSuccess(ctx, "search-projects", projects)
}
