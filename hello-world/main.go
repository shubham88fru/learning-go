package main

import (
	"net/http"
)

const port = ":8080"

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