package handlers

import (
	"github.com/eXoterr/SimpleFileOrganizerBackend/internal/config"
	"github.com/eXoterr/SimpleFileOrganizerBackend/internal/router"
)

func SetupHandlers(router *router.Router, cfg *config.Config) {
	router.Handler.GET("/files/list", ListFiles(cfg))
	router.Handler.GET("/files/library", ListLibrary(cfg))
	router.Handler.POST("/organize/film", OrganizeFilm(cfg))
}
