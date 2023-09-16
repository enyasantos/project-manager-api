package server

import (
	"log"
	"net/http"
	"os"

	"github.com/enyasantos/project-manager/router"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.New()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	router.InitializeRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
