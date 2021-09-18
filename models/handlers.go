package models

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/home.gtpl")
	if err != nil {
		log.Printf("Error Opening File: %v", err.Error())
		return
	}

	t.Execute(w, "Adebayo")
}