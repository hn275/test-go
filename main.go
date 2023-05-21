package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type MyThing struct {
	Foo string `json:"foo"`
}

func main() {
	r := chi.NewMux()

	r.Handle("/foo", http.HandlerFunc(handleDeez))

	log.Println("listen")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func handleDeez(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var buf MyThing
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		log.Fatal(err)
	}

	buf.Foo = "nutz"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTeapot)
	json.NewEncoder(w).Encode(&buf)
}
