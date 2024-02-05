package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

// Regular Page Controllers
func HomeController(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("../templates/main/home.html")
	if err != nil {
		http.Redirect(res, req, "/500", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(res, nil)
	if err != nil {
		http.Redirect(res, req, "/500", http.StatusInternalServerError)
		return
	}

}

// Error Page Controllers
func ErrorBadRequestController(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("asciiartweb/templates/errors/400.html")
	if err != nil {
		http.Error(res, "Well that's embarrassing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "Well that's embarrassing", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(res, "Welcome to the home page!")
}

func ErrorNotFoundController(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("../templates/errors/404.html")
	if err != nil {
		http.Error(res, "Well that's embarrassing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "Well that's embarrassing", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(res, "Welcome to the home page!")
}

func ErrorInternalController(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("../templates/errors/500.html")
	if err != nil {
		http.Error(res, "Well that's embarrassing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "Well that's embarrassing", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(res, "Welcome to the home page!")
}
