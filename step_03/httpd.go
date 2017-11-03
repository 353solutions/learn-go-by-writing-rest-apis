package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "שלום Gophers\n")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	msg := r.URL.Query().Get("msg")
	if msg == "" {
		msg = "PONG"
	}

	fmt.Fprintf(w, msg)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	now := time.Now()
	fmt.Fprintf(w, now.Format(time.RFC3339))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}
