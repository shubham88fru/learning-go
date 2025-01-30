package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	renderTemplate(w, "home.page")
}

func About(w http.ResponseWriter, r *http.Request) {
	// sum, _ := addValues(2, 3)
	// w.Write([]byte("About Page" + strconv.Itoa(sum)))
	renderTemplate(w, "about.page")

}

func addValues(x, y int) (int, error) {
	return x + y, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl + ".html")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Print("Error parsing template: ", err)
	}
}

func main() {

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		return
	}
}