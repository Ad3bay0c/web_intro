package models

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	blogs = Blogs{}
	message = Message{}
)

func init() {

}
func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/home.gtpl", "frontend/navbar.gtpl")
	if err != nil {
		log.Printf("Error Opening File: %v", err.Error())
		return
	}
	message.Data = blogs

	err = t.Execute(w, message)
	message.Message = ""
	message.Color = ""

	if err != nil {
		log.Printf("Error Opening File: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateBlog(w http.ResponseWriter, r *http.Request)  {
	title := strings.TrimSpace(r.FormValue("title"))
	details := strings.TrimSpace(r.FormValue("details"))

	blog := Blog{
		ID:      time.Now().Unix(),
		Title:   title,
		Details: details,
		Date:    time.Now().Local(),
	}
	blogs.Blogs = append(blogs.Blogs, blog)
	err := blogs.addToFile()
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message.Message = "Created Successfully"
	message.Color 	= "success"
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		message.Message = "Not a Valid ID"
		message.Color = "danger"
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

	for idx, v := range blogs.Blogs {
		if v.ID == id && len(blogs.Blogs)-1 != idx {
			blogs.Blogs = append(blogs.Blogs[:idx], blogs.Blogs[idx+1:]...)
		}else if v.ID == id && len(blogs.Blogs)-1 == idx{
			blogs.Blogs = blogs.Blogs[:idx]
		}
	}
	err = blogs.addToFile()
	if err != nil {
		log.Printf("%v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message.Message = "Deleted Successfully"
	message.Color = "success"

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
