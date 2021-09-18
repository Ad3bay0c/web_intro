package main

import (
	"github.com/Ad3bay0c/web_intro/models"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", models.Home)
	r.Post("/create", models.CreateBlog)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/delete", models.DeleteBlog)
	})

	log.Printf("Server Started at Localhost:4000")
	log.Fatalf("%v", http.ListenAndServe(":4000", r))
}
