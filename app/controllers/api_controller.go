package controllers

import (
	"github.com/dmitry-udod/codes_go/app/services"
	"github.com/gin-gonic/gin"
)

func FindFopById(c *gin.Context) {
	code := c.Param("code")

	record := services.SearchFop(code)

	if record.FullName != "" {
		c.JSON(200, record)
		return
	}

	EntityNotfound(c)
}

func EntityNotfound(c *gin.Context) {
	c.AbortWithStatusJSON(404, gin.H{"error": "Not Found"})
}
