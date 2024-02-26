package main

import (
    "github.com/blox-eng/work/config"
    "github.com/blox-eng/work/server"
    log "github.com/sirupsen/logrus"
    "net/http"
)

func main() {
    s := server.New()
    log.Info("Listening on port:", config.GetYamlValues().ServerConfig.Port)
    err := s.ListenAndServe()
    if err != http.ErrServerClosed {
        log.Fatalf("Listen: %s\n", err)
    }
    log.Info("service stopped")

}
