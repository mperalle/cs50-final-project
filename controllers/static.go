package controllers

import (
	"html/template"
	"net/http"

	"github.com/mperalle/cs50-final-project/views"
)

// function with a closure which return a HandlerFunc
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

// function to return the FAQ page
func FaqHandler(tpl views.Template) http.HandlerFunc {
	//hardcode of data to pass to the template
	//je veux faire quoi ? un slice de struct
	questions := []struct {
		Question string
		Answer   template.HTML
	}{{Question: "Is this web app for free?",
		Answer: "Yes it is !",
	}, {
		Question: "How can I contact you?",
		Answer:   `Email us - <a href="mailto:support@cs50.com">support@cs50.com</a>`,
	}, {
		Question: "What are your support hours?",
		Answer:   "The online support is open 24/7",
	},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
