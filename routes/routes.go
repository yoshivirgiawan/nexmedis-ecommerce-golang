package routes

import (
	"ecommerce/app/http/middlewares"
	v1 "ecommerce/routes/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	r.Static("/public", "./storage/public")
	r.Use(middlewares.RequestLogger())
	// Enable automatic redirect for routes with or without trailing slash
	r.RedirectTrailingSlash = true
	v1.SetupRouter(r)
	return r
}
