package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, data gin.H, templateName string) {
	responseEncodingType := c.Request.Header.Get("accept")
	switch responseEncodingType {
	case "application/json":
		c.JSON(http.StatusOK, data)
	case "application/xml":
		c.XML(http.StatusOK, data)
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
