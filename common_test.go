package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// setup and tear down function
func TestMain(m *testing.M) {
	// set gin to test mode
	gin.SetMode(gin.TestMode)

	// run tests
	os.Exit(m.Run())
}

// helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()

	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}

	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if !f(rr) {
		t.Fail()
	}
}

func saveLists() {
	tmpArticleList = articleList
}

func restoreLists() {
	articleList = tmpArticleList
}
