package main

import (
	"fmt"
	"github.com/Ad3bay0c/web_intro/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestModel_ViewBlog(t *testing.T) {
	req, err := http.NewRequest("GET", fmt.Sprintf("localhost:4000/%d/", 1632180490), nil)
	if err != nil {
		t.Fatalf("Couldn't create a request: %v", err.Error())
	}
	rec := httptest.NewRecorder()

	models.ViewBlog(rec, req)
	result := rec.Result()
	if result.StatusCode == http.StatusOK {
		t.Errorf("Expected status Ok, Got %v", result.StatusCode)
	}
	defer result.Body.Close()
	b, err := ioutil.ReadAll(result.Body)

	if err != nil {
		t.Fatalf("Could not read Response: %v", err.Error())
	}
	t.Fatalf("%v", string(b))
}
