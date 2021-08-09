package test

import (
	"ai4am.com/go-restapi/internal"
	"ai4am.com/go-restapi/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app *gin.Engine

func Setup() {
	if app == nil {
		configuration := config.GetConfiguration("config.yml")
		a := internal.App{Config: configuration, App: nil}
		a.Start("test.log")
		app = a.App
	}
}

func TestGetAllUsers(t *testing.T) {
	Setup()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(fmt.Sprintf("%s/api/users", server.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}