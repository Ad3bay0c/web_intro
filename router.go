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
	r.Get("/blogs", models.BlogPage)
	r.Post("/create", models.CreateBlog)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/delete", models.DeleteBlog)
		r.Get("/edit", models.EditBlog)
		r.Post("/update", models.UpdateBlog)
		r.Get("/", models.ViewBlog)
	})

	log.Printf("Server Started at Localhost:4000")
	log.Fatalf("%v", http.ListenAndServe(":4000", r))
}
