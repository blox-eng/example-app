package main

import (
	"net/http"

	"github.com/blox-eng/app/config"
	"github.com/blox-eng/app/internal/server"
	log "github.com/sirupsen/logrus"
)

//	@title			BLOX
//	@version		1.0
//	@description	This is the backend of Blox App service.

// @contact.name	API Support
// @contact.email	support@blox.io
func main() {
	s := server.New()
	log.Info("Listening on port:", config.GetYamlValues().ServerConfig.Port)
	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}
	log.Info("service stopped")
}
