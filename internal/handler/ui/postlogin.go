package uihandler

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/blox-eng/app/internal/hash"
	"github.com/blox-eng/app/internal/model"
	"github.com/blox-eng/app/internal/service"
	"github.com/blox-eng/app/internal/templates"
	"net/http"
	"time"
)

type PostLoginHandler struct {
	store             service.Service
	passwordHasher    *hash.PasswordHash
	sessionCookieName string
}

type PostLoginHandlerParams struct {
	Store             service.Service
	PasswordHasher    *hash.PasswordHash
	SessionCookieName string
}

func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		store:             params.Store,
		passwordHasher:    params.PasswordHasher,
		sessionCookieName: params.SessionCookieName,
	}
}
func (h *PostLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	users, err := h.store.ListUsers(&model.User{
		Email: email,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}
	if len(users) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	user := users[0]
	passwordIsValid, err := h.passwordHasher.ComparePasswordAndHash(password, user.Password)
	if err != nil || !passwordIsValid {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}
	session, err := h.store.CreateSession(&model.Session{
		UserID: user.ID,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID := user.ID
	sessionID := session.Session.SessionID

	cookieValue := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", sessionID, userID)))
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     h.sessionCookieName,
		Value:    cookieValue,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)

}
