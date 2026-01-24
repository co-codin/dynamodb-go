package routes

import "github.com/go-chi/chi"

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		
	}
}