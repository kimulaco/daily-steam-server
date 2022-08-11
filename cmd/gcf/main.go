package main

import (
	"log"
	"net/http"
	"os"

	GetNewApp "daily-steam.app/GetNewApp"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

type Router struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var routers = []Router{
	{
		Path:    "get-new-app",
		Handler: GetNewApp.GetNewApp,
	},
}

func main() {
	p := GetPort()

	for _, r := range routers {
		functions.HTTP(r.Path, r.Handler)
	}

	err := funcframework.Start(p)
	if err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

func GetPort() string {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	return port
}
