package controller

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
		Count   int
	}{
		Title:   "Welcome to the Counter App",
		Message: "This is the index page.",
		Count:   c.Counter.GetCount(),
	}
	err := c.View.Render(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c.Counter.Increment()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
