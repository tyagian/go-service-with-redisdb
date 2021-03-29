package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello World")
//}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/goodbye", goodbyeHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Good Bye!")
}
