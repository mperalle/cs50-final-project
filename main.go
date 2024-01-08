package main

import (
	"fmt"
	"net/http"
)

// declaration of the handler function of type http.HandlerFunc to pass it in http.HandleFunc
func handlerFunction(w http.ResponseWriter, r *http.Request) {
	//explicitely set the header content-type to html
	w.Header().Set("Content-Type", "text/html")
	//writing to ResponseWriter
	fmt.Fprint(w, "<h1>Welcome to my CS50 final project website!</h1>")
}

func main() {
	//register a handler to handle all incoming web requests
	http.HandleFunc("/", handlerFunction)

	//start the server on port :3000
	fmt.Println("Starting the server on port :3000 ...")
	http.ListenAndServe(":3000", nil)

}
