package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil &&
			strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestGetArticleJSON(t *testing.T) {
	r := getRouter(true)

	articleID := 1
	route := fmt.Sprintf("/article/view/%d", articleID)
	r.GET("/article/view/:articleID", getArticle)

	// Create a request to send to the above route
	req, _ := http.NewRequest(http.MethodGet, route, nil)
	req.Header.Set("accept", "application/json")

	expectedArticle, _ := getArticleByID(articleID)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		// Test that the article decoded is same

		var receivedArticle article
		err := json.NewDecoder(w.Body).Decode(&receivedArticle)

		pageOK := err == nil &&
			isArticleSame(expectedArticle, &receivedArticle)

		fmt.Printf("expected Article %+v, received article: %+v\n", expectedArticle, receivedArticle)
		return statusOK && pageOK
	})
}

func isArticleSame(a, b *article) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	return a.ID == b.ID && a.Title == b.Title && a.Content == b.Content
}
