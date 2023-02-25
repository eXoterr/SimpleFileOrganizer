package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	Addr    string
	Handler *httprouter.Router
}

func New(addr string) *Router {
	router := httprouter.New()
	router = SetCors(router)
	return &Router{
		Addr:    addr,
		Handler: router,
	}
}

func (r *Router) Start() error {
	return http.ListenAndServe(r.Addr, r.Handler)
}
