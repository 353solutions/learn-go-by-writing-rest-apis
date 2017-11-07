package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	db = map[string][]byte{}
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

func dbHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/db/"):]

	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	if r.Method == "GET" {
		value, ok := db[key]
		if !ok {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}
		w.Write(value)
		return
	}

	if r.Method == "POST" {
		defer r.Body.Close()
		value, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Can't read body", http.StatusInternalServerError)
			return
		}
		db[key] = value
		fmt.Fprintf(w, "OK\n")
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func keysHandler(w http.ResponseWriter, r *http.Request) {
	keys := []string{}

	for key := range db {
		keys = append(keys, key)
	}

	response := map[string]interface{}{
		"keys": keys,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		http.Error(w, "Can't encode keys", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/db/", dbHandler)
	http.HandleFunc("/keys", keysHandler)
	http.ListenAndServe(":8080", nil)
}
