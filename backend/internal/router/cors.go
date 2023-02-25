package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetCors(router *httprouter.Router) *httprouter.Router {
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Origin", "*")

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	return router
}
