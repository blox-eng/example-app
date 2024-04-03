package uihandler

import (
	"net/http"

	"github.com/blox-eng/app/internal/templates"
)

type GetRegisterHandler struct{}

func NewGetRegisterHandler() *GetRegisterHandler {
	return &GetRegisterHandler{}
}

func (h *GetRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.RegisterPage()
	er := templates.Layout(c, "My website").Render(r.Context(), w)

	if er != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
