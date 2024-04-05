package server

import (
	"fmt"

	"net"
	"net/http"
	"os"

	"github.com/blox-eng/app/config"
	_ "github.com/blox-eng/app/docs"
	// apihandler "github.com/blox-eng/app/internal/handler/api"
	uihandler "github.com/blox-eng/app/internal/handler/ui"
	"github.com/blox-eng/app/internal/hash"
	m "github.com/blox-eng/app/internal/middleware"
	"github.com/blox-eng/app/internal/service"
	"github.com/blox-eng/app/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	httpServer *http.Server
	router     *chi.Mux
}

func New() *Server {
	r := chi.NewRouter()
	storage.InitializeDatabaseConnector()

	pgClient := storage.NewClient(
		&storage.Config{
			DBConnection: "",
		})

	s := service.NewService(pgClient)
	passwordHasher := hash.NewHPasswordHash()

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	authMiddleware := m.NewAuthMiddleware(s, "session")

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			middleware.RequestID,
			m.TextHTMLMiddleware,
			m.CSPMiddleware,
			authMiddleware.AddUserToContext,
		)
		r.NotFound(uihandler.NewNotFoundHandler().ServeHTTP)

		r.Get("/", uihandler.NewHomeHandler().ServeHTTP)

		r.Get("/about", uihandler.NewAboutHandler().ServeHTTP)

		r.Get("/register", uihandler.NewGetRegisterHandler().ServeHTTP)

		r.Post("/register", uihandler.NewPostRegisterHandler(uihandler.PostRegisterHandlerParams{
			Store: s,
		}).ServeHTTP)
		r.Get("/login", uihandler.NewGetLoginHandler().ServeHTTP)

		r.Post("/login", uihandler.NewPostLoginHandler(uihandler.PostLoginHandlerParams{
			Store:             s,
			PasswordHasher:    passwordHasher,
			SessionCookieName: "session",
		}).ServeHTTP)

		r.Post("/logout", uihandler.NewPostLogoutHandler(uihandler.PostLogoutHandlerParams{
			SessionCookieName: "session",
		}).ServeHTTP)
	})

	//	r.Route("/api", func(r chi.Router) {
	//		r.Mount("/docs", httpSwagger.Handler(
	//			httpSwagger.URL(
	//				fmt.Sprintf(
	//					"http://%s:%s/api/docs/doc.json",
	//					config.GetYamlValues().ServerConfig.Server,
	//					config.GetYamlValues().ServerConfig.Port,
	//				))))
	//		r.Mount("/", apihandler.APIHandler(service))
	//	})
	server := newServer(r)

	return server
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
