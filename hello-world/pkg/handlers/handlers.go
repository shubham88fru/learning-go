package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page")

}

func addValues(x, y int) (int, error) {
	return x + y, nil
}

