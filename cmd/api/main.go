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
	// Create a new instance of Server
	s := server.New()

	// Print the port the server is listening on
	log.Info("Listening on port:", config.GetYamlValues().ServerConfig.Port)

	// Start the server
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		// Log any errors
		log.Fatalf("Listen: %s\n", err)
	}

	// Log a message when the server stops
	log.Info("Service stopped")
}
