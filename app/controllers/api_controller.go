package controllers

import (
	"github.com/dmitry-udod/codes_go/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FopFind(c *gin.Context) {
	params := Params()
	params["id"] = c.Param("code")
	records, _ := services.SearchFop(params)

	if len(records) > 0 && records[0].FullName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": records[0],
		})
		return
	}

	EntityNotfound(c)
}

func FopLatest(c *gin.Context) {
	params := Params()
	params["page"] = c.Request.URL.Query().Get("page")
	params["q"] = c.Request.URL.Query().Get("q")
	records, metadata := services.SearchFop(params)

	c.JSON(http.StatusOK, gin.H{
		"data": records,
		"metadata": metadata,
	})
	return
}

func LegalEntitiesLatest(c *gin.Context) {
	params := Params()
	params["page"] = c.Request.URL.Query().Get("page")
	params["q"] = c.Request.URL.Query().Get("q")
	records, metadata := services.SearchLegalEntities(params)

	c.JSON(http.StatusOK, gin.H{
		"data": records,
		"metadata": metadata,
	})
	return
}

func LegalEntityFind(c *gin.Context) {
	params := Params()
	params["id"] = c.Param("code")
	records, _ := services.SearchLegalEntities(params)

	if len(records) > 0 && records[0].FullName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": records[0],
		})
		return
	}

	EntityNotfound(c)
}

func EntityNotfound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}

func Params() map[string]string {
	return make(map[string]string)
}
