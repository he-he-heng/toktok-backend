package token

import (
	"fmt"
	"time"
	"toktok-backend/internal/adapter/config"
	"toktok-backend/internal/core/domain"

	"toktok-backend/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	key []byte

	accessDuration  time.Duration
	refreshDuration time.Duration
}

func New(config *config.Config) (*Token, error) {
	accessDuration, err := time.ParseDuration(config.Token.AccessDuration)
	if err != nil {
		return nil, errors.Wrap(domain.ErrTokenDuration, err)
	}

	refreshDuration, err := time.ParseDuration(config.Token.RefreshDuration)
	if err != nil {
		return nil, errors.Wrap(domain.ErrTokenDuration, err)
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
		"id":        user.ID,
		"role":      user.Role,
		"exp":       time.Now().Add(d).Unix(),
		"ita":       time.Now().Unix(),
		"tokenType": tokenType,
	})

	tokenString, err := token.SignedString(t.key)
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	return tokenString, nil
}

func (t *Token) VerifyToken(tokenString string) (*domain.TokenPlayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.Wrap(domain.ErrInvalidTokenSignMethod, err)
		}

		return t.key, nil
	})
	if err != nil {
		return nil, errors.Wrap(domain.ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, domain.ErrInvalidToken
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, domain.ErrExpiredToken
	}

	return &domain.TokenPlayload{
		ID:        int(claims["id"].(float64)),
		Role:      domain.RoleType(claims["role"].(string)),
		Exp:       int64(claims["exp"].(float64)),
		Ita:       int64(claims["ita"].(float64)),
		TokenType: domain.TokenType(claims["tokenType"].(string)),
	}, nil
}
