package main

import (
	"net/http"
	"strconv"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValues(2, 3)
	w.Write([]byte("About Page" + strconv.Itoa(sum)))
}

func addValues(x, y int) (int, error) {
	return x + y, nil
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