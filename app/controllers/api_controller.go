package controllers

import (
	"github.com/dmitry-udod/codes_go/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindFopById(c *gin.Context) {
	code := c.Param("code")

	record := services.SearchFop(code)

	if record.FullName != "" {
		c.JSON(http.StatusOK, record)
		return
	}

	EntityNotfound(c)
}

func EntityNotfound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}
