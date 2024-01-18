package controllers

import (
	"net/http"

	"github.com/mperalle/cs50-final-project/views"
)

// function with a closure which return a HandlerFunc
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
