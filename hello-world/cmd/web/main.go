package main

import (
	"myapp/pkg/handlers"
	"net/http"
)

const port = ":8080"

func main() {

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		return
	}
}