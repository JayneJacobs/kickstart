package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Guest struct {
	Name string
}

// Handler for the hello-gopher route
func HelloGuestHandler(w http.ResponseWriter, r *http.Request) {

	var guestname string
	guestname = r.URL.Query().Get("guestname")
	if guestname == "" {
		guestname = "Jayne"
	}
	guest := Guest{Name: guestname}
	renderTemplate(w, "templates/greeting.html", guest)

}

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}
