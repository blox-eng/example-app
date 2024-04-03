package uihandler

import (
	"net/http"

	"github.com/blox-eng/app/internal/model"
	"github.com/blox-eng/app/internal/service"
	"github.com/blox-eng/app/internal/templates"
)

type PostRegisterHandler struct {
	store service.Service
}

type PostRegisterHandlerParams struct {
	Store service.Service
}

func NewPostRegisterHandler(params PostRegisterHandlerParams) *PostRegisterHandler {
	return &PostRegisterHandler{
		store: params.Store,
	}
}
func (h *PostRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := h.store.CreateUser(&model.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c := templates.RegisterError()
		c.Render(r.Context(), w)
		return
	}
	c := templates.RegisterSuccess()
	err = c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
