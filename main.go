package main

import (
	"fmt"
	"net/http"
)

// declaration of function of type http.HandlerFunc to pass it in http.HandleFunc
func handlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my CS50 final project website!</h1>")
}

func main() {
	//call to HandleFunc method to specify how requests to route "/" should be handled
	http.HandleFunc("/", handlerFunction)

	//start the server on port :3000
	fmt.Println("Starting the server on port :3000 ...")
	http.ListenAndServe(":3000", nil)

}
