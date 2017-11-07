package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type DB struct {
	store map[string][]byte
	lock  sync.Mutex
}

func NewDB() *DB {
	return &DB{
		store: map[string][]byte{},
	}
}

var (
	db = NewDB()
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

func getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	db.lock.Lock()
	defer db.lock.Unlock()

	value, ok := db.store[key]
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	w.Write(value)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	defer r.Body.Close()
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusInternalServerError)
		return
	}

	db.lock.Lock()
	defer db.lock.Unlock()

	db.store[key] = value
	fmt.Fprintf(w, "OK\n")
}

func keysHandler(w http.ResponseWriter, r *http.Request) {
	keys := []string{}

	db.lock.Lock()
	defer db.lock.Unlock()

	for key := range db.store {
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
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", handler)
	rtr.HandleFunc("/ping", pingHandler)
	rtr.HandleFunc("/time", timeHandler)
	rtr.HandleFunc("/db/{key}", getHandler).Methods("GET")
	rtr.HandleFunc("/db/{key}", setHandler).Methods("POST")
	rtr.HandleFunc("/keys", keysHandler)
	http.ListenAndServe(":8080", rtr)
}
