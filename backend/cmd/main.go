package main

import (
	"log"

	"github.com/eXoterr/SimpleFileOrganizerBackend/internal/config"
	"github.com/eXoterr/SimpleFileOrganizerBackend/internal/handlers"
	"github.com/eXoterr/SimpleFileOrganizerBackend/internal/router"
	"github.com/eXoterr/SimpleFileOrganizerBackend/pkg/utils"
)

func main() {
	cfg := config.New()
	err := cfg.CheckConfig()
	log.Println("checking config...")
	utils.MustCheckError(err) // Fail if config is invalid
	log.Println("config valid")

	log.Printf("starting http server at http://%s", cfg.ListenAt)
	r := router.New(cfg.ListenAt)
	handlers.SetupHandlers(r, cfg)
	err = r.Start()

	utils.MustCheckError(err) // Fail if router can't start
}
