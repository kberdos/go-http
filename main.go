package main

import (
	"fmt"
	"github.com/kberdos/go-http/utils"
	"log"
	"net/http"
)

// The portNum that is being listened on
const portNum string = ":8080"

/*
Handles incoming HTTP requests.
writer: writes a response back to the client
req: houses the incoming request
*/

func Home(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Homepage\n")
}

func About(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "About\n")
}

func main() {
	log.Println("Starting up the http server")

	// registering the handler functions, creating paths
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// this route is being handled with a handler from an import
	http.HandleFunc("/team", utils.Team)

	// this route renders some HTMl
	http.HandleFunc("/content", utils.Content)

	log.Println("Started on port", portNum)
	fmt.Println("To close connection, CTRL+C")

	// set up the server. The second argument is nil to give us the standard ServeMux()
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
