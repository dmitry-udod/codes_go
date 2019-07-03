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

	router.Static("/public", "./public")
	router.LoadHTMLGlob("app/views/*")

	v1 := router.Group("/api/v1")
	{
		if os.Getenv("GIN_MODE") == "release" {
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

		v1.GET("/fop/:code", controllers.FindFopById)
	}

	router.GET("/", controllers.Main)

	return router
}