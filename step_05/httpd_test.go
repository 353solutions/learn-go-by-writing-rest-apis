package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error calling handler")
	}

	expected := "שלום Gophers\n"
	if string(out) != expected {
		t.Fatalf("Bad answer. Expected %q, got %q", expected, string(out))
	}
}
