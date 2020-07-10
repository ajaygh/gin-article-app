//routes.go
package main

func initRoutes() {
	router.GET("/", showIndexPage)

	// route to view a article
	router.GET("/article/view/:articleID", getArticle)
}
