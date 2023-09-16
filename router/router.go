package router

import (
	"github.com/enyasantos/project-manager/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	v1 := router.Group(basePath)

	projects := v1.Group("projects")
	{
		projects.GET("", handler.IndexProjectHandler)
	}

	project := v1.Group("project")
	{
		project.POST("", handler.CreateProjectHandler)
		project.GET("/:id", handler.ShowProjectHandler)
		project.DELETE("/:id", handler.DeleteProjectHandler)
		project.PUT("/:id", handler.PutProjectHandler)
		project.PATCH("/:id", handler.PatchProjectHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
