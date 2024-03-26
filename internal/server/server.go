package server

import (
	"fmt"

	"net"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/blox-eng/app/cmd/web"
	"github.com/blox-eng/app/config"
	_ "github.com/blox-eng/app/docs"
	"github.com/blox-eng/app/internal/handler"
	"github.com/blox-eng/app/internal/service"
	"github.com/blox-eng/app/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// TODO: Add authentication to all routes
// TODO: Add authorization to all routes
type Server struct {
	httpServer *http.Server
	router     *chi.Mux
}

func New() *Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
	)
	storage.InitializeDatabaseConnector()

	pgClient := storage.NewClient(
		&storage.Config{
			DBConnection: "",
		})

	s := service.NewService(pgClient)
	setupRoutesForUpdate(s, r)
	server := newServer(r)

	return server
}

func setupRoutesForUpdate(service service.Service, r *chi.Mux) {

	// plug in sub-routers for resources: feature gate
	// this pattern also allows for easy integration testing. see api_test.go

	r.Route("/api", func(r chi.Router) {
		r.Mount("/docs", httpSwagger.Handler(
			httpSwagger.URL(
				fmt.Sprintf(
					"http://%s:%s/api/docs/doc.json",
					config.GetYamlValues().ServerConfig.Server,
					config.GetYamlValues().ServerConfig.Port,
				))))
		r.Mount("/", handler.Handler(service))
	})

	// Serve static files
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/js/*", fileServer)
	r.Get("/", templ.Handler(web.Index()).ServeHTTP)
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", ":"+s.httpServer.Addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	return s.httpServer.Serve(l)
}

func newServer(r *chi.Mux) *Server {
	fmt.Println("****Server Started on", config.GetYamlValues().ServerConfig.Port, "****")
	return &Server{
		httpServer: &http.Server{Addr: config.GetYamlValues().ServerConfig.Port, Handler: r},
		router:     r,
	}
}
