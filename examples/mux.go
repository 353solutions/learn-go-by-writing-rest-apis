package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	fmt.Fprintf(w, "Hello %s\n", name)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/hi/{name}", hiHandler)
	http.ListenAndServe(":8080", rtr)
}
