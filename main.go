package main

import (
	"fmt"
	"net/http"
)

// declaration of the handler function of type http.HandlerFunc to pass it in http.HandleFunc
func homeFunction(w http.ResponseWriter, r *http.Request) {
	//explicitely set the header content-type to html
	w.Header().Set("Content-Type", "text/html")
	//writing the HTML response
	fmt.Fprint(w, "<h1>Welcome to my CS50 final project website!</h1>")
}

// declaration of the handler function for contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	//set header content-type to html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//writing the HTML response
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:marvin@cs50.com\">marvin@cs50.com</a>.</p>")
}

func main() {
	//register a handler function to handle all incoming web requests
	http.HandleFunc("/", homeFunction)
	//register a handler function for contact page
	http.HandleFunc("/contact", contactHandler)
	//start the server on port :3000
	fmt.Println("Starting the server on port :3000 ...")
	http.ListenAndServe(":3000", nil)

}
