package piscine

import (
	"html/template"
	"net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	// call functions that will get and store the artist info

	tmpl, err := template.ParseFiles("template/mainpage.html")
	if err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}
	// up until here is for main page
	/* -------------------------------------------------------------- */
	//
	if r.Method == http.MethodPost {
		if r.URL.Path != "/artist" {
			http.Error(w, "404 - Page Not Found", http.StatusNotFound)
			return
		}

		// call functions that will collect all data

		tmpl, err := template.ParseFiles("template/artistpage.html")
		if err != nil {
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		}
	}
}
