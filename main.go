package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello World")
//}

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello World")
	templates.ExecuteTemplate(w, "index.html", nil)
}
