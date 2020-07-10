// models.article_test.go

package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i := range alist {
		if alist[i].ID != articleList[i].ID ||
			alist[i].Title != articleList[i].Title ||
			alist[i].Content != articleList[i].Content {
			t.Fail()
			break
		}
	}
}
