package handlers

import (
	"net/http"
	"os"

	"github.com/eXoterr/SimpleFileOrganizerBackend/internal/config"
	"github.com/eXoterr/SimpleFileOrganizerBackend/internal/processors"
	"github.com/julienschmidt/httprouter"
)

func OrganizeFilm(cfg *config.Config) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		path := r.PostFormValue("path")

		finalName := processors.OrganizeFilm(path)

		os.Create(cfg.LibraryPath + "/" + finalName)

		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Origin", "*")
	}
}
