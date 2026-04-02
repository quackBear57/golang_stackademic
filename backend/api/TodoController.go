package api

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

type TodoResources struct {}

// Routes creates a REST router for the todos resource
func (rs TodoResources) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middelware..

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)	// GET /todos/{id} - read a single todo
	})

	return r
}

func (rs TodoResources) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("todo get " + id))
}