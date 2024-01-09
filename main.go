package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// declaration of the handler function of type http.HandlerFunc to pass it in http.HandleFunc
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//explicitely set the header content-type to html
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	//parse HTML template and execute it
	tpl, err := template.ParseFiles("templates/home.gohtml")
	//handle error in html parsing
	if err != nil {
		//logs error message out to terminal
		log.Printf("parsing template: %v", err)
		//writes 500 status code in the response
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		//prevents further code from executing
		return
	}

	err = tpl.Execute(w, nil)
	//handle errors during execution
	if err != nil {
		//logs error message out to terminal
		log.Printf("executing template: %v", err)
		//writes 500 status code in the response
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		//prevents further code from executing
		return
	}
}

// declaration of the handler function for contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	//explicitely set the header content-type to html
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	//parse HTML template and execute it
	tpl, err := template.ParseFiles("templates/contact.gohtml")
	//handle error in html parsing
	if err != nil {
		//logs error message out to terminal
		log.Printf("parsing template: %v", err)
		//writes 500 status code in the response
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		//prevents further code from executing
		return
	}

	err = tpl.Execute(w, nil)
	//handle errors during execution
	if err != nil {
		//logs error message out to terminal
		log.Printf("executing template: %v", err)
		//writes 500 status code in the response
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		//prevents further code from executing
		return
	}
}

// declaration of handler function for faq page
func faqHandler(w http.ResponseWriter, r *http.Request) {
	//set header content-type to html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//writing the HTML response
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li> <b>Is this application free to use ?</b> Yes it is completly free to use !</li>
		<li> <b>How can I get in touch with you ?</b> You can contact me via the email in the contact page</li>
	</ul>`)
}

func galleriesHandler(w http.ResponseWriter, r *http.Request) {
	//set header content-type to html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userID := chi.URLParam(r, "userID")
	//writing the HTML response
	w.Write([]byte(fmt.Sprintf("The userID is %v", userID)))
}
func main() {
	//register handler functions to a new Chi Routeur
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/galleries/{userID}", galleriesHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Error Page not found", http.StatusNotFound)
	})
	//start the server on port :3000
	fmt.Println("Starting the server on port :3000 ...")
	//pass the routeur to ListenAndServe
	http.ListenAndServe(":3000", r)
}
