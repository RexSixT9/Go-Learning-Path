package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func CreateToken(jwtSecret string, userID string, role string) (string, error) {
	if jwtSecret == "" {
		return "", errors.New("JWT secret is required")
	}
	now := time.Now().UTC()
	expirationTime := now.Add(7 * 24 * time.Hour) // Token valid for 7 days

	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(jwtSecret string, tokenString string) (Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return Claims{}, errors.New("token expired")
		}
		return Claims{}, fmt.Errorf("invalid token: %w", err)
	}
	if !token.Valid {
		return Claims{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return Claims{}, errors.New("failed to parse claims")
	}

	subject, err := claims.GetSubject()
	if err != nil || subject == "" {
		return Claims{}, errors.New("token does not contain a valid subject")
	}

	return *claims, nil
}
