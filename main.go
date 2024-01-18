package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mperalle/cs50-final-project/controllers"
	"github.com/mperalle/cs50-final-project/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	//create template from template files
	tpl, err := views.Parse(filepath)
	//handle errors during parsing
	if err != nil {
		//logs error message out to terminal
		log.Printf("processing template: %v", err)
		//writes 500 status code in the response
		http.Error(w, "There was an error processing the template.", http.StatusInternalServerError)
		//prevents further code from executing
		return
	}
	//write output of execution to the response
	tpl.Execute(w, nil)
}

// declaration of the handler function of type http.HandlerFunc to pass it in http.HandleFunc
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := "templates/home.html"
	executeTemplate(w, tplPath)
}

// declaration of the handler function for contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := "templates/contact.html"
	executeTemplate(w, tplPath)
}

// declaration of handler function for faq page
func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/faq.html")
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

	//parsing template for home page
	tpl, err := views.Parse("templates/home.html")
	if err != nil {
		panic(err)
	}
	//register handler with template already parsed for "/"
	r.Get("/", controllers.StaticHandler(tpl))

	//parsing template for contact page
	tpl, err = views.Parse("templates/contact.html")
	if err != nil {
		panic(err)
	}
	//register handler with template already parsed for "/contact"
	r.Get("/contact", controllers.StaticHandler(tpl))

	//parsing template for faq page
	tpl, err = views.Parse("templates/faq.html")
	if err != nil {
		panic(err)
	}
	//register handler with template already parsed for "/faq"
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.Get("/galleries/{userID}", galleriesHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Error Page not found", http.StatusNotFound)
	})

	//start the server on port :3000
	fmt.Println("Starting the server on port :3000 ...")
	//pass the routeur to ListenAndServe
	http.ListenAndServe(":3000", r)
}
