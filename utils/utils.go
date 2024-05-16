package utils

import (
	"fmt"
	"net/http"
)

// Just another standard route
func Team(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Our Team\n")
}

// A route that serves some html
func Content(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./content" {
		p = "./static/index.html"
		fmt.Println(p)
	}
	// fs := http.FileServer(http.Dir("../static"))
	// http.StripPrefix("/content/", fs).ServeHTTP(writer, r)
	http.ServeFile(w, r, p)
}
