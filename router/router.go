package router

import (
	"github.com/togo-mentor/go-practice-app/handlers"
	"github.com/go-chi/chi"
)

func Get() *chi.Mux {
	r := chi.NewRouter()

	todoHandler := handler.NewTodoHandler()
	r.Route("/todos", func(r chi.Router) {
		r.Post("/", todoHandler.Create)
		// r.Get("/detail", todoHandler.Get)
	})
	return r
}