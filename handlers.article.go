// handlers.article.go

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	//render html template
	render(c,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
		"index.html",
	)
}

func getArticle(c *gin.Context) {
	// Check if article id is valid
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		log.Println("article id not int or invalid int", err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	article, err := getArticleByID(articleID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		log.Println("no article with given article id", articleID, err)
		return
	}

	render(c,
		gin.H{
			"title":   article.Title,
			"content": article.Content,
		}, "article.html")
}
