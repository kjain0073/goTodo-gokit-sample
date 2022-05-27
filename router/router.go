package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/kjain0073/go-Todo/middlewareTodo"
)

func TodoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", middlewareTodo.FetchTodos)
		r.Post("/", middlewareTodo.CreateTodo)
		r.Put("/{id}", middlewareTodo.UpdateTodo)
		r.Delete("/{id}", middlewareTodo.DeleteTodo)
	})
	return rg
}
