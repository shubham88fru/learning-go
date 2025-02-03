package render

import (
	"fmt"
	"html/template"
	"net/http"
)


func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl + ".html", "templates/base.layout.html")
	err := parsedTemplate.ExecuteTemplate(w, tmpl+".html", nil)

	if err != nil {
		fmt.Print("Error parsing template: ", err)
	}
}