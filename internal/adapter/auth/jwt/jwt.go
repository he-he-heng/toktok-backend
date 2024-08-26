package jwt

import (
	"fmt"
	"time"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/infra/config"
	"toktok-backend/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	key []byte

	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewJWT(config *config.Config) (*Token, error) {
	accessDuration, err := time.ParseDuration(config.Token.AccessDuration)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServerError, err)
	}

	refreshDuration, err := time.ParseDuration(config.Token.RefreshDuration)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServerError, err)
	}

	return &Token{
		key:             []byte(config.Token.Key),
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}, nil

}

func (t *Token) CreateToken(tokenType domain.TokenType, user *domain.User) (string, error) {
	var d time.Duration
	switch tokenType {
	case domain.AccessToken:
		d = t.accessDuration
	case domain.RefreshToken:
		d = t.refreshDuration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       user.ID,
		"exp":       time.Now().Add(d).Unix(),
		"ita":       time.Now().Unix(),
		"tokenType": tokenType,
		"role":      user.Role,
	})

	tokenString, err := token.SignedString(t.key)
	if err != nil {
		return "", errors.Wrap(domain.ErrInternalServerError, err)
	}

	return tokenString, nil
}

func (t *Token) VerifyToken(tokenString string) (*domain.TokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Wrap(domain.ErrUnauthorized, fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}

		return t.key, nil
	})
	if err != nil {
		return nil, errors.Wrap(domain.ErrUnauthorized, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.Wrap(domain.ErrUnauthorized, err)
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, errors.Wrap(domain.ErrUnauthorized, err)
	}

	return &domain.TokenPayload{
		Iss:       int(claims["id"].(float64)),
		Exp:       int64(claims["exp"].(float64)),
		Ita:       int64(claims["ita"].(float64)),
		TokenType: domain.TokenType(claims["tokenType"].(string)),
		Role:      domain.UserRoleType(claims["role"].(string)),
	}, nil
}
