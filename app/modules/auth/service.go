package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	RefreshToken(tokenString string) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte(os.Getenv("JWT_SECRET"))

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID string) (string, error) {
	jwtTTL := 60

	if os.Getenv("JWT_TTL") != "" {
		if customTTL, err := strconv.Atoi(os.Getenv("JWT_TTL")); err == nil {
			jwtTTL = customTTL
		}
	}

	expirationTime := time.Now().Add(time.Duration(jwtTTL) * time.Minute)

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *jwtService) RefreshToken(tokenString string) (string, error) {
	// Parse token
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Generate new token with the same user ID
	userID := claims["user_id"].(string)
	newToken, err := s.GenerateToken(userID)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
