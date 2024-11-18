package auth

import (
	"errors"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Sub uuid.UUID
	jwt.RegisteredClaims
}

func GenerateAccessToken(userId uuid.UUID) (string, error) {
	claims := &JwtClaims{
		Sub: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(userId uuid.UUID) (string, error) {
	claims := &JwtClaims{
		Sub: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
			Issuer:    os.Getenv("JWT_REFRESH_ISSUER"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(signedToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		},
	)
	if err != nil {
		return nil, errors.New("failed to parse token")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Add(-5*time.Minute).Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}

func ValidateRefreshToken(signedToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_REFRESH_SECRET_KEY")), nil
		},
	)
	if err != nil {
		return nil, errors.New("failed to parse token")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Add(-60*time.Minute).Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
