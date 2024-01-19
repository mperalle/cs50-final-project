package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mperalle/cs50-final-project/controllers"
	"github.com/mperalle/cs50-final-project/templates"
	"github.com/mperalle/cs50-final-project/views"
)

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

	//register handler with template already parsed for "/"
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "layout-page.html", "home-page.html"))))

	//register handler with template already parsed for "/contact"
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.html"))))

	//register handler with template already parsed for "/faq"
	r.Get("/faq", controllers.FaqHandler(views.Must(views.ParseFS(templates.FS, "faq.html"))))

	//register handler with template already parsed for "/login"
	r.Get("/login", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "login.html"))))

	r.Get("/galleries/{userID}", galleriesHandler)

	//handle error
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Error Page not found", http.StatusNotFound)
	})

	//start the server on port :3000
	fmt.Println("Starting the server on port :3000 ...")
	//pass the routeur to ListenAndServe
	http.ListenAndServe(":3000", r)
}
