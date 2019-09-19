package router

import (
	"github.com/cnjack/throttle"
	"github.com/dmitry-udod/codes_go/app/controllers"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func SetupRouter() *gin.Engine  {
	router := gin.New()
	mode := os.Getenv("GIN_MODE")

	router.Static("/public", "./public")
	router.StaticFile("sitemap.txt", "./public/sitemap.txt")

	if _, err := os.Stat("app/views"); ! os.IsNotExist(err) {
		router.LoadHTMLGlob("app/views/*")
	}

	v1 := router.Group("/api/v1")
	{
		if mode == "release" {
			v1.Use(throttle.Policy(&throttle.Quota{
				Limit:  60,
				Within: time.Minute,
			}))
		}

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// FOP
		v1.GET("/fop/view/:code", controllers.FopFind)
		v1.GET("/fop/latest", controllers.FopLatest)

		// Legal Entities
		v1.GET("/legal-entities/latest", controllers.LegalEntitiesLatest)
		v1.GET("/legal-entities/view/:code", controllers.LegalEntityFind)

		// Legal Entities
		v1.GET("/terrorists", controllers.Terrorists)
		v1.GET("/terrorists/view/:code", controllers.TerroristFind)
	}

	router.GET("/", controllers.Main)

	router.NoRoute(func(c *gin.Context){
		c.File("app/views/index.tmpl")
	})

	return router
}