package jwt

import (
	"fmt"
	"time"
	"toktok-backend/internal/adapter/config"
	"toktok-backend/internal/core/domain"

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
		return nil, domain.ErrTokenDuration
	}

	refreshDuration, err := time.ParseDuration(config.Token.RefreshDuration)
	if err != nil {
		return nil, domain.ErrTokenDuration
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
		"uid":       user.UID,
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
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return t.key, nil
	})
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, domain.ErrInvalidToken
	}

	if time.Now().Unix() > claims["exp"].(int64) {
		return nil, domain.ErrExpiredToken
	}

	return &domain.TokenPlayload{
		UID:       claims["uid"].(int),
		Role:      claims["role"].(domain.RoleType),
		Exp:       claims["exp"].(int64),
		Ita:       claims["ita"].(int64),
		TokenType: claims["tokenType"].(domain.TokenType),
	}, nil
}
