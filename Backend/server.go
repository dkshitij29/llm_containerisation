package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
)

var tmpl = template.Must(template.ParseFiles("../Frontend/index.html"))

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit", submitHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Frontend"))))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		textInput := r.FormValue("textInput")
		response := Response{Message: fmt.Sprintf("You entered: %s", textInput)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
