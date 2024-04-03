package uihandler

import (
	"net/http"

	"github.com/blox-eng/app/internal/templates"
)

type AboutHandler struct{}

func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.About()
	er := templates.Layout(c, "My website").Render(r.Context(), w)

	if er != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return

	}
}
