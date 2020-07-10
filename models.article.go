// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{
		ID: 1, Title: "Spiritual", Content: "we don't feel real self",
	},
	{
		ID: 2, Title: "Macbeth", Content: "a pessimistic lover",
	},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range getAllArticles() {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("article not found")
}
