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

		v1.GET("/fop/view/:code", controllers.FopSearch)
		v1.GET("/fop/search/:q", controllers.FopSearch)
		v1.GET("/fop/latest", controllers.FopLatest)
	}

	router.GET("/", controllers.Main)

	return router
}