package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/static/", staticHandler)
	panic(http.ListenAndServe(":8080", nil))
}

// homeHandler handles normal requests for non-static files.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head></head><body><h1>Welcome Home!</h1></body></html>")
}

// staticHandler handles requests for static files on server.
func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
