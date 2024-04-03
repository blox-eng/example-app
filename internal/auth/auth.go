package auth

import (
	"github.com/blox-eng/app/internal/model"
	"github.com/golang-jwt/jwt"
)

type TokenAUth interface {
	GenerateToken(user model.User) (string, error)
	ValidateToken(token string) (jwt.MapClaims, error)
}
