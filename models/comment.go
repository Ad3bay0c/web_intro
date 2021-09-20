package models

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Comment struct {
	ID		int64	`json:"id"`
	Comment	string	`json:"comment"`
}


func CreateComment(w http.ResponseWriter, r *http.Request){
	param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Printf("Error Converting ID or Invalid ID")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if strings.TrimSpace(template.HTMLEscapeString(r.FormValue("comment"))) == "" {
		w.WriteHeader(http.StatusBadRequest)
		message.Message = "Please type in a valid comment"
		message.Color = "danger"
		http.Redirect(w, r, fmt.Sprintf("/%d/", id), http.StatusMovedPermanently)
		return
	}

	comment := &Comment{
		ID: time.Now().Unix(),
		Comment: template.HTMLEscapeString(strings.TrimSpace(r.FormValue("comment"))),
	}
	for idx, val := range blogs.Blogs {
		if val.ID == id {
			blogs.Blogs[idx].Comments[id] = comment
			blogs.Blogs[idx].Comment++
		}
	}
}
