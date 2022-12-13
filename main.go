package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/users/", func(w http.ResponseWriter, r *http.Request) {
		conn := initSQLConn()
		conn := UserRepository()
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
