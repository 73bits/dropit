package main

import (
	"log"
	"net/http"

	"github.com/73bits/dropit/internal/handler"
	"github.com/73bits/dropit/internal/repo"
	"github.com/73bits/dropit/internal/service"
)

func main() {
	mux := http.NewServeMux()

	repo, _ := repo.NewJSONRepo("data.json")
	service := service.NewTextService(repo)
	handler := handler.NewTextHandler(service)

	mux.HandleFunc("POST /text", handler.Create)
	mux.HandleFunc("GET /text/", handler.Get)

	log.Println("app running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
