package main

import (
	"github.com/quackBear57/golang_stackademic/backend/api"	// add api package to call the controller
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Mount("/todo", api.TodoResources{}.Routes())
	http.ListenAndServe(":5000", r)
}