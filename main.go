package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"
	"sebcio/conf"
	"sebcio/repository/users"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/users/", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("postgres", conf.POSTGRES_URI)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		userRepository := users.UserRepository{db}
		rows, err := userRepository.GetAll()

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
