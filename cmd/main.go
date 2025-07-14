package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		// fmt.Printf("ERROR %v", err)
		fmt.Println("ERROR", err)
	}

}
