package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greet/{name}", greet).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	var name = vars["name"]
	w.Write([]byte("Hello " + name + "!"))
}
