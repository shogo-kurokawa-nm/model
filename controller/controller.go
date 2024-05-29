package controller

import (
	"html/template"
	"model/model"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("src/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	counter, err := model.GetSessionValue(r, "counter")
	if err != nil {
		counter = 0
	}
	tmpl.Execute(w, counter)
}

func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	counter, _ := model.GetSessionValue(r, "counter")
	if counter == nil {
		counter = 0
	}
	model.SetSessionValue(r, w, "counter", counter.(int)+1)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
