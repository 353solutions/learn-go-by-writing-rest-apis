package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	db = make(map[string]interface{})
)

// HTTPJSONError writes HTTP error as JSON
func HTTPJSONError(w http.ResponseWriter, code http.ConnState, err error) {
	w.WriteHeader(int(code))
	enc := json.NewEncoder(w)
	enc.Encode(map[string]string{"error": err.Error()})
}

// SetHandler implements SET
func SetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if len(key) == 0 {
		HTTPJSONError(w, http.StatusBadRequest, fmt.Errorf("missing key"))
		return
	}

	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	var request struct {
		Value interface{} `json:"value"`
	}
	if err := dec.Decode(&request); err != nil {
		HTTPJSONError(w, http.StatusBadRequest, err)
		return
	}

	db[key] = request.Value
	json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
}

// GetHandler implements SET
func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if len(key) == 0 {
		HTTPJSONError(w, http.StatusBadRequest, fmt.Errorf("missing key"))
		return
	}

	val, ok := db[key]
	if !ok {
		val = nil
	}
	reply := map[string]interface{}{
		"key":   key,
		"value": val,
	}
	json.NewEncoder(w).Encode(reply)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/{key}", SetHandler).Methods("POST")
	rtr.HandleFunc("/{key}", GetHandler).Methods("GET")
	http.ListenAndServe(":8080", rtr)
}
