package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	//set the header content-type to html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	//handle errors during parsing
	if err != nil {
		//logs error message out to terminal
		log.Printf("processing template: %v", err)
		//writes 500 status code in the response
		http.Error(w, "There was an error processing the template.", http.StatusInternalServerError)
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

// declaration of the handler function of type http.HandlerFunc to pass it in http.HandleFunc
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := "templates/home.gohtml"
	executeTemplate(w, tplPath)
}

// declaration of the handler function for contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := "templates/contact.gohtml"
	executeTemplate(w, tplPath)
}

// declaration of handler function for faq page
func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/faq.gohtml")
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
