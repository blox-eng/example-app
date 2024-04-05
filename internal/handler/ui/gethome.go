package uihandler

import (
	"github.com/blox-eng/app/internal/middleware"
	"github.com/blox-eng/app/internal/model"
	"github.com/blox-eng/app/internal/templates"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.UserKey).(*model.User)
	if !ok {
		c := templates.GuestIndex()
		err := templates.Layout(c, "Blox").Render(r.Context(), w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
		return
	}

	c := templates.Index(user.Email)
	err := templates.Layout(c, "Blox").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
