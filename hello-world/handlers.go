package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page")

}

func addValues(x, y int) (int, error) {
	return x + y, nil
}

